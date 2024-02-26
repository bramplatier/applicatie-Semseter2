package main

import (
	"fmt"
	"time"
)

func main() {
	kentekens := []string{"GB-001-B", "HF-234-S", "HS-HIHI-HA", "test"}
	var groet = Bepaalgroet()
	var kenteken string

	fmt.Printf("%v! Welkom bij Fonteyn Vakantieparken\n", groet)
	fmt.Printf("hallo wat is uw kenteken?\n")
	fmt.Scanln(&kenteken)
	if BekijkKentekens(kenteken, kentekens) {
		fmt.Println("Het kenteken is geldig ga maar door")
	} else {
		fmt.Println("Het kenteken is niet geldig u mag niet door 😜)")
	}
	//test
}
func Bepaalgroet() string {
	hour := time.Now().Hour()
	var groet string
	if hour >= 7 && hour < 12 {
		groet = "goedemorgen"
	} else if hour >= 12 && hour < 18 {
		groet = "goedemiddag"
	} else if hour >= 18 && hour < 23 {
		groet = "goedenavond"
	} else {
		groet = "Sorry, de parkeerplaats is s nachts gesloten"
	}
	return groet
}

func BekijkKentekens(kenteken string, kentekens []string) bool {
	for _, y := range kentekens {
		if y == kenteken {
			return true
		}
	}
	return false
}
