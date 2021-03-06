package main

import (
	"github.com/spaceuptech/space-api-go"
	"fmt"
	"time"
	"math/rand"
)

func main() {
	api, err := api.New("books-app", "localhost:8081", false)
	if(err != nil) {
		fmt.Println(err)
	}
	v := map[string]interface{}{"params":"params"}
	for {
		resp, err := api.Call("service", "echo_func", v, 5000)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			if resp.Status == 200 {
				var v map[string]interface{}
				err:= resp.Unmarshal(&v)
				if err != nil {
					fmt.Println("Error Unmarshalling:", err)
				} else {
					fmt.Println("Result:", v)
				}
			} else {
				fmt.Println("Error Processing Request:", resp.Error)
			}
		}
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(6))*time.Second)
	}
}
