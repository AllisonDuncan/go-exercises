package ch01

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}
