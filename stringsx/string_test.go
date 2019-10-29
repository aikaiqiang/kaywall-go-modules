package stringsx

import (
	"fmt"
	"testing"
)

func TestHelloByLang(t *testing.T) {

	if greet, err := HelloByLang("kaywall", "fr"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greet)
	}
}

func TestHello(t *testing.T) {
	fmt.Println(Hello("kaywall"))
}
