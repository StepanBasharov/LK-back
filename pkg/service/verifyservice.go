package service

import (
	"LK_back/settings"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type RequestMTS struct {
	Number             string `json:"number"`
	Destination        string `json:"destination"`
	Text               string `json:"text"`
	TemplateResourceId uint64 `json:"template_resource_id"`
}

func GenerateVerifySmsText() string {
	rand.Seed(time.Now().UnixNano())
	start := 1000
	end := 9999
	code := rand.Intn((end - start + 1) + start)
	text := fmt.Sprintf("Ваш код верефикации: %d", code)

	return text
}

func SendSmsService(phone string) string {
	config := settings.InitConfigService()
	phoneService := config.PhoneForAuth
	apiKey := config.MtsApiKey
	mtsUrl := "https://api.exolve.ru/messaging/v1/SendSMS"

	smsRequest := RequestMTS{
		Number:             phone,
		Destination:        phoneService,
		Text:               GenerateVerifySmsText(),
		TemplateResourceId: rand.Uint64(),
	}
	data, errJson := json.Marshal(smsRequest)
	if errJson != nil {
		return fmt.Sprintf("Не удалось SMS отправить код на номер %s", phone)
	}
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, mtsUrl, bytes.NewBuffer(data))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	_, errReq := client.Do(req)
	if errReq != nil {
		return fmt.Sprintf("Не удалось SMS отправить код на номер %s", phone)
	}
	return fmt.Sprintf("SMS код отправлен на номер %s", phone)
}
