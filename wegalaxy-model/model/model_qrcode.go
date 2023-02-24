package model

type QrCode struct {
	Type      int    `json:"type"`
	Status    int    `json:"status"`
	BeginTime int64  `json:"beginTime"`
	EndTime   int64  `json:"endTime"`
	SpaceId   string `json:"spaceId"`
	Action    int    `json:"action"`
	JumpUrl   string `json:"jumpUrl"`
	Text      string `json:"text"`
}

type ProcessQrCodeReq struct {
	QrCodeId          string `json:"c"`
	QrCodeDisplayName string `json:"n"`
}

type ProcessQrCodeRsq struct {
	ErrorCode    int         `json:"errcode"`
	ErrorMessage string      `json:"errmsg"`
	Data         interface{} `json:"data"` // It varies because of different actions
}

type CheckInRspData struct {
	Action            string `json:"action"`
	IsMedalCollected  bool   `json:"isMedalCollected"`
	CollectMedalError string `json:"collectMedalError"`
	MedalName         string `json:"medalName"`
	JumpToUrl         string `json:"jumpToUrl"`
}
