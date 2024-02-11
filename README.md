## Ardor package

Pre alpha version. Not ready at all =)

### How to use

```
go get -u "github.com/scor2k/ardor"
```

```go
ardor := Ardor{Endpoint: "https://random.api.nxter.org/ardor"}
params := map[string]interface{}{
    "requestType": "getTime",
}
data, err := ardor.PostRequest(params)
```

