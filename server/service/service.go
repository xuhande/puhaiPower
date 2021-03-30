package service

import (
	"fmt"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/viper"
	"puhai/server/global"
	"puhai/server/util"
)

func GetData() []float32 {
	return util.GetFromDb()
}

func GenerateToken() string {
	bucket := viper.GetString("qiniu.bucket")
	fmt.Println(bucket)
	putPolicy := storage.PutPolicy{Scope: bucket}
	upToken := putPolicy.UploadToken(global.MAC)
	return upToken
}