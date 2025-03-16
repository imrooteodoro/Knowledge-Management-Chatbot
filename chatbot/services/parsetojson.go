package services

import "strings"

func ParseToJsonResponse(responseString string) string {

	responseString = strings.TrimPrefix(responseString, "```json\n")
	responseString = strings.TrimSuffix(responseString, "\n```")
	responseString = strings.ReplaceAll(responseString, "```", "")

	return responseString

}
