package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AccessKey struct {
	TelegramAccessKey string `json:"telegram_access_key"`
	VKAcessKey        string `json:"vk_access_key"`
}

func main() {
	accessKeyFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening the config file", err)
		return
	}
	defer accessKeyFile.Close()
	decoder := json.NewDecoder(accessKeyFile)
	var accessKey AccessKey
	if err := decoder.Decode(&accessKey); err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return
	}
	type PhotoSize struct {
		Type   string `json:"type"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Url    string `json:"url"`
	}
	type Attachment struct {
		Type  string `json:"type"`
		Photo struct {
			Sizes []PhotoSize `json:"sizes"`
		} `json:"photo"`
	}

	type Item struct {
		ID          int          `json:"id"`
		Text        string       `json:"text"`
		Attachments []Attachment `json:"attachments"`
	}

	type Response struct {
		Count int    `json:"count"`
		Items []Item `json:"items"`
	}

	type VKResponse struct {
		Response Response `json:"response"`
	}
	VkToken := accessKey.VKAcessKey
	groupDomain := "https://vk.com/a_day_in_life"
	ownerId := "-219631272"
	filter := "all"
	count := 1
	apiURL := fmt.Sprintf("https://api.vk.com/method/wall.get?owner_id=%s&domain=%s&count=%d&filter=%s&access_token=%s&v=5.131", ownerId, groupDomain, count, filter, VkToken)
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
	var vkresponse1 VKResponse
	err = json.Unmarshal(jsonData, &vkresponse1)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("VK API RESPONSE: ")

	fmt.Printf("Total Posts: %d\n", vkresponse1.Response.Count)
	for _, post := range vkresponse1.Response.Items {
		fmt.Printf("Post ID: %d\n", post.ID)
		fmt.Printf("Post Text: %s\n", post.Text)
		for _, attachment := range post.Attachments {
			if attachment.Type == "photo" {
				fmt.Println("Photo URLs:")
				for _, size := range attachment.Photo.Sizes {
					fmt.Printf("Size: %s, Width: %d, Height: %d, URL: %s\n", size.Type, size.Width, size.Height, size.Url)
				}
			}
		}
		fmt.Println("----------------------")

		telegramTOken := accessKey.TelegramAccessKey
		telegramURL := "https://api.telegram.org/bot" + telegramTOken + "/"
		chatID := "@test_api_group"
		contentType := "application/json"

		for _, post := range vkresponse1.Response.Items {
			message := fmt.Sprintf("Post ID: %d\nPost Text: %s", post.ID, post.Text) //formatting txt

			for _, attachment := range post.Attachments {
				if attachment.Type == "photo" {
					message += "\nPhoto URLs:\n"
					for _, size := range attachment.Photo.Sizes {
						message += fmt.Sprintf("Size: %s, Width: %d, Height: %d, URL: %s\n", size.Type, size.Width, size.Height, size.Url)
					}
				}
			}

			MessageData := struct {
				ChatID string `json:"chat_id"`
				Text   string `json:"text"`
			}{
				ChatID: chatID,
				Text:   message,
			}

			payload, err := json.Marshal(MessageData)
			if err != nil {
				fmt.Println("Error creating the JSON payload")
				return
			}

			telegramResponse, err := http.Post(telegramURL+"sendMessage", contentType, bytes.NewBuffer(payload))//sending txt
			if err != nil {
				fmt.Println("Error creating a POST request to Telegram:", err)
				return
			}
			defer telegramResponse.Body.Close()

			responseBody, err := io.ReadAll(telegramResponse.Body)
			if err != nil {
				fmt.Println("Error reading Telegram response:", err)
				return
			}

			fmt.Println("Telegram API Response:")
			fmt.Println(string(responseBody))

			fmt.Println("----------------------")
		}
	}
}

