package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shuklarituparn/echo/Models"
	"github.com/shuklarituparn/echo/keys"
	"io"
	"net/http"
	"os"
)

func main() {
	keys.LoadEnvVariables(".env")
	apiURL := fmt.Sprintf("https://api.vk.com/method/wall.get?owner_id=%d&domain=%s&count=%d&filter=%s&access_token=%s&v=5.131", keys.OwnerId, keys.GroupDomain, keys.Count, keys.Filter, keys.VkToken)
	vkresponse, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error getting the data from the page")
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(vkresponse.Body)

	responseFromVk, err := io.ReadAll(vkresponse.Body)
	if err != nil {
		fmt.Println("Error reading VK response")
		return
	}
	jsonData := responseFromVk
	var vkresponse1 Models.VKResponse
	err = json.Unmarshal(jsonData, &vkresponse1)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("VK API RESPONSE: ")
	fmt.Printf("Total Posts: %d\n", vkresponse1.Response.Count)
	for _, post := range vkresponse1.Response.Items {
		message := fmt.Sprintf("%s", post.Text)
		sendPhoto := Models.SendPhoto{ChatID: keys.TelegramChatId}
		isPhoto := false
		for _, attachment := range post.Attachments {
			if attachment.Type == "photo" {
				for _, size := range attachment.Photo.Sizes {
					if size.Type == "r" {
						sendPhoto.Photo = size.URL
						sendPhoto.Caption = post.Text
						isPhoto = true

					}
				}

			}
		}

		telegramToken, _ := os.LookupEnv("TELEGRAM_ACCESS_KEY")
		telegramURL := "https://api.telegram.org/bot" + telegramToken + "/"
		contentType := "application/json"

		NewMessageTelegram := Models.MessageData{
			ChatID: keys.TelegramChatId,
			Text:   message,
		}
		payload1, err := json.Marshal(NewMessageTelegram)
		if err != nil {
			fmt.Println("Error creating the JSON payload")
			return
		}
		payload2, err := json.Marshal(sendPhoto)
		if err != nil {
			fmt.Println("Error creating the send Image text")
			return
		}

		if !isPhoto {
			telegramResponse, err := http.Post(telegramURL+"sendMessage", contentType, bytes.NewBuffer(payload1))
			if err != nil {
				fmt.Println("Error creating a POST request to Telegram:", err)
				return
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(telegramResponse.Body)

		} else {
			telegramResponse, err := http.Post(telegramURL+"sendPhoto", contentType, bytes.NewBuffer(payload2))
			if err != nil {
				fmt.Println("Error creating a POST request to Telegram:", err)
				return
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(telegramResponse.Body)

			fmt.Println("Sent the posts successfully to the telegram!!")

		}

		fmt.Println("----------------------")
	}

}
