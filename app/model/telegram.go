package model

import (
	"fmt"
	"bytes"
	"net/http"
	"github.com/gin-gonic/gin"
	"git.puls.ru/devops1/sre/devops-api/types"
)

func MessageTelegram(ctx *gin.Context, chat int, token string){
    var message types.Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		fmt.Println(err)
		NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err := TelegramSend(message.Text, chat, token)
    if err != nil {
		fmt.Println(err)
        NewError(ctx, http.StatusInternalServerError, err)
        return
    }
	ctx.JSON(http.StatusOK, "OK")

}

func TelegramSend(text string, chat int, token string) (error) {
	data := []byte(fmt.Sprintf(`{"chat_id":%d, "text":"%s", "parse_mode":"HTML"}`, chat, text))
	tx := bytes.NewReader(data)
	resp, err := http.Post(fmt.Sprintf("%s%s/sendMessage", types.TelegramURL+"/bot", token), "application/json", tx)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return err
	  }
	return nil
}
