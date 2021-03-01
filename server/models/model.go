package models

type Data struct {
	//日期
	Date string `json:"date" gorm:"private key"`

	//电力负荷
	PowerLoad float32 `json:"powerLoad"`

	//温度
	Tem float32 `json:"tem"`
}

func NewData(date string, powerLoad float32) *Data {
	return &Data{
		Date:date,
		PowerLoad: powerLoad,
	}
}