package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type Dude struct {
	Name    string
	Ability string
}

func main() {

	dudes := [...]*Dude{
		&Dude{Name: "Tibi", Ability: "doing something else"},
	}
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(dudes))
	pick := *dudes[idx]

	var message bytes.Buffer
	message.WriteString(pick.Name)
	message.WriteString(" because loves ")
	message.WriteString(pick.Ability)
	fmt.Println(message.String())
}
