package app_error

import (
	"net/http"

	"github.com/degalaxy/gp-common/gp_error"
)

var (
	ErrNoRowsUpdated         = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_NO_ROWS_UPDATED), Message: "no rows updated: %v"}
	ErrDataNotFound          = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_DATA_NOT_FOUND), Message: "data not found: %v"}
	ErrHitLimit              = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_HIT_LIMIT), Message: "hit %s limit: %d"}
	ErrNoCouponLeft          = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_NO_COUPON_LEFT), Message: "no coupon left"}
	ErrHasBeenRedeemed       = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_HAS_BEEN_REDEEMED), Message: "has been redeemed"}
	ErrEmptyGift             = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_EMPTY_GIFT), Message: "empty gift"}
	ErrNotOwner              = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_NOT_OWNER), Message: "not owner"}
	ErrEnoughMedals          = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_NOT_ENOUGH_MEDALS), Message: "not enough medals"}
	ErrAccessDenied          = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_ACCESS_DENIED), Message: "access denied"}
	ErrInvalidTimePeriod     = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_INVALID_TIME_PERIOD), Message: "invalid time period: %s"}
	ErrNotSupportedOperation = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_NOT_SUPPORTED_OPERATION), Message: "not supported operation: %s"}

	// user error
	ErrUserAlreadyExists = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_USER_ALREADY_EXISTS), Message: "user already exists"}
	ErrUserNotExist      = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_USER_NOT_EXIST), Message: "user not exist"}
	ErrInvalidUser       = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_INVALID_USER), Message: "invalid user"}
	ErrAuthServiceError  = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_AUTHSERVICE_ERROR), Message: "error in Auth Service: [%d][%s]"}
	ErrMissingUserID     = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_MISSING_USER_ID), Message: "missing user ID"}

	// qrcode error
	ErrInvalidQrCode  = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_INVALID_QRCODE), Message: "invalid qr code: %v"}
	ErrRemainingUsage = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_REMAINING_USAGE), Message: "no remaining usage"}
	ErrNoHandler      = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_NO_HANDLER), Message: "no handler for %v"}

	// report error
	ErrReportFilePathInvalid = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_REPORT_FILE_PATH_INVALID), Message: "invalid filePath: %v"}
	ErrZipHandlerFailed      = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_ZIP_HANDLER_FAILED), Message: "zip report file failed: %v"}
	ErrSignNotMatch          = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_SIGN_NOT_MATCH), Message: "sign not match: %v"}

	// assets error
	ErrGetLongTermTokenFailed    = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_GET_LONG_TERM_TOKEN_FAILED), Message: "get da long term token failed: %v"}
	ErrGetDAUserInfoFailed       = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_GET_DA_USER_INFO_FAILED), Message: "get da user info failed: %v"}
	ErrGetContractInfoFailed     = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_GET_CONTRACT_INFO_FAILED), Message: "get contract info failed: %v"}
	ErrUploadCosPicFailed        = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_UPLOAD_COS_PIC_FAILED), Message: "upload cos pic failed: %v"}
	ErrGetUserDidFailed          = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_GET_USER_DID_FAILED), Message: "get user did failed: %v"}
	ErrGetDAUserIdFailed         = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_GET_DA_USER_ID_FAILED), Message: "get da user id failed: %v"}
	ErrCombinedIssueFailed       = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_COMBINED_ISSUE_FAILED), Message: "combine issue failed: %v"}
	ErrCallRanksContestantFailed = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_CALL_RANKS_CONTESTANT_FAILED), Message: "call ranks contestant failed: %v"}

	// egg error
	ErrLockUserFailed = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_EGG_LOCK_USER_FAILED), Message: "lock user failed: %s"}
	ErrLockEggFailed  = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_EGG_LOCK_EGG_FAILED), Message: "lock egg failed: %s"}

	// rabbit error
	ErrGetAssetDetailFailed        = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_GET_ASSET_DETAIL_FAILED), Message: "get asset detail failed: %s"}
	ErrCombineIssueFailed          = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_GET_COMBINED_ISSUE_FAILED), Message: "combine issue failed: %s"}
	ErrPreUploadCosFilesFailed     = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_PRE_UPLOAD_COS_FILES_FAILED), Message: "pre upload files to cos failed: %s"}
	ErrLockedCouponFound           = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_LOCKED_COUPON_FOUND), Message: "found %d locked coupon"}
	ErrInvalidImageUrl             = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_INVALID_IMAGE_URL), Message: "invalid image url: %s"}
	ErrEmptyImageInfoFromAppConfig = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_EMPTY_IMAGE_INFO_FROM_APP_CONFIG), Message: "get empty image info from app config"}
	ErrGetRandomUpgradePartsFailed = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_GET_RANDOM_UPGRADE_PARTS_FAILED), Message: "get random upgrade parts failed: %s"}

	ErrTooManyRequests = gp_error.GPError{Status: http.StatusOK, Code: int(RetCode_ERR_TOO_MANY_REQUESTS), Message: "%s"}
)

const (
	ErrorMsgSuccess               = "success"
	ErrMsgMissingUserId           = "Missing user id"
	ErrMsgQrCodeIdAndNameNotMatch = "id and name not match"
	ErrMsgQrCodeInactive          = "inactive"
	ErrMsgQrCodeBeforeValidFrom   = "curren time before valid from"
	ErrMsgQrCodeAfterValidTo      = "curren time after valid to"
)
