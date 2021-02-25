package models

type data struct {
	date string
	powerLoad float32
	tem float32
}

func NewData(date string, powerLoad float32) *data {
	return &data{
		date:date,
		powerLoad: powerLoad,
	}
}