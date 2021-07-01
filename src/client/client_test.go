package client_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(_ *testing.T) {
	fmt.Println(time.Now().String())
	<-time.After(2 * time.Second)
	fmt.Println(time.Now().String())
}
