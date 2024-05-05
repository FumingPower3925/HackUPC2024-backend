package ai

import (
	"context"

	oai "github.com/sashabaranov/go-openai"
)

var client *oai.Client

func init() {
	client = oai.NewClient("")
}

func GetResponse(request string) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		oai.ChatCompletionRequest{
			Model: oai.GPT3Dot5Turbo,
			Messages: []oai.ChatCompletionMessage{
				{
					Role:    oai.ChatMessageRoleUser,
					Content: "You are a helpful guide for blind people.  The blind person will ask you for one of this three options: { REST, BATHROOM, FLIGHT }. Your job will be to output ONLY one of this three options. Do NOT include explanations of your reasoning and only include one of those three words. BLIND PERSON: " + request,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func Speech2Text(path string) (string, error) {
	ctx := context.Background()

	req := oai.AudioRequest{
		Model:    oai.Whisper1,
		FilePath: path,
	}
	resp, err := client.CreateTranscription(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Text, nil
}
