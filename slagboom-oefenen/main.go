package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
	fmt.Printf("%s! Beste gebruiker, wat wil je doen?\n", Bepaalgroet())
	var choice int
	for {
		fmt.Println("1. Kenteken registreren")
		fmt.Println("2. Toegang park controleren")
		fmt.Println("3. Gebruiker verwijderen")
		fmt.Println("4. Gebruiker status wijzigen")
		fmt.Println("5. Exit")
		fmt.Print("Maak je keuze: ")

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			registerKenteken()
		case 2:
			checkToegangPark()
		case 3:
			removeUser()
		case 4:
			updateUserStatus()
		case 5:
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

	bookings, err := loadBookingsFromFile("bookings.json")
	if err != nil {
		log.Println(err)
		return
	}

	bookings = append(bookings, Booking{Name: voornaam, Kenteken: kenteken, Active: true})
	if err := writeBookingsToFile(bookings, "bookings.json"); err != nil {
		log.Println("Error writing to JSON file:", err)
		return
	}
	fmt.Println("Kenteken is succesvol toegevoegd aan het JSON-bestand")
}

func checkToegangPark() {
	var kenteken string

	fmt.Print("Hallo, wat is uw kenteken? ")
	fmt.Scanln(&kenteken)

	bookings, err := loadBookingsFromFile("bookings.json")
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

func removeUser() {
	bookings, err := loadBookingsFromFile("bookings.json")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Lijst van gebruikers:")
	for _, booking := range bookings {
		var status string
		if booking.Active {
			status = "actief"
		} else {
			status = "niet actief"
		}
		fmt.Printf("- Naam: %s, Kenteken: %s, Status: %s\n", booking.Name, booking.Kenteken, status)
	}

	var kenteken string
	fmt.Print("Wat is het kenteken van de gebruiker die u wilt verwijderen? ")
	fmt.Scanln(&kenteken)

	var updatedBookings []Booking
	var found bool

	for _, booking := range bookings {
		if booking.Kenteken != kenteken {
			updatedBookings = append(updatedBookings, booking)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Println("Gebruiker niet gevonden.")
		return
	}

	if err := writeBookingsToFile(updatedBookings, "bookings.json"); err != nil {
		log.Println("Error writing to JSON file:", err)
		return
	}
	fmt.Println("Gebruiker succesvol verwijderd")
}

func updateUserStatus() {
	bookings, err := loadBookingsFromFile("bookings.json")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Lijst van gebruikers:")
	for _, booking := range bookings {
		var status string
		if booking.Active {
			status = "actief"
		} else {
			status = "niet actief"
		}
		fmt.Printf("- Naam: %s, Kenteken: %s, Status: %s\n", booking.Name, booking.Kenteken, status)
	}

	var kenteken string
	var active bool

	fmt.Print("Wat is het kenteken van de gebruiker waarvan u de status wilt wijzigen? ")
	fmt.Scanln(&kenteken)

	fmt.Print("Wilt u deze gebruiker activeren? (ja/nee): ")
	var answer string
	fmt.Scanln(&answer)
	if strings.ToLower(answer) == "ja" {
		active = true
	} else {
		active = false
	}

	var found bool
	for i, booking := range bookings {
		if booking.Kenteken == kenteken {
			found = true
			bookings[i].Active = active
			break
		}
	}

	if !found {
		fmt.Println("Gebruiker niet gevonden.")
		return
	}

	if err := writeBookingsToFile(bookings, "bookings.json"); err != nil {
		log.Println("Error writing to JSON file:", err)
		return
	}
	if active {
		fmt.Println("Gebruiker succesvol geactiveerd")
	} else {
		fmt.Println("Gebruiker succesvol gedeactiveerd")
	}
}

func loadBookingsFromFile(filename string) ([]Booking, error) {
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

func Bepaalgroet() string {
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
	Active   bool   `json:"active"`
}
