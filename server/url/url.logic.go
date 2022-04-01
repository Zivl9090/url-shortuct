package url

import "url-shortcut/structs"

func Logic_GetUrl(shortcutUrl string) *structs.Url {
	var url *structs.Url = Dal_GetUrl(shortcutUrl)
	return url
}

func Logic_CreateUrl(shortcutUrl string) {
	Dal_CreateUrl(shortcutUrl)
}
