package persist

import (
	"fmt"
	"testing"
	"time"
)

type obj struct {
	Name   string
	Number int
	When   time.Time
}

func TestPersist(t *testing.T) {
	o := obj{
		Name:   "Mat",
		Number: 47,
		When:   time.Now(),
	}
	if err := Save("./file.tmp", o); err != nil {
		fmt.Println(err)
	}
	//load it back
	var o2 obj
	if err := Load("./file.tmp", &o2); err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(o2)
}
