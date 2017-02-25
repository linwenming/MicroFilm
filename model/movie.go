package model

type Movie struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Intro      string  `json:"intro"`
	CategoryId int64   `db:"category_id" json:"categoryId"`
	Director   string  `json:"director"`
	Actor      string  `json:"actor"`
	Highlight  string  `json:"highlight"`
	PlayTime   int64   `db:"play_time" json:"playTime"`
	PreviewUrl string  `db:"preview_url" json:"previewUrl"`
	MovieUrl   string  `db:"movie_url" json:"movieUrl"`
	Tags       string  `json:"tags"`
	Level      int    `json:"level"`
	Status     int     `json:"status"`
	Score      int     `json:"score"`
	PlayCount  int   `db:"play_count" json:"playCount"`
	ReplyCount int   `db:"reply_count" json:"replyCount"`
	ZanCount   int   `db:"zan_count" json:"zanCount"`
	Uploader   int64   `json:"uploader"`
	CreateTime int64   `db:"create_time" json:"createTime"`
	OnlineTime int64   `db:"online_time" json:"onlineTime"`
}

func NewMovie() *Movie {
	return &Movie{}
}