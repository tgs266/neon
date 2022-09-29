package services

import (
	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/api"
)

func AddCredentials(c *gin.Context, req api.AddCredentialsRequest) {
	if req.UsingBasic() {

	}
}
