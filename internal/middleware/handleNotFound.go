package midwares

import (
	"cms/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNotFound(c *gin.Context) {
	utils.JsonResponse(c, 404, 200404, http.StatusText(http.StatusNotFound), nil)
}
