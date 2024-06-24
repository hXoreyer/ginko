package code

import (
	_ "embed"

	"{{.Name}}/global"

	"github.com/gin-gonic/gin"
)

//go:embed code.go
var ByteCodeFile []byte

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	ServerError = 10101

	ParamError = 20101
)

func Text(code int) string {
	lang := global.App.Config.App.Lang

	if lang == "zh-cn" {
		return zhCNText[code]
	}

	if lang == "en-us" {
		return enUSText[code]
	}

	return zhCNText[code]
}

func Success(c *gin.Context, message string, data ...any) {
	if len(data) > 0 {
		c.JSON(200, gin.H{
			"code":    0,
			"message": message,
			"data":    data,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    0,
			"message": message,
		})
	}
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"error":   Text(code),
	})
}
