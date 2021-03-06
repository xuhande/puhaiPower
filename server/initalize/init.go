package initalize

import (
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"puhai/server/models"
	"strings"
)

func GormMysql() *gorm.DB {
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	ip := viper.GetString("mysql.ip")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	path := strings.Join([]string{username, ":", password,
		"@(", ip, ":", port, ")/", database, "?charset=utf8&parseTime=true&loc=Local"}, "")

	if db, err := gorm.Open(mysql.Open(path), &gorm.Config{}); err != nil {
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.RealDataMz{},
	)

	if err != nil {
		fmt.Println("register table error")
	}
}

func InitMAC() *qbox.Mac {
	ak, sk := viper.GetString("qiniu.ak"), viper.GetString("qiniu.sk")
	fmt.Println(ak)
	fmt.Println(sk)
	mac := qbox.NewMac(ak, sk)
	return mac
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read error")
		os.Exit(0)
	}
}
