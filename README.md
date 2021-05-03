# tillhub

With this repository you can access the Tillhub API with the following functions. We will continue to develop these as we go along, so if you want to collaborate, feel free.

If you need help, please feel free to ask us:  [info@jj-ideenschmiede.de](mailto:info@jj-ideenschmiede.de).

## Install

```console
go get github.com/jjideenschmiede/tillhub
```

## How to use it?

Here you can find an overview of all currently created functions.

### Authentication

```go
// Get bearer token
data, err := tillhub.Auth("email", "password")
if err != nil {
    fmt.Println(err)
}

// Print token
fmt.Println(data.token)
```