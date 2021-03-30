package global

import (
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"gorm.io/gorm"
)

var (
	PH_DB *gorm.DB
	MAC *qbox.Mac
)