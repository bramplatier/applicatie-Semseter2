package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func init() {
	logFile, err := os.OpenFile("trace.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("couldn't create logfile")
		os.Exit(1)
	}

	writer := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(writer)
}

func main() {
	fmt.Printf("%s! Beste gebruiker, wat wil je doen?\n", bepaalGroet())
	var choice int
	for {
		fmt.Println("1. Kenteken registreren")
		fmt.Println("2. Toegang park controleren")
		fmt.Println("3. Exit")
		fmt.Print("Maak je keuze: ")

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			registerKenteken()
		case 2:
			checkToegangPark()
		case 3:
			fmt.Println("Tot ziens!")
			os.Exit(0)
		default:
			fmt.Println("Ongeldige keuze, probeer opnieuw.")
		}
	}
}

func registerKenteken() {
	var voornaam, kenteken string

	fmt.Print("Wat is de gebruikersnaam? ")
	fmt.Scanln(&voornaam)
	fmt.Print("Wat is het kenteken? ")
	fmt.Scanln(&kenteken)

	bookings, err := laatBookingsNaarFile("bookings.json")
	if err != nil {
		log.Println(err)
		return
	}

	bookings = append(bookings, Booking{Name: voornaam, Kenteken: kenteken})
	if err := schrijfBookingsNaarFile(bookings, "bookings.json"); err != nil {
		log.Println("Error writing to JSON file:", err)
		return
	}
}

func checkToegangPark() {
	var kenteken string

	fmt.Print("Hallo, wat is uw kenteken? ")
	fmt.Scanln(&kenteken)

	bookings, err := laatBookingsNaarFile("bookings.json")
	if err != nil {
		log.Println(err)
		return
	}

	var found bool
	var gebruikersnaam string

	for _, booking := range bookings {
		if booking.Kenteken == kenteken {
			found = true
			gebruikersnaam = booking.Name
			break
		}
	}

	if found {
		fmt.Printf("Beste %s, het kenteken is geldig, ga maar door\n", gebruikersnaam)
	} else {
		fmt.Println("Het kenteken is niet geldig, u mag niet door 😜)")
	}
}

func laatBookingsNaarFile(filename string) ([]Booking, error) {
	var bookings []Booking

	jsonFile, err := os.Open(filename)
	if err != nil {
		return bookings, err
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&bookings)
	if err != nil {
		return bookings, err
	}

	return bookings, nil
}

func schrijfBookingsNaarFile(bookings []Booking, filename string) error {
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

func BekijkKentekens(bookings []Booking, kenteken string) bool {
	for _, booking := range bookings {
		if booking.Kenteken == kenteken {
			return true
		}
	}
	return false
}

func bepaalGroet() string {
	hour := time.Now().Hour()
	var groet string
	if hour >= 7 && hour < 12 {
		groet = "Goedemorgen"
	} else if hour >= 12 && hour < 18 {
		groet = "Goedemiddag"
	} else if hour >= 18 && hour < 23 {
		groet = "Goedenavond"
	} else {
		groet = "Sorry, de parkeerplaats is 's nachts gesloten"
	}
	return groet
}

type Booking struct {
	Name     string `json:"name"`
	Kenteken string `json:"kenteken"`
}
