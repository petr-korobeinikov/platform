# vault-integrated-service

## Vault access from `Go` code

```shell
go run cmd/service/main.go
```

## Vault access from `vault` cli

```shell
VAULT_ADDR='http://0.0.0.0:8200' VAULT_TOKEN=secret vault kv put secret/foo foo=foo
VAULT_ADDR='http://0.0.0.0:8200' VAULT_TOKEN=secret vault kv get secret/foo
```
