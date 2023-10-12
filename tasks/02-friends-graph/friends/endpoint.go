package friends

import (
	"friends-graph/user"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GraphController struct {
	graphService *GraphService
}

type GetFriendsGraphRequest struct {
	Id    string `valid:"required~Id is required, numeric~Id should be a number"`
	token string `valid:"required~Token is required"`
}

func NewGraphController(graphService *GraphService) *GraphController {
	return &GraphController{graphService: graphService}
}

func (c *GraphController) GetFriendsGraph(ctx *gin.Context) {
	request := GetFriendsGraphRequest{Id: ctx.Param("id"), token: ctx.Param("token")}

	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(request.Id, 10, 64)

	graph, err := c.graphService.BuildGraph(user.Id(id), 3)
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
