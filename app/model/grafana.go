package model

import (
	"fmt"
    "net/http"
	"git.puls.ru/devops1/sre/devops-api/types"
	"github.com/gin-gonic/gin"
)

func AlertGrafana(ctx *gin.Context ,chat int, token string) {
	var alert types.GrafanaAlert
	if err := ctx.ShouldBindJSON(&alert); err != nil {
		fmt.Println(err)
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	sendText := fmt.Sprintf("<b>%s</b>\nURL: %s\n",alert.Title,alert.RuleUrl)
	if len(alert.EvalMatches) > 0 {
		sendText = sendText + fmt.Sprintf("Metrics:\n %s: %.2f", alert.EvalMatches[0].Metric, alert.EvalMatches[0].Value)
	}

	err := TelegramSend(sendText, chat, token)
    if err != nil {
		fmt.Println(err)
        NewError(ctx, http.StatusInternalServerError, err)
        return
    }
	ctx.JSON(http.StatusOK, "OK")
}