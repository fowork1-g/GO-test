package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const DEFAULT_NUM_TO = 1000
const DEFAULT_FLAG_UNIC = true

type MatrixConf struct {
	x          int
	y          int
	numFrom    int
	numTo      int
	shouldUnic *bool
}

func PrepareMatrixConf(conf *MatrixConf) (*MatrixConf, error) {
	if conf.numFrom == 0 && conf.numTo == 0 {
		conf.numTo = DEFAULT_NUM_TO
	}

	if conf.shouldUnic == nil {
		flag := DEFAULT_FLAG_UNIC
		conf.shouldUnic = &flag
	}

	if conf.numFrom > conf.numTo {
		conf.numFrom, conf.numTo = conf.numTo, conf.numFrom
	}

	if *conf.shouldUnic == true && (*conf).x*(*conf).y > conf.numTo-conf.numFrom {
		return nil, errors.New(
			fmt.Sprintf(
				"the range for generating unique numbers is too small! Should be bigger then %d",
				conf.numTo-conf.numFrom,
			),
		)
	}

	if conf.x <= 0 || conf.y <= 0 {
		return nil, errors.New(
			"the number of elements in the matrix horizontally and vertically must be greater than zero",
		)
	}

	return conf, nil
}

func MakeMatrix(conf MatrixConf) [][]int {
	var mem map[int]bool = nil

	if conf.shouldUnic != nil && *conf.shouldUnic {
		mem = make(map[int]bool)
	}

	gen := makeNumberGenerator(conf, mem)

	matrix := make([][]int, conf.x)

	for i := 0; i < conf.x; i++ {
		matrix[i] = make([]int, conf.y)

		for n := 0; n < conf.y; n++ {
			matrix[i][n] = gen()
		}
	}

	return matrix
}

func makeNumberGenerator(conf MatrixConf, memForUnics map[int]bool) func() int {
	simplGen := func() int {
		return rand.Intn(conf.numTo-conf.numFrom) + conf.numFrom
	}

	if memForUnics == nil {
		return simplGen
	} else {
		return func() int {
			for {
				new := simplGen()
				if !memForUnics[new] {
					memForUnics[new] = true
					return new
				}
			}
		}
	}
}
