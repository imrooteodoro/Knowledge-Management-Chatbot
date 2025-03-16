package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botToken := ""
	if botToken == "" {
		log.Fatal("Erro: TELEGRAM_BOT_TOKEN não definido")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		for update := range updates {
			if update.Message == nil {
				continue
			}

			userMessage := update.Message.Text
			chatID := update.Message.Chat.ID

			apiResponse := callAPI(userMessage)

			botResponse, ok := apiResponse["bot_response"].(string)
			if !ok {
				botResponse = "Erro: Resposta do bot não encontrada."
			}

			msg := tgbotapi.NewMessage(chatID, botResponse)

			htmlResponse, _ := apiResponse["html_response"].(string)
			if htmlResponse != "" {
				msg.ParseMode = "HTML"
				msg.Text = htmlResponse
			}

			_, err := bot.Send(msg)
			if err != nil {
				log.Println("Erro ao enviar mensagem:", err)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "Bot está funcionando",
		})
	})

	port := "8080"
	fmt.Printf("Iniciando servidor HTTP na porta %s...\n", port)
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor HTTP:", err)
	}
}

func callAPI(userMessage string) map[string]interface{} {
	apiURL := ""

	jsonData := map[string]string{"user_message": userMessage}
	jsonValue, _ := json.Marshal(jsonData)

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Println("Erro ao chamar API:", err)
		return map[string]interface{}{"bot_response": "Erro ao processar sua mensagem."}
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println("Erro ao decodificar resposta da API:", err)
		return map[string]interface{}{"bot_response": "Erro ao processar sua mensagem."}
	}

	llmResponse, exists := result["llm_response"].(map[string]interface{})
	if !exists {
		log.Println("Erro: llm_response não encontrado")
		return map[string]interface{}{"bot_response": "Erro: Resposta inválida."}
	}

	botResponse, exists := llmResponse["bot_response"].(string)
	if !exists {
		log.Println("Erro: bot_response não encontrado")
		return map[string]interface{}{"bot_response": "Erro: Resposta inválida."}
	}

	htmlResponse, _ := llmResponse["html_response"].(string)

	return map[string]interface{}{
		"bot_response":  botResponse,
		"html_response": htmlResponse,
	}
}
