package services

import (
	"log"
	"strconv"
)

func StringToInt(n string) int {
	i, err := strconv.Atoi(n)
	if err != nil {
		log.Println(err)
	}
	return i
}
