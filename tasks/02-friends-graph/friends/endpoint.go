package friends

import (
	"friends-graph/user"
	"friends-graph/vk"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GraphController struct {
	graphService *GraphService
	vkService    *vk.Service
}

type GetFriendsGraphRequest struct {
	Id string `valid:"required~Id is required, numeric~Id should be a number"`
}

func NewGraphController(graphService *GraphService) *GraphController {
	return &GraphController{graphService: graphService}
}

func (c *GraphController) GetFriendsGraph(ctx *gin.Context) {
	request := GetFriendsGraphRequest{Id: ctx.Param("id")}

	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(request.Id, 10, 64)

	initiator, err := c.vkService.GetUser(user.Id(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch information about initiator"})
		return
	}

	graph, err := c.graphService.BuildGraph(initiator, 3)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to build friends graph"})
		return
	}

	nodes := graph.BFS(user.Id(id), 3)

	var users []*user.User
	for _, node := range nodes {
		users = append(users, node.User)
	}

	ctx.JSON(http.StatusOK, users)
}
