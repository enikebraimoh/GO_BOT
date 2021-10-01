package main

import (
    "fmt"
    fb "github.com/huandu/facebook/v2"
)

func main() {
    res, _ := fb.Get("/538744468", fb.Params{
        "fields": "first_name",
        "access_token": "a-valid-access-token",
    })
    fmt.Println("Here is my Facebook first name:", res["first_name"])
}
