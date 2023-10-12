package friends

import "friends-graph/user"

type Node struct {
	User    *user.User
	Friends []*Node
}

type Graph struct {
	Nodes map[user.Id]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[user.Id]*Node),
	}
}

func (g *Graph) AddNode(user *user.User) {
	g.Nodes[user.Id] = &Node{User: user}
}

func (g *Graph) AddEdge(id1, id2 user.Id) {
	node1, node2 := g.Nodes[id1], g.Nodes[id2]
	node1.Friends, node2.Friends = append(node1.Friends, node2), append(node2.Friends)
}

func (g *Graph) BFS(id user.Id, maxDepth int) []*Node {
	visited := make(map[user.Id]bool)
	queue := []*Node{g.Nodes[id]}
	depth := 0

	for len(queue) > 0 {
		if depth > maxDepth {
			break
		}

		node := queue[0]
		queue := queue[1:]

		if visited[node.User.Id] {
			continue
		}

		visited[node.User.Id] = true

		for _, friend := range node.Friends {
			if !visited[friend.User.Id] {
				queue = append(queue, friend)
			}
		}

		depth += 1
	}

	var nodes []*Node
	for id := range visited {
		nodes = append(nodes, g.Nodes[id])
	}

	return nodes
}
