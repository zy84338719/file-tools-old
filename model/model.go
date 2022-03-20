package model

import "sync"

var WG sync.WaitGroup

type UserInfo struct {
	Code int `json:"code"`
	Data struct {
		UserList []struct {
			DiscoverType int `json:"discover_type"`
			UserID       int `json:"user_id"`
			Covers       []struct {
				Ord          int    `json:"ord"`
				Name         string `json:"name"`
				MusicID      int    `json:"music_id"`
				VerifyStatus int    `json:"verify_status"`
				CoverID      int    `json:"cover_id"`
				Type         int    `json:"type"`
				Url          string `json:"url"`
			} `json:"covers"`
		} `json:"user_list"`
	} `json:"data"`
	Message string `json:"message"`
}
