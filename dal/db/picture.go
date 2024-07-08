package db

type Image struct {
	Pid int64  `json:"pid" gorm:"primaryKey;autoincrement"`
	Url string `json:"url"`
}
