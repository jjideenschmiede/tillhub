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

### Read all products

```go
// Read all products
products, err := tillhub.ReadProducts(data.User.Id, "Bearer "+data.Token)
if err != nil {
    fmt.Println(err)
}

// Print products
fmt.Println(products)
```

### Create an product

```go
// Create product data
body := CreateProductBody{
	true,
    "product",
    "J&J Testprodukt",
    "673ba8d1-f544-48be-b432-122633e92715",
    "f85910a6-94c4-40bd-a0ac-386ef880c857",
    "JJTEST2021",
    CreateProductBodyAttributes{},
    []CreateProductBodyCodes{},
    nil,
    nil,
    nil,
    CreateProductBodyImages{},
    nil,
    true,
    nil,
    nil,
    nil,
    CreateProductBodyPrices{
        []CreateProductBodyPricesDefaultPrices{
            {
                CreateProductBodyPricesDefaultPricesAmount{
                    69,
                    82.11,
                },
                3,
                "EUR",
                20,
                2,
            },
        },
        []CreateProductBodyPricesBranchPrices{},
    },
    nil,
    true,
    false,
    nil,
    nil,
    CreateProductBodyManufacturer{
        nil,
    },
    CreateProductBodySupplier{
        nil,
    },
    nil,
    "simple",
    nil,
    []CreateProductBodyTags{},
    false,
    nil,
    CreateProductBodyConfiguration{
        true,
        CreateProductBodyConfigurationPricing{
            false,
        },
    },
    nil,
    nil}

// Create product function
create, err := CreateProduct(body, data.User.Id, "Bearer "+data.Token)
if err != nil {
    fmt.Println(err)
}

// Print product data
fmt.Println(create)
```