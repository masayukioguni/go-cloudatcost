package main

import (
	"fmt"
	"github.com/masayukioguni/go-digitalocean/digitalocean"
	"os"
)

var applicationName = "go-tugboat"
var version string

func main() {
	APIKey := os.Getenv("DIGITALOCEAN_API_TOKEN")
	client, _ := digitalocean.NewClient(&digitalocean.Option{APIKey: APIKey})

	account, hr, err := client.AccountService.GetUserInformation()

	if err != nil {
		fmt.Printf("error: %v\n\n", err)
		return
	}

	if hr.StatusCode != 200 {
		fmt.Printf("http response error: %+v %+v \n\n", hr, err)
		return
	}

	fmt.Printf("%s(%s) verified:%t limit:%d\n",
		account.Email, account.UUID, account.EmailVerified, account.DropletLimit)

}
