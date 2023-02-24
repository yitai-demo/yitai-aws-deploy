package model

import "github.com/degalaxy/helper/dahelper"

type QueryCustomerNonMdotAssetsReq struct {
	Did               string                    `json:"did"`
	SceneId           string                    `json:"sceneId"`
	ContractAddresses []string                  `json:"contractAddresses"`
	PageInfo          *dahelper.DAPageInfoParam `json:"pageInfo"`
}

type QueryCustomerNonMdotAssetsRsp struct {
	ErrorCode          int                    `json:"errcode"`
	ErrorMessage       string                 `json:"errmsg"`
	QueryCustomerAsset *QueryCustomerAssetRet `json:"data"`
}

type QueryCustomerAssetRet struct {
	AssetList []*AssetDetail        `json:"assetList"`
	PageInfo  *dahelper.DAPageInfo2 `json:"pageInfo"`
}

type GetAssetDetailReq struct {
	ContractAddress string `json:"contractAddress"`
	ContractMerits  string `json:"contractMerits"`
	Did             string `json:"did"`
}

type GetAssetDetailRsp struct {
	ErrorCode    int          `json:"errcode"`
	ErrorMessage string       `json:"errmsg"`
	AssetDetail  *AssetDetail `json:"data"`
}

type AssetDetail struct {
	AssetId            string                   `json:"id"`
	Type               string                   `json:"type"`
	IssuerId           string                   `json:"issuerId"`
	IssuerName         string                   `json:"issuerName"`
	OwnerName          string                   `json:"ownerName"`
	OwnerId            string                   `json:"ownerId"`
	ContractAddress    string                   `json:"contractAddress"`
	ContractType       string                   `json:"contractType"`
	Value              int                      `json:"value"`
	TransactionDetails []*TransactionDetail     `json:"transactionDetails"`
	CustomMetadata     *dahelper.CustomMetadata `json:"customMetadata"`
	Materials          []*Material              `json:"materials"`
}

type Material struct {
	Id              string `json:"id"`
	ContractAddress string `json:"contractAddress"`
}

type TransactionDetail struct {
	Id              string `json:"id"`
	ContractAddress string `json:"contractAddress"`
	ContractMerits  string `json:"contractMerits"`
	TransactionType string `json:"transactionType"`
	FromOwner       string `json:"fromOwner"`
	ToOwner         string `json:"toOwner"`
	CreatedAt       string `json:"createdAt"`
}

type TransactionType string

const (
	TRANSACTION_STATUS_ISSUE    TransactionType = "issue"
	TRANSACTION_STATUS_TRANSFER TransactionType = "transfer"
	TRANSACTION_STATUS_WRITEOFF TransactionType = "writeOff"
)

type BatchSafeSendReq struct {
	ContractAddress string                  `json:"contractAddress"`
	Items           []*BatchSafeSendReqItem `json:"items"`
	BlockchainId    string                  `json:"blockchainId"`
	Did             string                  `json:"did"`
	TransactionId   string                  `json:"transactionId"`
}

type BatchSafeSendReqItem struct {
	AssetId       string `json:"assetId"`
	RecipientName string `json:"recipient"`
}

type BatchSafeSendRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

type CombinedIssueWithItemUriReq struct {
	RecipientUserId   string                   `json:"recipientUserId"`
	Did               string                   `json:"did"`
	Assets            []*AssetDetail           `json:"assets"`
	ContractAddress   string                   `json:"contractAddress"`
	ImageUrl          string                   `json:"imageUrl"`
	ImageMd5          string                   `json:"imageMd5"`
	ImageThumbnailUrl string                   `json:"imageThumbnailUrl"`
	ImageThumbnailMd5 string                   `json:"imageThumbnailMd5"`
	IssueRule         string                   `json:"issueRule"`
	CustomMetadata    *dahelper.CustomMetadata `json:"customMetadata"`
	BlockchainId      string                   `json:"blockchainId"`
	Remark            string                   `json:"remark"`
	EventId           string                   `json:"eventId"`
	TransactionId     string                   `json:"transactionId"`
	SpaceId           string                   `json:"spaceId"`
}

type CombinedIssueWithItemUriRsp struct {
	ErrorCode    int                     `json:"errcode"`
	ErrorMessage string                  `json:"errmsg"`
	AssetDetail  *dahelper.DAAssetDetail `json:"data"`
}

type ImageThumbnailUrlRsp struct {
	ErrorCode         int    `json:"errcode"`
	ErrorMessage      string `json:"errmsg"`
	ImageThumbnailUrl string `json:"data"`
}

type BatchIssueReq struct {
	Did               string                   `json:"did"`
	DaToken           string                   `json:"daToken"`
	ContractAddress   string                   `json:"contractAddress"`
	BlockchainId      string                   `json:"blockchainId"`
	RecipientUserId   string                   `json:"recipientUserId"`
	IssueRule         string                   `json:"issueRule"`
	CustomMetadata    *dahelper.CustomMetadata `json:"customMetadata"`
	Remark            string                   `json:"remark"`
	ImageUrl          string                   `json:"imageUrl"`
	ImageMd5          string                   `json:"imageMd5"`
	ImageThumbnailUrl string                   `json:"imageThumbnailUrl"`
	ImageThumbnailMd5 string                   `json:"imageThumbnailMd5"`
	TransactionId     string                   `json:"transactionId"`
}

type BatchIssueRsp struct {
	ErrorCode    int                     `json:"errcode"`
	ErrorMessage string                  `json:"errmsg"`
	AssetDetail  *dahelper.DAAssetDetail `json:"data"`
}

type UploadFileExtension string

const (
	BLANK_CONPONENTS_NUMBER         int = 3
	Welcome_BLANK_CANVAS_NFT_NUMBER int = 5

	GetUserByIdUri = "/v1/wegalaxy-foundation/internal/users/get-by-id"
	CosMainFolder  = "Wegalaxy"
	EmptyDAUserDId = "0x0000000000000000000000000000000000000000"

	UPLOAD_FILE_EXTENSION_PNG UploadFileExtension = ".png"
)

type GetContractInfoReq struct {
	SpaceIds          []string `json:"spaceIds"`
	ContractAddresses []string `json:"contractAddresses"`
	AssetTypes        []string `json:"assetTypes"`
	IsShownInWallet   []bool   `json:"isShownInWallet"`
}

type GetContractInfoRsp struct {
	ErrorCode     int             `json:"errcode"`
	ErrorMessage  string          `json:"errmsg"`
	ContractInfos []*ContractInfo `json:"data"`
}

type ContractInfo struct {
	Id              string `json:"id"`
	ContractAddress string `json:"contractAddress"`
	SpaceId         string `json:"spaceId"`
	BlockchainId    string `json:"blockchainId"`
	ContractType    string `json:"contractType"`
	AssetType       string `json:"assetType"`
	IsShownInWallet bool   `json:"isShownInWallet"`
}

type UploadImageRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	ImageUrl     string `json:"imageUrl"`
	Md5          string `json:"md5"`
}

type DAAdminUserLoginReq struct {
	UserId string `json:"userId"`
}

type DAAdminUserLoginRsp struct {
	ErrorCode     int    `json:"errcode"`
	ErrorMessage  string `json:"errmsg"`
	LongTermToken string `json:"longTermToken"`
}

type ListCustomerTransactionDetailReq struct {
	Did               string                    `json:"did"`
	ContractAddresses []string                  `json:"contractAddress"`
	SceneId           string                    `json:"sceneId"`
	PageInfo          *dahelper.DAPageInfoParam `json:"pageInfo"`
}

type ListCustomerTransactionDetailRsp struct {
	ErrorCode    int                                `json:"errcode"`
	ErrorMessage string                             `json:"errmsg"`
	Data         *ListCustomerTransactionDetailData `json:"data"`
}

type ListCustomerTransactionDetailData struct {
	TransactionDetails []*TransactionDetail  `json:"transactionDetails"`
	PageInfo           *dahelper.DAPageInfo2 `json:"pageInfo"`
}

// 查询用户mdot账户列表
type ListCustomerMdotsReq struct {
	Did                   string               `json:"did"`
	ContractAddresses     []string             `json:"contractAddresses,omitempty"`
	ContractTemplateIds   []string             `json:"contractTemplateIds,omitempty"`
	ContractType          string               `json:"contractType,omitempty"`
	CurrentPath           string               `json:"currentPath,omitempty"`
	MatchChildrenInstance bool                 `json:"matchChildrenInstance,omitempty"`
	PageInfo              *dahelper.DAPageInfo `json:"pageInfo,omitempty"`
	SceneId               string               `json:"sceneId"`
}

type ListCustomerMdotsRsp struct {
	ErrorCode    int                                  `json:"errcode"`
	ErrorMessage string                               `json:"errmsg"`
	Data         *dahelper.DAListCustomerMdotsRspData `json:"data"`
}

// 查询用户mdot交易流水列表
type ListCustomerMdotTransactionDetailReq struct {
	Did                   string               `json:"did"`
	ContractAddresses     []string             `json:"contractAddresses"`
	ContractTemplateIds   []string             `json:"contractTemplateIds"`
	CurrentPath           string               `json:"currentPath"`
	FromOwner             string               `json:"fromOwner"`
	MatchChildrenInstance bool                 `json:"matchChildrenInstance"`
	PageInfo              *dahelper.DAPageInfo `json:"pageInfo"`
	SceneId               string               `json:"sceneId"`
	ToOwner               string               `json:"toOwner"`
}

type ListCustomerMdotTransactionDetailRsp struct {
	ErrorCode    int                                               `json:"errcode"`
	ErrorMessage string                                            `json:"errmsg"`
	Data         *dahelper.DAListCustomerMdotTransactionDetailData `json:"data"`
}

// 空投mdot资产
type BatchIssueMdotsReq struct {
	Did             string                               `json:"did"`
	DAAdminToken    string                               `json:"daAdminToken"`
	BlockchainId    string                               `json:"blockchainId"`
	IssueRule       string                               `json:"issueRule"`
	ContractAddress string                               `json:"contractAddress"`
	Items           []*dahelper.DABatchIssueMdotsReqItem `json:"items"`
}

type BatchIssueMdotsRsp struct {
	ErrorCode    int                     `json:"errcode"`
	ErrorMessage string                  `json:"errmsg"`
	AssetDetail  *dahelper.DAAssetDetail `json:"data"`
}
