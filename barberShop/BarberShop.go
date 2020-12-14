package barberShop

import (
	"fmt"
	"math/rand"
)

const TotalBarbers = 3
const TotalClients = 5

type Client struct {
	DNI uint32
}

type BarberShop struct {
	Positions int
	Wallet    int
	Client
}

func GenerateDNI() Client {
	var client Client
	client.DNI = rand.Uint32()
	return client
}

func SendClient(AllClient int) []Client {
	clients := make([]Client, 0)
	for i := 0; i < AllClient; i++ {
		DNIGenerator := GenerateDNI()
		clients = append(clients, DNIGenerator)

		//time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
	}

	return clients
}

func GetClientByPosition() {

	AllClients := SendClient(TotalClients)

	for _, client := range AllClients {
		fmt.Println(client)

	}
}

//-------------------------------------------------------------------EXECUTION----------

//-------------------------------------------------------------------EXECUTION----------

func CreatingBarbers() {
	barbers := make([]BarberShop, TotalBarbers)
	for position, _ := range barbers {
		_ = BarberShop{
			Positions: position,
			Wallet:    rand.Intn(100),
		}
		//barberMethod.Positions = position
		//barberMethod.Wallet = rand.Intn(100)
		//fmt.Println(barberMethod)
	}
	fmt.Println(barbers)

}

func CreateBarber(barber *BarberShop, position int) {
	barber.Positions = position
	barber.Wallet = rand.Intn(100)
}

func CreatingBarbers2() {
	barbers := make([]BarberShop, TotalBarbers)
	for i, v := range barbers {
		CreateBarber(&v, i)
		fmt.Println(v)
	}
	fmt.Println(barbers)
}

//------------------------------------------------------------------------------------------------

func GetClientByPosition2() {
	//Generate Clients
	AllClients := SendClient(TotalClients)
	//Only read all clients
	for _, client := range AllClients {
		fmt.Println(client)
	}
	CreatingBarbers2()

}

//------------------------------------------------------------------------------------------------

//func AvailableBarber (barbers []BarberShop) bool{}

type Persona struct {
	Name string
	Age  int
}

func PrintPersona() {

	persona := Persona{
		Name: "Julio",
		Age:  24,
	}
	AddYear(&persona)
	fmt.Println(persona)
}

func AddYear(persona *Persona) {
	persona.Age++
}
