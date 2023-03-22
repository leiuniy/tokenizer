package tiktokens

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"testing"
)

func TestEncodingForModel(t *testing.T) {
	fmt.Println(os.TempDir())
	encoder, err := EncodingForModel("gpt-3.5-turbo")
	if err != nil {
		log.Fatal(err)
	}
	str := "tiktoken is great!"
	encoded := encoder.Encode(str, nil, nil)
	fmt.Println("We can look at each token and what it represents:")
	for _, token := range encoded {
		fmt.Printf("%d -- %s\n", token, encoder.Decode([]int{token}))
	}
	decoded := encoder.Decode(encoded)
	fmt.Printf("We can decode it back into: %s\n", decoded)

	messages := []openai.ChatCompletionMessage{
		{Role: "system", Content: "你是一个幽默风趣、见多识广、温柔可爱的小助手"},
		{Role: "user", Content: "你叫什么名字，可以告诉我吗？"},
		{Role: "assistant", Content: "当然可以，我的名字是小智。"},
		{Role: "user", Content: "你好，小智"},
		{Role: "assistant", Content: "你好，有什么可以帮助你的吗？"},
		{Role: "user", Content: "你多大了"},
		{Role: "assistant", Content: "作为一个聊天机器人，我没有实际的年龄，我只是一个程序。"},
		{Role: "user", Content: "你长什么样子"},
		{Role: "assistant", Content: "我没有具体的外貌，因为我只是一个聊天机器人程序，没有实际的形体。"},
	}

	// every message follows <im_start>{role/name}\n{content}<im_end>
	// every reply is primed with <im_start>assistant
	numTokens := len(messages)*4 + 2
	for _, message := range messages {
		numTokens += len(encoder.Encode(message.Role, nil, nil))
		numTokens += len(encoder.Encode(message.Content, nil, nil))
		if message.Name != "" {
			numTokens += len(encoder.Encode(message.Name, nil, nil))
			numTokens -= 1
		}
	}
	fmt.Println("token num:", numTokens)
}
