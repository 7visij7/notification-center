package model

import (
	"fmt"
	"strings"
    "net/http"
	"git.puls.ru/devops1/sre/devops-api/types"
	"github.com/gin-gonic/gin"
)

func AlertRedash(ctx *gin.Context, chat int, token string) {
	var alert types.Redash
	if err := ctx.ShouldBindJSON(&alert); err != nil {
		fmt.Println(err)
		NewError(ctx, http.StatusBadRequest, err)
		return
	
	}
	sendText := fmt.Sprintf("<b>[%s]</b> %s\nURL: %s/queries/%d\n%s %s %d",strings.ToTitle(alert.RedashAlerts.State), alert.RedashAlerts.Name, alert.Url_base, alert.RedashAlerts.Query_id,alert.RedashAlerts.Option.Column, alert.RedashAlerts.Option.Op, alert.RedashAlerts.Option.Value) 
	err := TelegramSend(sendText, chat, token)
    if err != nil {
		fmt.Println(err)
        NewError(ctx, http.StatusInternalServerError, err)
        return
    }
	ctx.JSON(http.StatusOK, "OK")
}