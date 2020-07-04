package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	cont := 0
	for true {
		publish(cont)
		time.Sleep(time.Second/2)
		cont ++
	}
}

func publish (num int)  {
	nc, _ := nats.Connect("nats://0.0.0.0:4222")
	strs := `{"Nombre": "Jimmie Dalton","Departamento": "Chiquimula","Edad": 66,"Forma de contagio": "Comunitario","Estado": "Recuperado"  }`
	nc.Publish("foo", []byte(strs))
	nc.Close()
	fmt.Println(num)
}





