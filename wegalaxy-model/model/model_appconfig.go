package model

type SimpleTrueOrFalseRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	IsSuccess    bool   `json:"data"`
}
type AppConfig struct {
	AppId string `json:"appId"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GetAppConfigReq struct {
	AppId string `json:"appId"`
	Name  string `json:"name"`
}

type GetAppConfigRsp struct {
	ErrorCode    int        `json:"errcode"`
	ErrorMessage string     `json:"errmsg"`
	AppConfig    *AppConfig `json:"data"`
}
