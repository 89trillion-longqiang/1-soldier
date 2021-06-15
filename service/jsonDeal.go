package service

import (
	"awesomeProject/soldierInfo"
)

// GetArmyAssembled /通过稀有度，阶段来获取已解锁的所有士兵的信息
func GetArmyAssembled(rarity string , stage string ,cvc string, Umap map[string]*soldierInfo.Info) (data []*soldierInfo.Info,conditon string){
	if rarity == "" || stage == "" || cvc == "" {
		return nil,"error"
	}
	for _,v :=range Umap{
		if v.Rarity == rarity  && v.Desc <= stage {
			if v.UnlockArena == "1"{
				data = append(data,v)
			}
		}
	}
	return data,"success"
}

// GetRarity /通过id来获取稀有度
func GetRarity(id string,Umap map[string]*soldierInfo.Info) string{
	for k,v:=range Umap{
		if id == k{
			return  v.Rarity
		}
	}

	return ""
}

// GetCombatPoints /通过id来获取战力
func GetCombatPoints(id string,Umap map[string]*soldierInfo.Info)  string{
	for k,v:=range Umap{
		if id == k{
			return  v.CombatPoints
		}
	}

	return ""
}

// GetSoldiersJson /通过阶段来获取已解锁的士兵信息
func GetSoldiersJson(stage string,Umap map[string]*soldierInfo.Info) (data []*soldierInfo.Info) {
	if stage == "" {
		return
	}
	for _, v := range Umap {
		if v.Desc <= stage {
			data = append(data, v)
		}
	}
	return
}