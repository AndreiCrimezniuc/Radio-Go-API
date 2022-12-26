package router_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (env *DIEnv) HandleRoot(g *gin.Context) {
	g.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
