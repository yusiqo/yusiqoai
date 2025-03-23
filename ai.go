package yusiqoai


import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const apiURL = "https://chat.yusiqo.com/api/chat/completions"

type ChatGPT struct {
	APIKey string
}


type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

func (c *ChatGPT) SendMessage(prompt string) (string, error) {
	reqBody, _ := json.Marshal(Request{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "system", Content: "Sen bir yardımcı AI'sın."},
			{Role: "user", Content: prompt},
		},
	})

	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return "", readErr
	}

	if resp.StatusCode != http.StatusOK {
		// Hata mesajını API'den al ve göster
		var apiError Response
		json.Unmarshal(body, &apiError)

		if apiError.Error.Message != "" {
			return "", errors.New(fmt.Sprintf("ChatGPT API hatası: %s", apiError.Error.Message))
		}

		return "", errors.New(fmt.Sprintf("ChatGPT API hatası: HTTP %d - %s", resp.StatusCode, string(body)))
	}

	var response Response
	json.Unmarshal(body, &response)

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", errors.New("Yanıt alınamadı")
}
