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
	ex0113()
	ex0115()
	act0102()
	act0103()
	act0104()
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

func ex0113() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}
	fmt.Printf("count1: %#v\ncount2: %#v\ncount3: %#v\nt: %#v\n", count1, count2, count3, t)
}

func ex0115() {
	var count int
	fmt.Println("count: ", count)
	add5Value(count)
	fmt.Println("count: ", count)
	add5Point(&count)
	fmt.Println("count: ", count)
}

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value: ", count)
}

func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point: ", *count)
}

func act0102() {
	a, b := 5, 10
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

const (
	January = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func act0103() {
	count := 5
	var message string
	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Less than 5"
	}
	fmt.Println(message)
}

func act0104() {
	count := 0
	if count < 5 {
		count = 10
	}
	fmt.Println(count == 10)
}
