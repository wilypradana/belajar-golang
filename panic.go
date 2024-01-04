package main

import "fmt"

func endApp()  {
	fmt.Println("Selesai memanggil function")
}

func runApplication(error bool)  {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("aplikasi dijalankan")
}

func main() {
	runApplication(false)
}