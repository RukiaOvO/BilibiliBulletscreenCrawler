package utils

import (
	"bilibiliBulletscreenCrawler/database/model"
	"bytes"
	"compress/flate"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var MultipageData model.MultipageData
var Data model.XmlData

func GetCidByBvNum(bvNum string) error {
	getUrl := fmt.Sprintf("https://api.bilibili.com/x/web-interface/wbi/view/detail?platform=web&bvid=%s", bvNum)
	req, _ := http.NewRequest("GET", getUrl, nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	refererStr := fmt.Sprintf("https://www.bilibili.com/video/%s", bvNum)
	req.Header.Set("Referer", refererStr)
	req.Header.Set("authority", "api.bilibili.com")
	req.Header.Set("Cookie", "i-wanna-go-back=-1; buvid4=ABBB5DB9-32E0-E10F-FF96-7EB2AE6ED6F723884-022072517-Lbg0LAZDGBXEw6XYm2wm3Q%3D%3D; buvid_fp_plain=undefined; rpdid=|(u)luk)))Rm0J'uYYmJ)Y~)J; buvid3=369FBA10-1B0C-E35B-5DE0-5C936B7B3F7B69401infoc; b_nut=1691928469; b_ut=5; _uuid=455BEB10A-C6C1-FD3D-7438-F8871DFEF351068183infoc; hit-new-style-dyn=1; hit-dyn-v2=1; LIVE_BUVID=AUTO9216921107297984; DedeUserID=102770566; DedeUserID__ckMd5=5c782754d1330dce; is-2022-channel=1; CURRENT_BLACKGAP=0; enable_web_push=DISABLE; header_theme_version=CLOSE; FEED_LIVE_VERSION=V_WATCHLATER_PIP_WINDOW3; fingerprint=e07802331e41d53ceb628cefdd46ed92; CURRENT_QUALITY=80; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjA1MzQ1ODgsImlhdCI6MTcyMDI3NTMyOCwicGx0IjotMX0.eu5umMYzaijg9PQGPNfsJdlOj5ekIDRB9biz960Z6R0; bili_ticket_expires=1720534528; home_feed_column=5; browser_resolution=1707-932; PVID=10; bsource=search_baidu; bp_t_offset_102770566=951735448231739392; buvid_fp=e07802331e41d53ceb628cefdd46ed92; CURRENT_FNVAL=4048; b_lsid=B8210A492_190924C12AA; SESSDATA=5bbeb3f6%2C1735993400%2Ccdf82%2A72CjBp_g3ZrJx_sFDYxzD8_47SL7dtzz8eEelLMFjElOlNLpE5dfW49HRKKI_3NA5APVgSVkNsR19EaGZjUkJPcTRVV0xCc0wyMDlkdWdNa0I3Qnh2MnNzaHJLanFIYTFPZzJrQXpqRnBMYVpnNEZSNXg5Q2poWXRZWXlwb3VYWk82UUMyV19VNjFBIIEC; bili_jct=2f2c5f13c3cefb203236e14d7557423e; sid=6o6wfv2x")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send request: %v\n", err)
		return err
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Failed to receive response: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	doc, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(doc, &MultipageData)
	if err != nil {
		fmt.Printf("Failed to read decodeData: %v\n", err)
	}
	return nil
}

func GetPageDataByCid(Cid string) error {
	barrageStr := fmt.Sprintf("https://api.bilibili.com/x/v1/dm/list.so?oid=%s", Cid)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", barrageStr, nil)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send request: %v\n", err)
		return err
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Failed to receive response: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	doc, _ := io.ReadAll(resp.Body)
	return XmlDataControl(doc)
}

func XmlDataControl(in []byte) error {
	xmlDoc, err := FlateDecode(in)
	if err != nil {
		fmt.Printf("Failed to read decodeData: %v\n", err)
	}
	err = xml.Unmarshal(xmlDoc, &Data)
	if err != nil {
		fmt.Printf("Failed to unmarshal xmlData: %v\n", err)
		return err
	}

	for c, v := range Data.BulletscreenXmlData {
		Data.BulletscreenXmlData[c].P = strings.Split(v.P, ",")[6]
	}
	return nil
}

func FlateDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer reader.Close()

	return io.ReadAll(reader)
}
