package repositories

import (
	"archive/zip"
	"fmt"
	"github.com/BrianAllred/goydl"
	"github.com/Dadard29/go-youtube-dl/api"
	"github.com/Dadard29/go-youtube-dl/models"
	"github.com/bogem/id3v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)


const Store = "./placeholder"
const archiveFile = "mp3_files.zip"

func CheckPlaceholder(token string) bool {
	path2check := path.Join(Store, token)
	if _, err := os.Stat(path2check); os.IsNotExist(err) {
		if err := os.Mkdir(path2check, 0755); err != nil {
			return false
		}
		return true
	}
	return true
}

func Download(vModel models.VideoModel, path string) {

	vUrl := vModel.GetUrl()
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.Options.Output.Value = path
	youtubeDl.Options.ExtractAudio.Value = true
	youtubeDl.Options.AudioFormat.Value = "mp3"

	cmd, err := youtubeDl.Download(vUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SetID3v2Tags(path string, model models.VideoModel) error {
	tag, err := id3v2.Open(path, id3v2.Options{Parse: true})
	if err != nil {
		return err
	}
	defer tag.Close()

	// Set simple text frames.
	tag.SetTitle(model.Title)
	tag.SetArtist(model.Artist)
	tag.SetAlbum(model.Album)
	tag.SetYear(string(model.Date.Year()))

	// Write it to file.
	if err = tag.Save(); err != nil {
		return err
	}

	return nil
}

func RenameFile(srcFile string, dstFile string) error {
	return os.Rename(srcFile, dstFile)
}

func CleanPlaceholder(token string) error {
	return cleanPath(path.Join(Store, token))
}

func cleanPath(srcPath string) error {
	dir, err := ioutil.ReadDir(srcPath)
	if err != nil {
		return err
	}
	for _, d := range dir {
		err := os.RemoveAll(path.Join([]string{srcPath, d.Name()}...))
		if err != nil {
			return err
		}
	}

	return nil
}

func CleanMp3Files(token string) error {
	p := path.Join(Store, token)
	dir, err := ioutil.ReadDir(p)
	if err != nil {
		return err
	}
	for _, f := range dir {
		if f.IsDir() || !strings.Contains(f.Name(), ".mp3") {
			continue
		}

		err := os.Remove(path.Join(p, f.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func ArchiveFiles(token string) error {
	srcPath := path.Join(Store, token)
	zipFile, err := os.Create(path.Join(srcPath, archiveFile))
	if err != nil {
		return err
	}

	defer zipFile.Close()
	archive := zip.NewWriter(zipFile)

	dir, err := ioutil.ReadDir(srcPath)
	if err != nil {
		return err
	}

	for _, d := range dir {
		if d.IsDir() || !strings.Contains(d.Name(), ".mp3") {
			continue
		}

		f, err := archive.Create(d.Name())
		if err != nil {
			return err
		}

		originFile, err := os.OpenFile(path.Join(srcPath, d.Name()), os.O_RDONLY, 0644)
		if err != nil {
			return err
		}

		body, err := ioutil.ReadAll(originFile)
		if err != nil {
			return err
		}

		_, err = f.Write(body)
		if err != nil {
			return err
		}
	}

	err = archive.Close()
	if err != nil {
		return err
	}

	return nil
}

func NewStatus(token string, message string) {
	s := models.Status{
		Token:        token,
		DateStarted:  time.Now(),
		DateFinished: time.Now(),
		Progress:     0,
		Done:         false,
		Message:      message,
	}

	api.Api.Database.Orm.Create(&s)
}

func UpdateStatus(token string, progress int, message string) {
	var s models.Status
	api.Api.Database.Orm.Where(&models.Status{
		Token: token,
	}).Last(&s)

	s.Progress = progress
	s.Message = message

	api.Api.Database.Orm.Save(&s)
}

func EndStatus(token string, message string) {
	var sList []models.Status
	api.Api.Database.Orm.Where(&models.Status{
		Token: token,
		Done: false,
	}).Find(&sList)

	for _, s := range sList {
		if s.Done == false {
			s.Done = true
			s.Progress = 100
			s.Message = message
			s.DateFinished = time.Now()

			api.Api.Database.Orm.Save(&s)
		}
	}
}

func GetStatus(token string) (models.Status, error) {
	var s models.Status
	api.Api.Database.Orm.Where(&models.Status{
		Token: token,
	}).Last(&s)

	if s.Token != token {
		// no status found yet
		return models.Status{
			StatusId:     0,
			Token:        token,
			DateStarted:  time.Time{},
			DateFinished: time.Time{},
			Progress:     100,
			Done:         true,
			Message:      "no download done yet",
		}, nil
	}

	return s, nil
}

