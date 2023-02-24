package model

type TimePeriod struct {
	Id                string `json:"id"`
	TimePeriodGroupId string `json:"timePeriodGroupId"`
	SpaceId           string `json:"spaceId"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Status            int    `json:"status"`
	ValidFromUnixSec  int64  `json:"validFromUnixSec"`
	ValidToUnixSec    int64  `json:"validToUnixSec"`
}

type GetActiveTimePeriodsRsp struct {
	ErrorCode    int           `json:"errcode"`
	ErrorMessage string        `json:"errmsg"`
	TimePeriods  []*TimePeriod `json:"data"`
}

type GetTimePeriodByIdRsp struct {
	ErrorCode    int         `json:"errcode"`
	ErrorMessage string      `json:"errmsg"`
	TimePeriod   *TimePeriod `json:"data"`
}

type TimePeriodUserAssetAmount struct {
	Id                string `json:"id"`
	OwnerId           string `json:"ownerId"`
	TimePeriodAssetId string `json:"timePeriodAssetId"`
	AssetName         string `json:"assetName"`
	Amount            int    `json:"amount"`
	ProcessingCount   int    `json:"processingCount"`
	AmountLimit       int    `json:"amountLimit"`
	Version           int    `json:"version"`
}

type GetUserTimePeriodAssetAmountsRsp struct {
	ErrorCode                  int                          `json:"errcode"`
	ErrorMessage               string                       `json:"errmsg"`
	TimePeriodUserAssetAmounts []*TimePeriodUserAssetAmount `json:"data"`
}

type GetSingleUserTimePeriodAssetAmountRsp struct {
	ErrorCode                 int                        `json:"errcode"`
	ErrorMessage              string                     `json:"errmsg"`
	TimePeriodUserAssetAmount *TimePeriodUserAssetAmount `json:"data"`
}

type InitUserTimePeriodAssetAmountsReq struct {
	UserId       string `json:"userId"`
	TimePeriodId string `json:"timePeriodId"`
}

type UpdateUserTimePeriodAssetAmountsReq struct {
	UserId                      string `json:"userId"`
	TimePeriodUserAssetAmuontId string `json:"timePeriodUserAssetAmuontId"`
	NewAmount                   int    `json:"newAmount"`
	NewProcessingCount          int    `json:"newProcessingCount"`
	CurrentVersion              int    `json:"currentVersion"`
}
