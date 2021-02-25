package main

import (
	"fmt"
	"puhai/server/models"
)

func main() {
	data := models.NewData("9.30", 350)
	fmt.Println(data)
}
