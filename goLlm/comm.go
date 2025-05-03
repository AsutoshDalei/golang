package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Model       string    `json:"model"`
	CreatedTime time.Time `json:"created_at`
	Message     Message   `json:"message"`
	Done        bool      `json:"done"`

	TotalDuration      int64 `json:"total_duration"`
	LoadDuration       int   `json:"load_duration"`
	PromptEvalCount    int   `json:"prompt_eval_count"`
	PromptEvalDuration int   `json:"prompt_eval_duration"`
	EvalCount          int   `json:"eval_count"`
	EvalDuration       int   `json:"eval_duration"`
}

func talk(url string, convo Request) (*Response, error) {
	jsn, err := json.Marshal(&convo)

	if err != nil {
		return nil, err
	}
	client := http.Client{}

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsn))

	if err != nil {
		return nil, err
	}

	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()

	ollamaResp := Response{}

	err = json.NewDecoder(httpResp.Body).Decode(&ollamaResp)
	return &ollamaResp, err

}

const ollamaURL = "http://localhost:11434/api/chat"

func main() {
	reader := bufio.NewReader(os.Stdin)

	var memory []Message

	fmt.Println("Chat started.")

	for {
		fmt.Print("You: ")
		inp, _ := reader.ReadString('\n')
		inp = strings.TrimSpace(inp)

		if inp == "exit" {
			break
		}

		usrMsg := &Message{
			Role:    "user",
			Content: inp,
		}
		memory = append(memory, *usrMsg)

		req := &Request{
			Model:    "llama3.2:1b",
			Stream:   false,
			Messages: memory,
		}

		resp, err := talk(ollamaURL, *req)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		memory = append(memory, resp.Message)
		fmt.Printf("Bot: %s\n", resp.Message.Content)

	}

}
