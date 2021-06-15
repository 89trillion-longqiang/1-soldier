package router

import (
	"awesomeProject/ctrl"
	"github.com/gin-gonic/gin"
)
func SetUpRoute() *gin.Engine{
	r := gin.Default()
	// GET：请求方式；/getinfo：请求的路径
	// 当客户端以GET方法请求/getinfo，会执行后面的匿名函数

	r.GET("/getstage", ctrl.GetStage)
	r.GET("/getArmyAssembled", ctrl.GetArmyAssembled)
	r.GET("/idGetRarity", ctrl.IdGetRarity)
	r.GET("/idGetCombatPoints",ctrl.IdGetCombatPoints)
	r.GET("/getSoldiersJson", ctrl.GetSoldiersJson)

	return r
}
