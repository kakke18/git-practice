package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"time"
)

func main() {
	zun := "ズン"
	doko := "ドコ"
	zundoko := []string{zun, doko}
	zunzunzunzundoko := []string{zun, zun, zun, zun, doko}
	result := []string{}

	rand.Seed(time.Now().UnixNano())

	for {
		time.Sleep(1 * time.Second)
		word := zundoko[rand.Intn(2)]
		fmt.Printf("%s ", word)
		result = append(result, word)
		if len(result) >= 5 {
			length := len(result)
			answer := result[length-5:]
			if reflect.DeepEqual(answer, zunzunzunzundoko) {
				fmt.Println("キ・ヨ・シ！")
				os.Exit(0)
			}
		}
	}
}
