package model

type Space struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	EntryPointType string `json:"entryPointType"`
	EntryPointUrl  string `json:"entryPointUrl"`
	ImageUrl       string `json:"imageUrl"`
	CosFolder      string `json:"cosFolder"`
	DisplayOrder   int    `json:"displayOrder"`
	Status         int    `json:"status"`
}

type GetSpaceByIdReq struct {
	SpaceId string `json:"spaceId"`
}
type GetSpacesByNamesReq struct {
	SpaceNames []string `json:"spaceNames"`
}

type GetSpaceRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	Space        *Space `json:"data"`
}

type GetSpacesRsp struct {
	ErrorCode    int      `json:"errcode"`
	ErrorMessage string   `json:"errmsg"`
	Spaces       []*Space `json:"data"`
}

// --- HasAccess 判断当前用户是否有某个Space的权限
type HasAccessReq struct {
	SpaceId string `json:"spaceId"`
}

type HasAccessBySpaceNameReq struct {
	SpaceName string `json:"spaceName"`
}

type HasAccessRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	HasAccess    bool   `json:"data"`
}
