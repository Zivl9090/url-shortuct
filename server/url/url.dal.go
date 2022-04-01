package url

import (
	"url-shortcut/db"
	"url-shortcut/structs"
)

func Dal_GetUrl(shortcutUrl string) *structs.Url {
	var url structs.Url
	db.Db.First(&url, "shortcut_url = ?", "TEST")
	return &url
}

func Dal_CreateUrl(originalUrl string) {
	db.Db.Create(&structs.Url{OriginalUrl: originalUrl, ShortcutUrl: "TEST"})
}
