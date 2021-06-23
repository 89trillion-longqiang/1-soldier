package ctrl

import (
	"awesomeProject/handle"
	"github.com/gin-gonic/gin"
)

func GetStage(c *gin.Context)  {
	stage := c.Query("stage")
	if len(stage) == 0{
		c.JSON(200, gin.H{
			"condition":"parameter error",
		})
		return
	}
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
	if len(rarity)==0 && len(stage)==0 && len(cvc)==0{
		c.JSON(200, gin.H{
			"condition":"parameter error",
		})
		return
	}
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
	if len(id) == 0 && len(ret) ==0 {
		c.JSON(200, gin.H{
			"condition":"parameter error",
		})
		return
	}
	c.JSON(200, gin.H{
		"condition":ret["condition"],
		"id":ret["id"],
		"Rarity" :ret["Rarity"],
	})
}
func IdGetCombatPoints(c *gin.Context){
	id := c.Query("id")
	if len(id) == 0 {
		c.JSON(200, gin.H{
			"condition":"parameter error",
		})
		return
	}
	ret := handle.HandleIdGetCombatPoints(id)

	c.JSON(200, gin.H{
		"condition":ret["condition"],
		"id":ret["id"],
		"CombatPoints" :ret["CombatPoints"],
	})

}

func GetSoldiersJson(c *gin.Context){
	stage := c.Query("stage")
	if len(stage) ==0 {
		c.JSON(200, gin.H{
			"condition":"parameter error",
		})
		return
	}

	retMap , retInfo := handle.HandleGetSoldiersJson(stage)
	c.JSON(200, gin.H{
		"condition":retMap["condition"],
		"stage" : retMap["stage"],
		"data" : retInfo,
	})
}