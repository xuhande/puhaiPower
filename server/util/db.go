package util

import (
	"puhai/server/global"
	"puhai/server/models"
)

func GetFromDb() []float32 {
	var data []models.Data
	pl := make([]float32, 0)
	global.PH_DB.Find(&data)
	for _, v := range data {
		pl = append(pl, v.PowerLoad)
	}
	return pl
}