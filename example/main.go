package main

import (
	"context"
	"fmt"
	"prompts"
	"strings"

	gogpt "github.com/sashabaranov/go-gpt3"
)

const (
	CHATGPTAPITOKEN = ""
)

var his []gogpt.ChatCompletionMessage

func main() {
	order := prompts.New()
	order.Type = prompts.ActorComposer

	actorName, prompt := order.Prompt(true)
	fmt.Println(actorName)

	gpt := gogpt.NewClient(CHATGPTAPITOKEN)
	req := gogpt.ChatCompletionRequest{
		Model: gogpt.GPT3Dot5Turbo,
		Messages: append(his, gogpt.ChatCompletionMessage{
			Role:    "user",
			Content: prompt,
		}),
	}

	res, err := gpt.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return err.Error()
	}

	fmt.Println(res)
}

func load() {
	acts := prompts.Load()

	var s string

	for i := 0; i < len(acts); i++ {
		s = strings.ReplaceAll(acts[i].Actor, " ", "")
		fmt.Println(s)
	}
}
