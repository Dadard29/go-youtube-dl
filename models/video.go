package models

import "time"

type VideoModel struct {
	Id int `gorm:"type:int;index:id;primary_key;auto_increment"`
	VideoId string `gorm:"type:varchar(30);index:video_id"`
	Token string `gorm:"type:varchar(30);index:token"`
	Title string `gorm:"type:varchar(30);index:title"`
	Album string `gorm:"type:varchar(30);index:album"`
	Artist string `gorm:"type:varchar(30);index:artist"`
	Date time.Time `gorm:"type:date;index:date"`
}

func (VideoModel) TableName() string {
	return "video"
}

func (v VideoModel) GetUrl() string {
	return "https://youtube.com/watch?v=" + v.VideoId
}

type VideoJson struct {
	VideoId string
	Title string
	Album string
	Artist string
	Date string
}

func NewVideoJson(v VideoModel) VideoJson {
	return VideoJson{
		VideoId: v.VideoId,
		Title:   v.Title,
		Album:   v.Album,
		Artist:  v.Artist,
		Date:    v.Date.String(),
	}
}

type VideoYoutubeModel struct {
	Kind     string `json:"kind"`
	Etag     string `json:"etag"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
				Standard struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"standard"`
				Maxres struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"maxres"`
			} `json:"thumbnails"`
			ChannelTitle         string   `json:"channelTitle"`
			Tags                 []string `json:"tags"`
			CategoryID           string   `json:"categoryId"`
			LiveBroadcastContent string   `json:"liveBroadcastContent"`
			Localized            struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"localized"`
		} `json:"snippet"`
		ContentDetails struct {
			Duration        string `json:"duration"`
			Dimension       string `json:"dimension"`
			Definition      string `json:"definition"`
			Caption         string `json:"caption"`
			LicensedContent bool   `json:"licensedContent"`
			Projection      string `json:"projection"`
		} `json:"contentDetails"`
		Statistics struct {
			ViewCount     string `json:"viewCount"`
			LikeCount     string `json:"likeCount"`
			DislikeCount  string `json:"dislikeCount"`
			FavoriteCount string `json:"favoriteCount"`
			CommentCount  string `json:"commentCount"`
		} `json:"statistics"`
	} `json:"items"`
}

type PlaylistYoutubeModel struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []PlaylistYoutubeItem
}

type PlaylistYoutubeItem struct {
	Kind    string `json:"kind"`
	Etag    string `json:"etag"`
	ID      string `json:"id"`
	Snippet struct {
		PublishedAt time.Time `json:"publishedAt"`
		ChannelID   string    `json:"channelId"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Thumbnails  struct {
			Default struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"default"`
			Medium struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"medium"`
			High struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"high"`
		} `json:"thumbnails"`
		ChannelTitle string `json:"channelTitle"`
		PlaylistID   string `json:"playlistId"`
		Position     int    `json:"position"`
		ResourceID   struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"resourceId"`
	} `json:"snippet"`
}

