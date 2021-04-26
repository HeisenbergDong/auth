package v1

import (
	"auth/auth/dto"
	"auth/auth/service"
	"auth/auth/vo"
	"auth/gin/response"
	"auth/global"
	"auth/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UpdateAuth(c *gin.Context){
	var updateAuthDto dto.UpdateAuthDto
	_ = c.ShouldBindJSON(&updateAuthDto)
	if err := utils.Verify(updateAuthDto, dto.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateAuth(updateAuthDto.AuthorityId, updateAuthDto.AuthInfos); err != nil {
		global.LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func GetPolicyPathByAuthorityId(c *gin.Context) {
	var updateAuthDto dto.UpdateAuthDto
	_ = c.ShouldBindJSON(&updateAuthDto)
	if err := utils.Verify(updateAuthDto, dto.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	paths := service.GetPolicyPathByAuthorityId(updateAuthDto.AuthorityId)
	response.OkWithDetailed(vo.PolicyPathVo{Paths: paths}, "获取成功", c)
}