package main

import (
	"fmt"
	"friends-graph/friends"
	"friends-graph/user"
	"friends-graph/vk"
	"net/http"
)

func main() {
	token := "-"

	httpClient := &http.Client{}

	vkService := vk.NewVKService(httpClient, token)
	graphService := friends.NewGraphService(vkService)

	testUserId := user.Id(210700286)

	user, err := vkService.GetUser(testUserId)
	if err != nil {
		fmt.Printf("Failed to get user: %v\n", err)
		return
	}
	fmt.Printf("Got user: %+v\n", user)

	// Test building the friends graph
	graph, err := graphService.BuildGraph(user, 3)
	if err != nil {
		fmt.Printf("Failed to build friends graph: %v\n", err)
		return
	}

	users := graph.GetUsersFromGraph()
	fmt.Printf("Got %d users in friends graph\n", len(users))
}
