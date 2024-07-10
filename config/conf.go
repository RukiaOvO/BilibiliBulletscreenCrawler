package config

import (
	"bilibiliBulletscreenCrawler/database/model"
	"bilibiliBulletscreenCrawler/utils"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func HashMapDataCheck(barrage *model.BulletscreenData, vp *viper.Viper, kw string) {
	if exist := strings.Contains(barrage.Value, kw); exist && kw != "RukiaOvO" {
		data := vp.GetString(barrage.P)
		if data == "" {
			vp.Set(barrage.P, utils.CheckHashCode(barrage.P))
			if err := vp.WriteConfig(); err != nil {
				log.Fatalf("Failed to write config file: %v", err)
			}
			fmt.Printf("HachMap add:  %s\n", barrage.P)
		} else {
			fmt.Printf("Data exists: %v-%v\n", barrage.P, data)
		}
		barrage.Uid = vp.GetString(barrage.P)
	} else {
		barrage.Uid = "null"
	}
}
