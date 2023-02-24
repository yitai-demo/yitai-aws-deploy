package model

type CommReportReq struct {
	ReportData *CommReportData `json:"reportData"`
}

type CommReportData struct {
	ReportDataInfo     *CommReportDataInfo   `json:"reportDataInfo"`
	ReportDataItemList []*CommReportDataItem `json:"reportDataItemList"`
}

type CommReportDataInfo struct {
	AppVersion string `json:"appVersion"`
	Platform   int    `json:"platform"`
	UserId     string `json:"userId"`
	LoginType  int    `json:"loginType"`
	EventTime  int64  `json:"eventTime"`
}

type CommReportDataItem struct {
	OperId string `json:"operId"`
	Ext1   string `json:"ext1"`
	Ext2   string `json:"ext2"`
	Ext3   string `json:"ext3"`
	Ext4   string `json:"ext4"`
	Ext5   string `json:"ext5"`
	Ext6   string `json:"ext6"`
	Ext7   string `json:"ext7"`
	Ext8   string `json:"ext8"`
	Ext9   string `json:"ext9"`
	Ext10  string `json:"ext10"`
}

type CommReportRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

type GetReportReq struct {
	Sign       string `json:"sign"`
	ReportName string `json:"report_name"`
}

type QueryReportDirReq struct {
	Sign string `json:"sign"`
}

type QueryReportDirRsp struct {
	Reports []string `json:"reports"`
}
