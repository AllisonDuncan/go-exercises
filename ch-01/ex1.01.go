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
	ex01()
	ex0109()
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

func ex0109() {
	//main course
	var total float64 = 2 * 13
	fmt.Println("Sub :", total)
	//drinks
	total += (4 * 2.25)
	fmt.Println("Sub :", total)
	//discount
	total -= 5
	fmt.Println("Sub :", total)
	//tip
	tip := total * 0.2
	fmt.Println("Tip :", tip)
	total = total + tip
	fmt.Println("Total :", total)
	//split bill
	split := total / 2
	fmt.Println("Split :", split)
	//reward
	visitCount := 24
	visitCount += 1
	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("Reward")
	}

}
