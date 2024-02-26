package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

//TODO: als geen kenteken vraag alsnog op elke tenteken

func main() {
	var groet = Bepaalgroet()
	var kenteken string

	jsonFile, err := os.Open("bookings.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully opened bookings.json")
	defer jsonFile.Close()

	var bookings []Booking
	err = json.NewDecoder(jsonFile).Decode(&bookings)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("%v! Welkom bij Fonteyn Vakantieparken\n", groet)
	fmt.Printf("Hallo, wat is uw kenteken?\n")
	fmt.Scanln(&kenteken)
	if BekijkKentekens(bookings, kenteken) {
		fmt.Println("Het kenteken is geldig, ga maar door")
	} else {
		fmt.Println("Het kenteken is niet geldig, u mag niet door 😜)")
	}

	fmt.Print("Wilt u dit kenteken toevoegen aan het JSON-bestand? (ja/nee): \n")
	var answer string
	fmt.Scanln(&answer)
	var voornaam string

	if answer == "ja" {
		fmt.Print("wat is het gebruikers naam?\n")

		fmt.Scanln(&voornaam)
		fmt.Print("wat is het kenteken?\n")
		fmt.Scanln(&kenteken)
	}
	if answer == "ja" {
		bookings = append(bookings, Booking{Name: voornaam, Kenteken: kenteken})
		if err := writeBookingsToFile(bookings, "bookings.json"); err != nil {
			fmt.Println("Error writing to JSON file:", err)
			return
		}
		fmt.Println("Kenteken is succesvol toegevoegd aan het JSON-bestand")
	}
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
		groet = "Sorry, de parkeerplaats is 's nachts gesloten"
	}
	return groet
}

func BekijkKentekens(bookings []Booking, kenteken string) bool {
	for _, booking := range bookings {
		if booking.Kenteken == kenteken {
			return true
		}
	}
	return false
}

type Booking struct {
	Name     string `json:"name"`
	Kenteken string `json:"kenteken"`
}

func writeBookingsToFile(bookings []Booking, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(bookings); err != nil {
		return err
	}
	return nil
}
