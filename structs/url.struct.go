package structs

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	OriginalUrl string
	ShortcutUrl string
}
