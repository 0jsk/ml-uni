package main

import (
	"fmt"
	"friends-graph/vk"
	"log"
	"net/http"
)

func main() {
	httpClient := http.Client{}
	token := "-"

	service := vk.NewVKService(&httpClient, token)

	friends, err := service.GetFriendsList(177804866)
	if err != nil {
		log.Fatalf("Failed to fetch friend list: %v", err)
	}

	fmt.Println("Friend list:")
	for _, friend := range friends {
		fmt.Println(friend)
	}
}
