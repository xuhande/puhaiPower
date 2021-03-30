package models

type RealDataMz struct {
	//日期
	Date string `json:"date" gorm:"private key"`

	//电力负荷
	PowerLoad float32 `json:"powerLoad"`

	//温度
	Temperature float32 `json:"tem"`
}

type Login struct {

}