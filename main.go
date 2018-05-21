package main

import (
	"time"

)

func main() {
	RunTimer()
}
func RunTimer() {
	GetDashboardTimer := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-GetDashboardTimer.C:
			GetDashboardData()
		}
	}
}
func GetDashboardData() {
	go func() {

	}()
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}