package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Resp{Code: 0, Msg: "成功", Data: data})
}

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, Resp{Code: -1, Msg: msg})
}
