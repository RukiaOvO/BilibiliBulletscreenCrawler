package model

type MultipageData struct {
	Data Data `json:"data"`
}

type Data struct {
	View View `json:"view"`
}

type View struct {
	Page   []Page `json:"pages"`
	Videos int    `json:"videos"`
}

type Page struct {
	Cid int64 `json:"cid"`
}
