package controller

import (
	"log"
	"net/http"

	"github.com/brmcoder/line-bot-dict/service"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

var dictService = service.NewDictionaryService()

func (server *Server) NewWebhookController(router *gin.Engine) {
	line := router.Group("/webhook")
	line.POST("", server.webhookHandler)
}

func (server *Server) webhookHandler(ctx *gin.Context) {
	bot, err := linebot.New(server.config.LineChannelSecret, server.config.LineChannelAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(ctx.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	log.Println(events)

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				word := message.Text
				definition, err := dictService.GetWordDefinition(word)
				if err != nil {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(err.Error())).Do()
					continue
				}
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(definition)).Do()
			}
		}
	}

	ctx.Status(http.StatusOK)
}
