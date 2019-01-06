package main

import (
	"fmt"
	"net/url"

	"github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

func main() {
	u, err := url.Parse("https://api.pagerduty.com/oncalls")
	if err != nil {
		fmt.Println(err)
	}

	v := url.Values{}

	v.Set("includes[]", "users")
	v.Set("time_zone", "UTC")
	v.Set("schedule_id", "3")
	u.RawQuery = v.Encode()
	fmt.Println(u.String())

	response, _ := golambdainvoke.Run(8001, "payload")
	fmt.Println(string(response))
}
