package purchases

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/purchase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("purchases")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":purchase_ID", controller.GetByID)
		v10.DELETE(":purchase_ID", controller.Delete)
		v10.PATCH(":purchase_ID", controller.Updated)
	}

	return route
}
