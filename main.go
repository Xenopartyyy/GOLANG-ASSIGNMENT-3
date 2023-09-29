package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Weather struct {
	Water int
	Wind  int
}

func (w *Weather) SetValue(water int, wind int) {
	w.Water = water
	w.Wind = wind
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	for {
		randNumWater := rand.Intn(101)
		randNumWind := rand.Intn(101)

		v := Weather{}
		v.SetValue(randNumWater, randNumWind)

		formatJson, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(formatJson))

		if randNumWater <= 5 {
			fmt.Println("status water : aman")
		} else if randNumWater >= 5 && randNumWater <= 8 {
			fmt.Println("status water : siaga")
		} else if randNumWater > 8 {
			fmt.Println("status water : bahaya")
		}

		if randNumWind <= 6 {
			fmt.Println("status wind : aman")
		} else if randNumWind >= 7 && randNumWind <= 15 {
			fmt.Println("status wind : siaga")
		} else if randNumWind > 15 {
			fmt.Println("status wind : bahaya")
		}

		fmt.Println("")

		time.Sleep(2 * time.Second)
	}
}
