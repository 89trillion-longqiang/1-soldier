package ctrl

import (
	"awesomeProject/handle"
	"github.com/gin-gonic/gin"
)

func GetStage(c *gin.Context)  {
	stage := c.Query("stage")
	retMap,retInfo := handle.HandleGetStage(stage)
	c.JSON(200, gin.H{
		"condition":retMap["condition"],
		"stage" : retMap["stage"],
		"data" : retInfo,
	})
}
func GetArmyAssembled(c *gin.Context){
	rarity := c.Query("rarity")
	stage  := c.Query("stage")
	cvc := c.Query("cvc")
	dataRes ,condition:= handle.HandleGetArmyAssembled(rarity,stage,cvc,*handle.UMAP)
	c.JSON(200, gin.H{
		"condition":condition,
		"stage" : stage,
		"data" : dataRes,
	})
}

func IdGetRarity(c *gin.Context){
	id := c.Query("id")
	ret := handle.HandleIdGetRarity(id)

	c.JSON(200, gin.H{
		"condition":ret["condition"],
		"id":ret["id"],
		"Rarity" :ret["Rarity"],
	})
}
func IdGetCombatPoints(c *gin.Context){
	id := c.Query("id")
	ret := handle.HandleIdGetCombatPoints(id)

	c.JSON(200, gin.H{
		"condition":ret["condition"],
		"id":ret["id"],
		"CombatPoints" :ret["CombatPoints"],
	})

}

func GetSoldiersJson(c *gin.Context){
	stage := c.Query("stage")
	retMap , retInfo := handle.HandleGetSoldiersJson(stage)

	c.JSON(200, gin.H{
		"condition":retMap["condition"],
		"stage" : retMap["stage"],
		"data" : retInfo,
	})
}