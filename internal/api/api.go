package api

import (
	"fmt"
	"yokanban-cli/internal/http"
)

func List() {
	// TODO to implement
}

func Get() {
	// TODO to implement
}

func Create() {
	// TODO to implement
}

func Delete() {
	// TODO to implement
}

func Test() {
	fmt.Println("Test()")
	token := GetAccessToken()
	body := http.Get(RouteTest, token)

	fmt.Println(body)
}
