package routes

import (
	"encoding/json"
	"fmt"
	"laskerbot/controller"
	"laskerbot/models"
	"laskerbot/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessage(server *gin.Engine) {

	var user_message models.Message

	server.POST("/message", func(ctx *gin.Context) {
		err := ctx.ShouldBindJSON(&user_message)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		response := controller.ControllerLLM(user_message.UserMessage)

		responseString := fmt.Sprintf("%v", response[0])

		responseString = services.ParseToJsonResponse(responseString)

		var responseFinal struct {
			LLMResponse struct {
				BotResponse string `json:"bot_response"`
				HmtlResponsa string  `json:"html_response"`
			} `json:"llm_response"`
		}

		err = json.Unmarshal([]byte(responseString), &responseFinal)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Erro ao processar a resposta do LLM",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"llm_response": responseFinal.LLMResponse,
		})
	})
}
