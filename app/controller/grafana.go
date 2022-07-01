package controller

import (
    "git.puls.ru/devops1/sre/devops-api/model"
	"git.puls.ru/devops1/sre/devops-api/types"
	"github.com/gin-gonic/gin"
    _ "git.puls.ru/devops1/sre/devops-api/docs" // docs is generated by Swag CLI, you have to import it.
  )

// @Summary Send message to "Ecomm" telegram chat
// @Description Send message to telegram chat
// @Tags grafana
// @Accept json
// @Produce json
// @Param message body types.Alert true "Send message"
// @Success 200 {array} types.Message
// @Failure 400,404 {array} model.HTTPError
// @Failure 500 {array} model.HTTPError
// @Security ApiKeyAuth
// @Router /grafana/ecomm [post]
func (c *Controller) GrafanaEcomm(ctx *gin.Context) {
	model.AlertGrafana(ctx, types.ChatIdEcomm, types.BotTokenDevops)
}

// @Summary Send message to "PULS Markirovka Alerts" telegram chat
// @Description Send message "PULS Markirovka Alerts" to telegram chat
// @Tags grafana
// @Accept json
// @Produce json
// @Param message body types.Alert true "Send message"
// @Success 200 {array} types.Message
// @Failure 400,404 {array} model.HTTPError
// @Failure 500 {array} model.HTTPError
// @Security ApiKeyAuth
// @Router /grafana/markirovka [post]
func (c *Controller) GrafanaMarkirovka(ctx *gin.Context) {
	model.AlertGrafana(ctx, types.ChatIdMarkirovka, types.BotTokenMarkirovka)
}

// @Summary Send message to "PULS DevOps Monitor" telegram chat
// @Description Send message "PULS DevOps Monitor" to telegram chat
// @Tags grafana
// @Accept json
// @Produce json
// @Param message body types.Alert true "Send message"
// @Success 200 {array} types.Message
// @Failure 400,404 {array} model.HTTPError
// @Failure 500 {array} model.HTTPError
// @Security ApiKeyAuth
// @Router /grafana/devops [post]
func (c *Controller) GrafanaDevops(ctx *gin.Context) {
	model.AlertGrafana(ctx, types.ChatIdDevops, types.BotTokenDevops)
}

// @Summary Send message to "GRVS_support_alerts" telegram chat
// @Description Send message "GRVS_support_alerts" to telegram chat
// @Tags grafana
// @Accept json
// @Produce json
// @Param message body types.Alert true "Send message"
// @Success 200 {array} types.Message
// @Failure 400,404 {array} model.HTTPError
// @Failure 500 {array} model.HTTPError
// @Security ApiKeyAuth
// @Router /grafana/ovs-support [post]
func (c *Controller) GrafanaOVSSupport(ctx *gin.Context) {
	model.AlertGrafana(ctx, types.ChatIdOVSSupport, types.BotTokenDevops)
}

// @Summary Send message to "kt_metrics_alert" telegram chat
// @Description Send message "kt_metrics_alert" to telegram chat
// @Tags grafana
// @Accept json
// @Produce json
// @Param message body types.Alert true "Send message"
// @Success 200 {array} types.Message
// @Failure 400,404 {array} model.HTTPError
// @Failure 500 {array} model.HTTPError
// @Security ApiKeyAuth
// @Router /grafana/markirovka-ktteam [post]
func (c *Controller) GrafanaMarkirovkaKTTeam(ctx *gin.Context) {
	model.AlertGrafana(ctx, types.ChatIdMarkirovkaKTTeam, types.BotTokenDevops)
}

// @Summary Send message to "DBA_channel" telegram chat
// @Description Send message "DBA_channel" to telegram chat
// @Tags grafana
// @Accept json
// @Produce json
// @Param message body types.Alert true "Send message"
// @Success 200 {array} types.Message
// @Failure 400,404 {array} model.HTTPError
// @Failure 500 {array} model.HTTPError
// @Security ApiKeyAuth
// @Router /grafana/ovs-dba [post]
func (c *Controller) GrafanaOVSDBA(ctx *gin.Context) {
	model.AlertGrafana(ctx, types.ChatIdOVSDBA, types.BotTokenDevops)

}