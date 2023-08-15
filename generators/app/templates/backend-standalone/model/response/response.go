package response

import (
	"<%= displayName %>/global"
	"<%= displayName %>/utils/e"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"fmt"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Warning struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Result(code int, data interface{}, msg string, ctx *gin.Context) {
	// 开始时间
	ctx.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(e.SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(e.SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(e.SUCCESS, data, "操作成功", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(e.SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(e.ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(e.ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailedMessage(message string, c *gin.Context) {
	Result(e.MessagedError, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

func WsWithMessage(ws *websocket.Conn, message string) {
	data := fmt.Sprintf("{\"type\":\"msg\",\"message\":\"%s\"}", message)
	err := ws.WriteMessage(1, []byte(data))
	if err != nil {
		global.LOG.Error(err.Error())
	}
}

func WsWithError(ws *websocket.Conn, message string) {
	data := fmt.Sprintf("{\"type\":\"err\",\"message\":\"%s\"}", message)
	err := ws.WriteMessage(1, []byte(data))
	if err != nil {
		global.LOG.Error(err.Error())
	}
}

func WsWithType(ws *websocket.Conn, ty string, message string) {
	data := fmt.Sprintf("{\"type\":\"%s\",\"message\":\"%s\"}", ty, message)
	err := ws.WriteMessage(1, []byte(data))
	if err != nil {
		global.LOG.Error(err.Error())
	}
}
