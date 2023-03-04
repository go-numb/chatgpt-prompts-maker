# chatgpt-prompts-maker

This package creates instructions that give CHATGPT various roles.


## Usage

```go 
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
	acts := prompts.New()

    // set actor and create prompt
	isJapanse = false
    acts.Type = prompts.ActorComposer
	actorName, prompt := order.Prompt(isJapanse)
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
		log.Fatal(err)
	}

	fmt.Println(res)
}

```


## Author

[@_numbP](https://twitter.com/_numbP)