package model

import "github.com/degalaxy/helper/dahelper"

type GetPartCandidatesReq struct {
	SpaceId    string   `json:"spaceId"`
	AssetTypes []string `json:"assetTypes"`
}

type GetPartCandidatesRsp struct {
	ErrorCode      int                        `json:"errcode"`
	ErrorMessage   string                     `json:"errmsg"`
	PartCandidates map[string]*PartCandidates `json:"data"`
}

type PartCandidates struct {
	Items []*PartCandiateItem `json:"items"`
}

type PartCandiateItem struct {
	ImageUrl string `json:"imageUrl"`
	ImageMd5 string `json:"imageMd5"`
	Name     string `json:"name"`
	SubName  string `json:"subName"`
}

type CollectWelcomeGiftsReq struct {
	Did string `json:"did"`
}

type CollectTimePeriodAssetsReq struct {
	Did          string `json:"did"`
	TimePeriodId string `json:"timePeriodId"`
}

type PreUploadCosReq struct {
	RootFolderName string `json:"rootFolderName"`
	CosFolderName  string `json:"cosFolderName"`
}

type PreUploadCosRsp struct {
	ErrorCode    int                    `json:"errcode"`
	ErrorMessage string                 `json:"errmsg"`
	Items        []*PreUploadCosRspItem `json:"items"`
}

type PreUploadCosRspItem struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	FilePath     string `json:"filePath"`
}

type RanksContestantReq struct {
	EventId       string                   `json:"eventId"`
	AssetContract string                   `json:"assetContract"`
	AssetId       string                   `json:"assetId"`
	CreatorId     string                   `json:"creatorId"`
	ImageUrl      string                   `json:"imageUrl"`
	Metadata      *dahelper.CustomMetadata `json:"metadata"`
}

type RanksContestantRsp struct {
	ErrCode    string `json:"errcode"`
	ErrMessage string `json:"message"`
}

type UpgradeCombinedIssueReq struct {
	RecipientUserId string         `json:"recipientUserId"`
	Did             string         `json:"did"`
	Assets          []*AssetDetail `json:"assets"`
	ContractAddress string         `json:"contractAddress"`
}

type UpgradeCombinedIssueRsp struct {
	ErrorCode    int                     `json:"errcode"`
	ErrorMessage string                  `json:"errmsg"`
	AssetDetail  *dahelper.DAAssetDetail `json:"data"`
}
