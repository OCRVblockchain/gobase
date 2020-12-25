package swagger

import (
	"github.com/gin-gonic/gin"
	//nolint: golint //reason: blank import because of swagger docs init
	ginSwagger "github.com/swaggo/gin-swagger"                // gin-swagger middleware
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
	_ "github.com/wajox/gobase/api"
	"github.com/wajox/gobase/internal/web/controllers/apiv1"
)

var (
	_ apiv1.Controller = (*Controller)(nil)
)

type Controller struct {
	apiv1.BaseController
}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
