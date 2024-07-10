package filecontrol

import (
	"bilibiliBulletscreenCrawler/config"
	"bilibiliBulletscreenCrawler/utils"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup

func InitFile(bvNum string) {
	fileAdr := fmt.Sprintf("./database/data/%s.txt", bvNum)
	file, err := os.OpenFile(fileAdr, os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open file: %v, recreating\n", err)
		if file, err = os.Create(fileAdr); err != nil {
			log.Fatalf("Failed to create file: %v\n", err)
		}
	}
	defer file.Close()

	_, err = file.WriteString("--------------------------------Barrage List--------------------------------\nNumber   Crc32Code   DecodeUid   Comment\n")
	if err != nil {
		log.Fatalf("Failed to init file when writing: %v\n", err)
	}
}

func WriteFileData(bvNum string, keyWord string) {
	vp := viper.New()
	vp.AddConfigPath("./config/local")
	vp.SetConfigType("yaml")
	vp.SetConfigName("hashMap")
	if err := vp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("Failed to find config file: %v, recreating", err)
			if _, err := os.Create("./config/local/hashMap.yaml"); err != nil {
				log.Fatalf("Failed to recreate config file: %v", err)
			}
		} else {
			log.Fatalf("Failed to read config file: %v", err)
		}
	}

	wg.Add(len(utils.Data.BulletscreenXmlData))
	for c := range utils.Data.BulletscreenXmlData {
		mu.Lock()
		go func(i int) {
			defer wg.Done()
			config.HashMapDataCheck(&utils.Data.BulletscreenXmlData[i], vp, keyWord)
			fileAdr := fmt.Sprintf("./database/data/%s.txt", bvNum)
			file, err := os.OpenFile(fileAdr, os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("Failed to open file: %v\n", err)
			}
			defer file.Close()

			data := fmt.Sprintf("%5d     %s      %s     %s\n", i+1, utils.Data.BulletscreenXmlData[i].P, utils.Data.BulletscreenXmlData[i].Uid, utils.Data.BulletscreenXmlData[i].Value)
			_, err = file.WriteString(data)
			if err != nil {
				log.Fatalf("Failed to write data: %v\n", err)
			}
		}(c)
		mu.Unlock()
	}
	wg.Wait()
}
