package controller

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/degalaxy/gp-common/gp_error"
	"github.com/degalaxy/gp-common/log"
	"github.com/degalaxy/gp-common/util"
	us_domain "github.com/degalaxy/wegalaxy-service/wegalaxy-foundation/user-space/domain"
	"github.com/degalaxy/wegalaxy-service/wegalaxy-model/app_error"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	// ERROR_CODE_SUCCESS = 0
	// ERROR_MEG_SUCCESS  = "success"

	HTTP_COOKIE_ACCESS_TOKEN = "accessToken"
	HTTP_COOKIE_TOKEN        = "token"

	HTTP_HEADER_AUTHORIZATION = "Authorization"

	DEFAULT_LOCK_TIME = time.Minute * 5

	ISO_DATE_TIME_FORMAT = "2006-01-02T15:04:05Z07:00"

	MAX_PAGE_SIZE = int32(100)

	QUERY_OPERATOR_EQUAL            = "equal"
	QUERY_OPERATOR_NOT_EQUAL        = "notEqual"
	QUERY_OPERATOR_GREATER          = "greater"
	QUERY_OPERATOR_LESS             = "less"
	QUERY_OPERATOR_GREATER_OR_EQUAL = "greaterOrEqual"
	QUERY_OPERATOR_LESS_OR_EQUAL    = "lessOrEqual"

	QUERY_FILTER_VALUE_TYPE_STRING = "string"
	QUERY_FILTER_VALUE_TYPE_FLOAT  = "float"
	QUERY_FILTER_VALUE_TYPE_INT    = "int"
	QUERY_FILTER_VALUE_TYPE_COLUMN = "column"
	QUERY_FILTER_VALUE_TYPE_UUID   = "uuid"

	HTTP_HEADER_ACCESS_TOKEN = "Token"
)

var QUERY_OPERATOR_MAPPING = map[string]string{
	QUERY_OPERATOR_EQUAL:            "=",
	QUERY_OPERATOR_NOT_EQUAL:        "<>",
	QUERY_OPERATOR_GREATER:          ">",
	QUERY_OPERATOR_LESS:             "<",
	QUERY_OPERATOR_GREATER_OR_EQUAL: ">=",
	QUERY_OPERATOR_LESS_OR_EQUAL:    "<=",
}

var QUERY_COL_MAPPING = map[string]map[string]string{}

type Lock struct {
	Key   string
	Value string
}

func GenLockKeyForUpdateUser(loginId string) string {
	return "login|" + loginId + "|act|updateuser"
}

func GenLockKeyForUser(loginId string) string {
	return "login|" + loginId + "|lock"
}

func GenerateGuid() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x%x", b, uint64(time.Now().UnixNano()))
}

func GetUserById(c *ControllerContext, id string) (*us_domain.User, error) {
	// user, err := c.UserRepository.Get(c.DbConnection.DB, id)
	// if err != nil {
	// 	log.Logger.Error(fmt.Sprintf("[ControllerCommon | GetUserById] Cannot retrieve user with id %s : %v", id, err))
	// 	return nil, gp_error.ErrInternalError
	// }
	// if user == nil {
	// 	log.Logger.Error(fmt.Sprintf("[ControllerCommon | GetUserById] Invalid user: id = %s", id))
	// 	return nil, app_error.ErrInvalidUser
	// }
	return &us_domain.User{}, nil
}

func getTokenFromReqHeader(ctx *gin.Context) string {
	authInHeader := ctx.GetHeader("Authorization")
	if len(authInHeader) == 0 {
		log.Infof("[CommonController | getTokenFromReqHeader] No Authorization in header")
		return ""
	}
	tokenParts := strings.Split(authInHeader, " ") // in the format of "Beaer xxxxxx"
	if len(tokenParts) != 2 || len(tokenParts[1]) == 0 {
		log.Infof("[CommonController | getTokenFromReqHeader] No token in Authorization header")
		return ""
	}
	return tokenParts[1]
}

func getTokenFromCookie(ctx *gin.Context) string {
	cookie, err := ctx.Request.Cookie(HTTP_COOKIE_TOKEN)
	if err != nil {
		log.Infof("[CommonController | getTokenFromCookie] Cannot get cookie %s: %v", HTTP_COOKIE_TOKEN, err.Error())
		return ""
	}
	if cookie == nil || cookie.Value == "" {
		log.Infof("[CommonController | getTokenFromCookie] No or blank cookie %s", HTTP_COOKIE_TOKEN)
		return ""
	}
	return cookie.Value
}

// GetUserIdFromToken 检查参数并获取userID
func GetUserIdFromToken(c *ControllerContext, ctx *gin.Context) (string, error) {
	token, err := GetTokenFromContext(ctx)
	if err != nil {
		log.Errorf("[CommonController | GetUserIdFromToken] GetUserIdFromToken failed, err:%v", err)
		return "", err
	}

	claims, err := c.VerifyTokenService.TokenToClaim(token)
	if err != nil {
		log.Errorf("[CommonController | GetUserIdFromToken] TokenToClaim failed, err:%v", err)
		return "", err
	}
	userId := claims.UserId
	if len(userId) == 0 {
		log.Errorf("[ControllerCommon | GetUserIdFromToken] " + app_error.ErrMsgMissingUserId)
		return "", gp_error.ErrInvalidToken
	}
	log.Infof("[CommonController | GetUserIdFromToken] userId = %s", userId)
	return userId, nil
}

// GetDATokenFromToken 检查参数并获取 daToken
func GetDATokenFromToken(c *ControllerContext, ctx *gin.Context) (string, error) {
	token, err := GetTokenFromContext(ctx)
	if err != nil {
		log.Errorf("[CommonController | GetUserIdFromToken] GetUserIdFromToken failed, err:%v", err)
		return "", err
	}

	claims, err := c.VerifyTokenService.TokenToClaim(token)
	if err != nil {
		log.Errorf("[CommonController | GetUserIdFromToken] TokenToClaim failed, err:%v", err)
		return "", err
	}
	daToken := claims.DaToken
	if len(daToken) == 0 {
		log.Errorf("[ControllerCommon | GetUserIdFromToken] " + app_error.ErrMsgMissingUserId)
		return "", gp_error.ErrInvalidToken
	}
	log.Infof("[CommonController | GetUserIdFromToken] daToken = %s", daToken)
	return daToken, nil
}

func UnmarshalInputFromRequestBody(ctx *gin.Context, inputPtr interface{}) error {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("[ControllerCommon | UnmarshalInputFromRequestBody] Error reading request: %v" + err.Error()))
		return gp_error.ErrRequest
	}

	err = json.Unmarshal(body, inputPtr)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("[ControllerCommon | UnmarshalInputFromRequestBody]  Error unmarshalling request %v %v", err, string(body)))
		return gp_error.ErrRequest
	}

	return nil
}

func GetTokenFromContext(ctx *gin.Context) (string, error) {
	token := getTokenFromReqHeader(ctx)
	if len(token) == 0 {
		token = getTokenFromCookie(ctx)
	}
	if len(token) == 0 {
		log.Errorf("[ControllerCommon | GetTokenFromContext] No token in header or cookie")
		return "", gp_error.ErrInvalidToken
	}
	return token, nil
}

func SendRequestAndParseResponse(reqPtr, respPtr interface{}, method, url string, headers map[string]string) error {
	// log.Infof("[ControllerCommon | SendRequestAndParseResponse] sending request to %s", url)
	var reqBodyReader io.Reader = nil
	if reqPtr != nil {
		reqBodyBytes, err := json.Marshal(reqPtr)
		if err != nil {
			log.Errorf("[ControllerCommon | SendRequestAndParseResponse] marshal input failed, err: %v", err)
			return gp_error.GPErrorf(gp_error.ErrInvalidInput, err.Error())
		}
		//log.Infof("[ControllerCommon | SendRequestAndParseResponse] requestBody = %s", string(reqBodyBytes)))
		reqBodyReader = bytes.NewBuffer(reqBodyBytes)
	}

	httpReq, err := http.NewRequest(method, url, reqBodyReader)
	if err != nil {
		log.Errorf("[ControllerCommon | SendRequestAndParseResponse] new http reqeust failed, err: %v", err)
		return gp_error.GPErrorf(gp_error.ErrHttpError, err.Error())
	}

	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}

	client := util.GetHttpClient("")
	httpResp, err := client.Do(httpReq)
	if err != nil {
		log.Errorf("[ControllerCommon | SendRequestAndParseResponse] send http reqeust failed, err: %v", err)
		return gp_error.GPErrorf(gp_error.ErrHttpError, err.Error())
	}
	defer httpResp.Body.Close()
	respBodyBytes, _ := ioutil.ReadAll(httpResp.Body)
	log.Infof("[ControllerCommon | SendRequestAndParseResponse], respBodyBytes: %v", zap.Any("responseBody", string(respBodyBytes)))
	if httpResp.StatusCode != http.StatusOK {
		log.Errorf("[ControllerCommon | SendRequestAndParseResponse] http call return non OK: [%d] %s",
			httpResp.StatusCode, string(respBodyBytes))
		return gp_error.GPErrorf(gp_error.ErrHttpError, fmt.Sprintf("[%d] %s", httpResp.StatusCode, string(respBodyBytes)))
	}

	err = json.Unmarshal(respBodyBytes, respPtr)
	if err != nil {
		log.Errorf("[ControllerCommon | SendRequestAndParseResponse] marshal response failed, err: %v, response: %s",
			err, string(respBodyBytes))
		return gp_error.GPErrorf(gp_error.ErrInvalidResponse, string(respBodyBytes))
	}

	return nil
}

func TrimAndFirstToLowerCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
