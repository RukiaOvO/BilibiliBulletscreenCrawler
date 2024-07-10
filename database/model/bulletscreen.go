package model

import "encoding/xml"

type XmlData struct {
	XmlName             xml.Name           `xml:"i"`
	ChatServer          string             `xml:"chatserver"`
	ChatId              int                `xml:"chatid"`
	Mission             int                `xml:"mission"`
	MaxLimit            int                `xml:"maxlimit"`
	State               int                `xml:"state"`
	Real_Name           int                `xml:"real_name"`
	Source              string             `xml:"source"`
	BulletscreenXmlData []BulletscreenData `xml:"d"`
}

type BulletscreenData struct {
	P     string `xml:"p,attr"`
	Value string `xml:",chardata"`
	Uid   string
}
