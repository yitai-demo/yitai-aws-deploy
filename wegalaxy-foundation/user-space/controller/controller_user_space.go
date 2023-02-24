package controller

import (
	"net/http"

	"github.com/degalaxy/gp-common/gp_error"
	"github.com/degalaxy/gp-common/log"
	common_ctrl "github.com/degalaxy/wegalaxy-service/wegalaxy-foundation/common/controller"
	"github.com/degalaxy/wegalaxy-service/wegalaxy-model/model"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	controllerContext *common_ctrl.ControllerContext
}

func NewUserController(controllerContext *common_ctrl.ControllerContext) (*UserController, error) {
	return &UserController{
		controllerContext: controllerContext,
	}, nil
}

// @Summary      Get user information
// @Description  Get user information by ID
// @ID           get-user
// @Tags         user
// @Param        Authorization header string  false  "Token passed in Authorization in 'Bearer xxxx' format. The alternative is to pass 'token=xxxx' in cookie"
// @Produce      json
// @Success      200 {object} model.GetUserInfoRsp
// @Router       /users [get]
func (c *UserController) GetUser(ctx *gin.Context) {
	userId, err := common_ctrl.GetUserIdFromToken(c.controllerContext, ctx)
	if err != nil {
		log.Errorf("[UserController | GetUser] get userId failed, err: %v", err)
		gp_error.RaiseHttpError(ctx, err)
		return
	}

	// Get user
	user, err := common_ctrl.GetUserById(c.controllerContext, userId)
	if err != nil {
		log.Errorf("[UserController | GetUser] Failed to get user, err: %v", err)
		gp_error.RaiseHttpError(ctx, err)
		return
	}

	rsp := &model.GetUserInfoRsp{
		User: &model.User{
			Id:     userId,
			UserId: user.UserId,
			Did:    user.Did,
			Status: int(user.Status),
		},
	}
	ctx.JSON(http.StatusOK, rsp)
}
