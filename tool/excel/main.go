package main

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/spf13/cast"
	"io/ioutil"
)

func main() {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")
	info, _ := ioutil.ReadDir("F:\\all-file")
	var location = 1
	for _, f := range info {
		_ = xlsx.SetCellValue("Sheet1", "A" + cast.ToString(location), f.Name())
		_ = xlsx.SetCellValue("Sheet1", "B" + cast.ToString(location), cast.ToString(f.Size()/100)+"Kb")
		_ = xlsx.SetCellValue("Sheet1", "C" + cast.ToString(location), "1")
		location += 1
	}
	xlsx.SetActiveSheet(index)
	_ = xlsx.SaveAs("source.xlsx")
}