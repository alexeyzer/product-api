package datastruct

const MediaTableName = "media"

type ContentType string

const (
	photo_ContentType ContentType = "PHOTO"
	video_ContentType ContentType = "VIDEO"
)

type Media struct {
	ID          int64       `db:"id"`
	ProductID   int64       `db:"product_id"`
	ContentType ContentType `db:"content_type"`
	Url         string      `db:"url"`
}
