package vk

import (
	"encoding/json"
	"errors"
	"fmt"
	"friends-graph/user"
	"net/http"
)

type Error struct {
	ErrorCode     int    `json:"error_code"`
	ErrorMsg      string `json:"error_msg"`
	RequestParams []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"request_params"`
}

type ErrorResponse struct {
	Error `json:"error"`
}

type FriendsResponse struct {
	Response struct {
		Count int64          `json:"count"`
		Items []user.UserDTO `json:"items"`
	} `json:"response"`
}

type ServiceInterface interface {
	GetFriendsList(id string) ([]user.User, error)
}

type Service struct {
	client *http.Client
	token  string
}

func NewVKService(httpClient *http.Client, token string) *Service {
	return &Service{
		client: httpClient,
		token:  token,
	}
}

func (s *Service) GetFriendsList(id user.Id) ([]user.User, error) {
	friendsMethodURL := fmt.Sprintf(
		"https://api.vk.com/method/friends.get?user_id=%d&access_token=%s&order=name&fields=name&name_case=nom&v=5.154",
		id,
		s.token)

	resp, err := s.client.Get(friendsMethodURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to VK API: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var vkErr ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&vkErr); err != nil {
			return nil, fmt.Errorf("failed to decode VK API error response: %w", err)
		}
		return nil, errors.New(vkErr.Error.ErrorMsg)
	}

	var friendsResp FriendsResponse
	if err := json.NewDecoder(resp.Body).Decode(&friendsResp); err != nil {
		return nil, fmt.Errorf("failed to decode VK API response: %w", err)
	}

	var friends []user.User
	for _, item := range friendsResp.Response.Items {
		name := fmt.Sprintf("%s %s", item.FirstName, item.LastName)
		friendAsUser := user.User{Id: item.Id, Name: name, IsClosed: !item.CanAccessClosed && item.IsClosed}

		friends = append(friends, friendAsUser)
	}

	return friends, nil
}
