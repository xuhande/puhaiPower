package service

import (
	"puhai/server/util"
)

func GetData() []float32 {
	return util.GetFromDb()
}