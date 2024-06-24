package user

import (
	"fmt"

	"{{.Name}}/api/user"
	"{{.Name}}/common/code"
	"{{.Name}}/common/request"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var form user.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println(err)
		code.Error(c, code.ParamError, request.GetErrorMsg(form, err))
		return
	}
	code.Success(c, "success")
}
