package krisbot

import (
	"fmt"

	"github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

func main() {
	response, _ := golambdainvoke.Run(8001, "payload")
	fmt.Println(response)
}
