package main

import (
	"bilibiliBulletscreenCrawler/pkg/filecontrol"
	"bilibiliBulletscreenCrawler/utils"
	"fmt"
	"log"
	"strconv"
)

func main() {
	bvNum := "BV1fb421J7um" //bilibiliVideo's Bvid
	keyWord := "RukiaOvO"   //Bulletscreen keywords (if you don't want to search special bulletscreen, use the keyword "RukiaOvO")
	err := utils.GetCidByBvNum(bvNum)
	if err != nil {
		log.Fatalf("Failed to get Cid: %v\n", err)
	}
	for i := 1; i <= utils.MultipageData.Data.View.Videos; i++ {
		fileName := fmt.Sprintf("%s[P%d]", bvNum, i)
		filecontrol.InitFile(fileName)

		if err := utils.GetPageDataByCid(strconv.FormatInt(utils.MultipageData.Data.View.Page[i-1].Cid, 10)); err != nil {
			log.Fatalf("Failed to get page data: %v\n", err)
		}
		filecontrol.WriteFileData(fileName, keyWord)
	}
}
