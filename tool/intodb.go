package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RealDataMz struct {
	PowerLoad float32
	Temperature float32
}

func main() {
	//f, _ := os.Open("../data/train.csv")
	//reader := bufio.NewReader(f)
	dsn := "mysql_nj_jay_com:000000@tcp(nj-jay.com:3306)/mysql_nj_jay_com?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("连接成功")
	err = db.AutoMigrate(&RealDataMz{})
	if err != nil {
		fmt.Println("创建失败")
		return
	}
	//var n int
	//for {
	//	n += 1
	//	bytes, _, _ := reader.ReadLine()
	//	if strings.Split(string(bytes), ",")[1] == "名称" {
	//		continue
	//	}
	//	if strings.Split(string(bytes), ",")[1] != "民正重钢" {
	//		break
	//	}
	//	pl := cast.ToFloat32(strings.Split(string(bytes), ",")[3])
	//	t := cast.ToFloat32(strings.Split(string(bytes), ",")[14])
	//
	//	rd := &RealDataMz{
	//		PowerLoad: pl,
	//		Temperature:t,
	//	}
	//	db.Create(rd)
	//}
	var realDataMz RealDataMz
	db.First(&realDataMz)
	fmt.Println(realDataMz)
}