package managers

import (
	"errors"
	"fmt"
	"github.com/Dadard29/go-youtube-dl/models"
	"github.com/Dadard29/go-youtube-dl/repositories"
	"io/ioutil"
	"path"
)

func GetDownloadFile(token string) (string, error) {
	// check if operation occurring now
	s, err := repositories.GetStatus(token)
	if err == nil && !s.Done {
		return "", errors.New("an operation is currently being processed")
	}

	if !repositories.CheckPlaceholder(token) {
		err := errors.New("error checking placeholder")
		return "", err
	}

	p := path.Join(repositories.Store, token)
	dir, err := ioutil.ReadDir(p)
	if err != nil {
		return "", err
	}

	if len(dir) == 0 {
		return "", errors.New("your placeholder is emtpy")
	}

	if len(dir) > 1 {
		return "", errors.New("your placeholder is a f*ckin mess")
	}

	return path.Join(repositories.Store, token, dir[0].Name()), nil


}

func GetDownloadStatus(token string) (models.StatusJson, error) {
	var j models.StatusJson
	s, err := repositories.GetStatus(token)

	if err != nil {
		return j, err
	}

	return models.NewStatusJson(s), nil
}

func CancelDownload(token string) {
	repositories.EndStatus(token, "cancelled by user")
}

func Download(token string, videoId string) error {
	// check if operation occurring now
	s, err := repositories.GetStatus(token)
	if err == nil && !s.Done {
		return errors.New("an operation is currently being processed")
	}

	vModel, err := repositories.VideoGet(token, videoId)
	if err != nil {
		return err
	}

	if !repositories.CheckPlaceholder(token) {
		err := errors.New("error checking placeholder")
		return err
	}
	repositories.CleanPlaceholder(token)
	repositories.NewStatus(token, "download started")

	tempFilename := "tmp"
	tempFilePath := path.Join(repositories.Store, token, tempFilename + ".mp3")

	go func() {
		logger.Info("download started")
		repositories.Download(vModel, path.Join(repositories.Store, token), tempFilename)
		err = repositories.SetID3v2Tags(tempFilePath, vModel)
		if err != nil {
			logger.Error(err.Error())
			repositories.CleanPlaceholder(token)
			repositories.EndStatus(token, "error setting ID3 tags")
			return
		}

		filename := fmt.Sprintf("%s.mp3", vModel.Title)
		outputFile := path.Join(repositories.Store, token, filename)
		err = repositories.RenameFile(tempFilePath, outputFile)
		if err != nil {
			logger.Error(err.Error())
			repositories.CleanPlaceholder(token)
			repositories.EndStatus(token, "error storing file")
			return
		}


		repositories.EndStatus(token, "download done")
		logger.Info("download done")
	}()


	return nil
}

func DownloadAll(token string) error {
	// check if operation occurring now
	s, err := repositories.GetStatus(token)
	if err == nil && !s.Done {
		return errors.New("an operation is currently being processed")
	}

	vList, err := repositories.VideoGetList(token)
	if err != nil {
		return err
	}

	if !repositories.CheckPlaceholder(token) {
		err := errors.New("error checking placeholder")
		return err
	}
	repositories.CleanPlaceholder(token)
	repositories.NewStatus(token, "download started")

	tempFilename := "tmp"
	tempFilePath := path.Join(repositories.Store, token, tempFilename + ".mp3")

	go func() {
		logger.Info("download started")

		var status = 0.00
		var unit = 100.0 / float64(len(vList))
		for _, m := range vList {
			repositories.Download(m, path.Join(repositories.Store, token), tempFilename)

			err = repositories.SetID3v2Tags(tempFilePath, m)
			if err != nil {
				logger.Error(fmt.Sprintf("error setting id3v2 tags: %v", err))
				continue
			}

			filename := fmt.Sprintf("%s.mp3", m.Title)
			outputFile := path.Join(repositories.Store, token,filename)
			err = repositories.RenameFile(tempFilePath, outputFile)
			if err != nil {
				logger.Error(fmt.Sprintf("error renaming file: %v", err))
				continue
			}

			status += unit
			repositories.UpdateStatus(token, int(status),
				fmt.Sprintf("progress is %d%%", int(status)))

			logger.Info(fmt.Sprintf("(%f%% done) downloaded and stored %s", status, filename))
		}

		err := repositories.ArchiveFiles(token)
		if err != nil {
			logger.Error(err.Error())
			repositories.CleanPlaceholder(token)
			repositories.EndStatus(token, "error archiving mp3 files")
		}

		err = repositories.CleanMp3Files(token)
		if err != nil {
			logger.Error(err.Error())
			repositories.CleanPlaceholder(token)
			repositories.EndStatus(token, "error cleaning up mp3 files")
		}


		repositories.EndStatus(token, "download done")
		logger.Info("download done")
	}()


	return nil
}