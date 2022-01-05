package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
)

var (
	token     string
	vaultAddr string
)

func init() {
	_ = godotenv.Load(".platform/env/.env")

	token = os.Getenv("COMPONENT_VAULT_VAULT_VAULT_DEV_ROOT_TOKEN_ID")
	vaultAddr = os.Getenv("COMPONENT_VAULT_VAULT_VAULT_DEV_LISTEN_ADDRESS")
}

func main() {
	config := &api.Config{
		Address: "http://" + vaultAddr,
	}

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	client.SetToken(token)

	input := map[string]interface{}{
		"data": map[string]interface{}{
			"foo": "foo",
		},
	}

	_, err = client.Logical().Write("secret/data/foo", input)
	if err != nil {
		fmt.Println(err)
		return
	}

	secret, err := client.Logical().Read("secret/data/foo")
	if err != nil {
		fmt.Println(err)
		return
	}

	if secret == nil {
		fmt.Println("secret not found")
		return
	}

	m, ok := secret.Data["data"]
	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
		return
	}

	fmt.Printf("foo: %v\n", m.(map[string]interface{})["foo"])
}
