package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetUpdates(token string, offset int) ([]Update, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d", token, offset)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result struct {
		OK     bool     `json:"ok"`
		Result []Update `json:"result"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Result, nil
}

func SendMessage(token string, message Message) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	payload := fmt.Sprintf("chat_id=%d&text=%s", message.Chat, message.Text)

	resp, err := http.Post(
		url,
		"application/json",
		strings.NewReader(payload),
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram error: %s", string(body))
	}

	return nil
}
