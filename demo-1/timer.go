package main

import "fmt"
import "time"

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println(timer1)
	<-timer1.C
	fmt.Println("Timer1 is expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<- timer2.C
		fmt.Println("Timer2 is expired")
	}()
	
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Print("Timer2 is stopped")
	}
}