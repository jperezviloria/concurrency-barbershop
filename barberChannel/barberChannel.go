package barberChannel

import (
	"fmt"
	"math/rand"
	"time"
)

type BarberShop struct {
	Position int
	Wallet   int
}

type Client struct {
	DNI int
}

func SendInfo(channelA chan<- Client, totalInfo int) {

	for i := 0; i < totalInfo; i++ {

		channelA <- Client{DNI: rand.Int()}
		fmt.Printf("-----------------------> Creating id: %v \n", channelA)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	}
	fmt.Println("finish sendInfo")
}

//func DataBase () []BarberShop{
//
//	barbers := []BarberShop{
//		{
//			Position: 1,
//			Wallet: 0,
//		},
//		{
//			Position: 2,
//			Wallet: 0,
//		},
//		{
//			Position: 3,
//			Wallet: 0,
//		},
//	}
//	return barbers
//}

//func FindBarberByPosition() {
//	allBarbers := DataBase()
//	for _,v := range allBarbers{
//		fmt.Printf("barber %d have a wallet with %d\n", v.Position, v.Wallet)
//	}
//}

func ReceiveInfo(channelA <-chan Client, channelB <-chan struct{}) {

	var valA int
	var valB int
	var valC int
	for {
		select {
		case i := <-channelA:
			//time.Sleep(1 * time.Second)
			fmt.Printf("init case i: %d in time %s\n", i.DNI, time.Now())
			fmt.Println(i)
			fmt.Println("case i, wait 7 seconds to follow")
			time.Sleep(3 * time.Second)
			valA++
			fmt.Println(valA)
		case j := <-channelA:
			//time.Sleep(1 * time.Second)
			fmt.Printf("init case j: %d in time %s\n", j.DNI, time.Now())
			fmt.Println(j)
			fmt.Println("case j, wait 7 seconds to follow")
			time.Sleep(3 * time.Second)
			valB++
			fmt.Println(valB)
		case k := <-channelA:
			//time.Sleep(1 * time.Second)
			fmt.Printf("init case k: %d in time %s\n", k.DNI, time.Now())
			fmt.Println(k)
			fmt.Println("case k, wait 7 seconds to follow")
			time.Sleep(3 * time.Second)
			valC++
			fmt.Println(valC)
		case <-channelB:
			break
		}
	}
}

func MainInfo() {
	channelA := make(chan Client, 1)
	channelB := make(chan struct{})

	go ReceiveInfo(channelA, channelB)
	SendInfo(channelA, 100)
	channelB <- struct{}{}
}
