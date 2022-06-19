package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica("1")
	generica(true)
	generica([2]string{"1","2"})
	generica(map[string]string{"Olá": "Olá"})

	mapa := map[interface{}]interface{} {
	1 : "String",
	float32(100) : true,
	"String": "String",
	true : float64(12),
	}
	fmt.Println(mapa)
}
