### Initialise (will return keys and a token)
```
curl -X PUT http://0.0.0.0:8200/v1/sys/init --data '{"secret_shares":1, "secret_threshold":1}'
```
### Use one of the keys to unseal
```
curl -X PUT http://0.0.0.0:8200/v1/sys/unseal --data '{"key":"a5e665962f544dd16471c120c5500a7906cfbaeb3f18ae0fc6c5c71d444f0a90"}'
```
### Use the root token to store something
```
curl -X PUT http://0.0.0.0:8200/v1/secret/foo/bar -H "X-Vault-Token: f9bd7f8c-4234-e3de-acad-076682dd2733" --data '{"some_api_key":"bzzp"}'