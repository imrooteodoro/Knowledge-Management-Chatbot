package controller

import (
	"laskerbot/gemini"
	"laskerbot/models"

	"github.com/google/generative-ai-go/genai"
)

func ControllerLLM(user_message string) []genai.Part {

	system_context := models.SystemPrompt()

	ai_response := gemini.Llm(user_message, system_context)

	return ai_response

}
