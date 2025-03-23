package yusiqoai

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const apiURL = "https://chat.yusiqo.com/api/chat/completions"

type ChatGPT struct {
	APIKey string
}

type Request struct {
	Model    string  `json:"model"`
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

	if resp.StatusCode != 200 {
		return "", errors.New("ChatGPT API hatası")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var response Response
	json.Unmarshal(body, &response)

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", errors.New("Yanıt alınamadı")
}
