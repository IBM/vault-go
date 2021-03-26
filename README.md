# vault-go

Vault go contains the go definitions of the Vault types imported into go programs

All of these types are defined as Kubernetes CRD's. A future implentation will build controllers that allow a k8s cluster to maintain the state of Vault's configuration.

## Types

- [JWTRole](api/v1/jwtrole_types.go)
- [PKIRole](api/v1/pkirole_types.go)
- [SSHRole](api/v1/sshrole_types.go)
- [VaultAuth](api/v1/vaultauth_types.go)
- [VaultEndpoint](api/v1/vaultendpoint_types.go)
- [VaultNamespace](api/v1/vaultnamespace_types.go)
- [VaultPolicy](api/v1/vaultpolicy_types.go)
- [VaultRole](api/v1/vaultrole_types.go)

## Apps using vault-go

- [vault-cli](https://github.com/ibm/vault-cli)

example command to add an api

```bash
kubebuilder create api --group vault --version v1 --kind VaultEndpoint
```