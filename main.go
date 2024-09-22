package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	conf, err := PrepareMatrixConf(&MatrixConf{x: 5, y: 5, numTo: 100})
	if err != nil {
		log.Fatal(err)
	}

	matrix := MakeMatrix(*conf)

	for _, line := range matrix {
		fmt.Println(line)
	}

	fmt.Println("")

	shouldUnic := false
	conf, err = PrepareMatrixConf(&MatrixConf{x: 5, y: 5, numTo: 20, shouldUnic: &shouldUnic})
	if err != nil {
		log.Fatal(err)
	}

	matrix = MakeMatrix(*conf)

	for _, line := range matrix {
		fmt.Println(line)
	}
}
