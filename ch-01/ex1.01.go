package ch01

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Main() {
	ex01()
	act01()
}

func ex01() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}

func act01() {
	var (
		fname          string = "Allison"
		lname          string = "Duncan"
		age            int    = 29
		peanut_allergy bool   = false
	)

	fmt.Println(fname, lname, age, peanut_allergy)
}
