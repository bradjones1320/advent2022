package main

import (
	"fmt"
	"testing"
)

func TestBlah(t *testing.T) {

	c := [...]string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	// bigElf := 0
	// tempElf := 0

	for i, cal := range c {

		fmt.Println(i, " - ", cal)
	}

}
