package main

import (
	"os"
	"fmt"
	"github.com/hashicorp/vault/api"

)

var (
	token = os.Getenv("VAULT_TOKEN")
	vault_addr = os.Getenv("VAULT_ADDR")
	vault_path = os.Getenv("VAULT_PATH")
)

func main() {
	config := &api.Config{
		Address: vault_addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.SetToken(token)
	c := client.Logical()
	secret, err := c.Read(vault_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.Create(".env")
	if err != nil {
        fmt.Println(err)
    }
    defer f.Close()

	for k, v := range secret.Data {
		str := fmt.Sprintf("%s=%s\n", k,v) 
		_, err := f.WriteString(str)

        if err != nil {
            fmt.Println(err)
        }
    }
}