package main

import (
	"github.com/spaceuptech/space-api-go"
	"github.com/spaceuptech/space-api-go/api/model"
	"fmt"
	"time"
	"math/rand"
)

func main() {
	api, err := api.New("books-app", "localhost:8081", false)
	if(err != nil) {
		fmt.Println(err)
	}
	db := api.Mongo()
	for {
		unsubscribe := db.LiveQuery("books").Subscribe(func(liveData *model.LiveData, changeType string) () {
			var v []interface{}
			liveData.Unmarshal(&v)
			fmt.Println(v)
		}, func(err error) () {
			fmt.Println(err)
		})
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(6))*time.Second)
		unsubscribe()
	}
}
