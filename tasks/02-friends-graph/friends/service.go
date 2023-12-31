package friends

import (
	"friends-graph/user"
	"friends-graph/vk"
)

type GraphService struct {
	vkService *vk.Service
}

func NewGraphService(vkService *vk.Service) *GraphService {
	return &GraphService{vkService: vkService}
}

func (s *GraphService) BuildGraph(initiator *user.User, maxDepth int) (*Graph, error) {
	graph := NewGraph()
	queue := []user.Id{initiator.Id}
	visited := make(map[user.Id]bool)
	depth := 0

	for len(queue) > 0 {
		if depth > maxDepth {
			break
		}

		id := queue[0]
		queue = queue[1:]

		if visited[id] {
			continue
		}

		visited[id] = true
		graph.AddNode(initiator)

		friends, err := s.vkService.GetFriendsList(id)
		if err != nil {
			return nil, err
		}

		for _, friend := range friends {
			if !visited[friend.Id] {
				queue = append(queue, friend.Id)
			}

			if _, ok := graph.Nodes[friend.Id]; !ok {
				graph.AddNode(&friend)
			}

			graph.AddEdge(id, friend.Id)
		}

		depth++
	}

	return graph, nil
}
