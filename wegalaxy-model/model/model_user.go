package model

type UserRole struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Status int    `json:"status"`
}

type UserTag struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type User struct {
	Id     string      `json:"id"`
	Status int         `json:"status"`
	UserId string      `json:"user_id"`
	Did    string      `json:"did"`
	Roles  []*UserRole `json:"roles"`
	Tags   []*UserTag  `json:"tags"`
}

type GetUserInfoRequest struct {
	Channel        string `json:"channel" example:"email"`
	Identifier     string `json:"identifier" example:"tester@test.com"`
	ChallengeId    string `json:"challengeId" example:"550e8400-e29b-41d4-a716-446655440000"`
	ChallengeValue string `json:"challengeValue" example:"ABCD0000"`
}

type UserAccount struct {
	AccountType    string `json:"accountType"`
	BlockchainId   int    `json:"blockchainId"`
	AccountAddress string `json:"accountAddress"`
	Did            string `json:"did"`
}

type GetUserInfoResponseFromDA struct {
	ErrorCode    int            `json:"errcode"`
	ErrorMessage string         `json:"errmsg"`
	Username     string         `json:"username"`
	UserId       string         `json:"userId"`
	UserAccounts []*UserAccount `json:"accounts"`
	Tags         []*UserTag     `json:"tags"`
}

type SpaceAccessCheckFromDAReq struct {
	SpaceId string `json:"space_id"`
}

type SpaceAccessCheckFromDARsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	HasAccess    bool   `json:"hasAccess"`
}

type GetUserInfoFromDAResponse struct {
	ErrorCode    int            `json:"errcode"`
	ErrorMessage string         `json:"errmsg"`
	Username     string         `json:"username"`
	UserId       string         `json:"userId"`
	UserAccounts []*UserAccount `json:"accounts"`
	Tags         []*UserTag     `json:"tags"`
}

type GetUserByIdReq struct {
	UserId string `json:"userId"`
}

type GetUserInfoRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	User         *User  `json:"data"`
}

type BatchGetUserByUserIdReq struct {
	UserIds []string `json:"userId"`
}

type BatchGetUserByUserIdRsp struct {
	ErrorCode    int     `json:"errcode"`
	ErrorMessage string  `json:"errmsg"`
	Users        []*User `json:"data"`
}

// --- UpdateUserRoles 更新用户角色
type UpdateUserRolesReq struct {
	RolesList []*UserRole `json:"rolesList"`
}

type UpdateUserRolesRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

// --- GetUserSpaceByUserTags 根据userTags获取当前用户能参与的space
type GetUserSpaceByUserTagsReq struct {
	TagIds []string `json:"tags"`
}

type GetUserSpaceByUserTagsRsp struct {
	ErrorCode    int      `json:"errcode"`
	ErrorMessage string   `json:"errmsg"`
	SpaceInfo    []*Space `json:"data"`
}

// --- UpdateUserRoleReq 更新用户角色
type UpdateUserRoleReq struct {
	RoleIds []string `json:"roleIds"`
}

type UpdateUserRoleRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	IsSuccessful bool   `json:"data"`
}

// -- GetAllUserRoleRsp
type GetAllUserRoleRsp struct {
	ErrorCode    int         `json:"errcode"`
	ErrorMessage string      `json:"errmsg"`
	UserRoles    []*UserRole `json:"data"`
}

type CheckIfUserHasATagReq struct {
	TagName string `json:"tagName"`
}

type CheckIfUserHasATagRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	HasTag       bool   `json:"data"`
}

type UserExistReq struct {
	UserId string `json:"userId"`
}

type UserExistRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	IsExist      bool   `json:"isExist"`
}
