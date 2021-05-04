# tillhub

With this repository you can access the Tillhub API with the following functions. We will continue to develop these as we go along, so if you want to collaborate, feel free.

If you need help, please feel free to ask us:  [info@jj-ideenschmiede.de](mailto:info@jj-ideenschmiede.de).

## Install

```console
go get github.com/jjideenschmiede/tillhub
```

## How to use it?

Here you can find an overview of all currently created functions.

### Basic authentication

To access the API you need a current Bearer Token. You will receive this token with further information about this function. [Here](https://developer.tillhub.com/tutorials/auth/basic/#basic-auth-flow) you can see the call in tillhub documentation.

```go
// Get bearer token
data, err := tillhub.AuthBasic("email", "password")
if err != nil {
    fmt.Println(err)
}

// Print token
fmt.Println(data.token)
```

### API key authentication

To access the API you need a current Bearer Token. You will receive this token with further information about this function. [Here](https://developer.tillhub.com/tutorials/auth/key/) you can see the call in tillhub documentation.

```go
// Get bearer token
data, err := tillhub.AuthKey("accoundId", "apiKey")
if err != nil {
    fmt.Println(err)
}

// Print token
fmt.Println(data.token)
```

### Read all products

If you want to read out all products, you can do it with this function. Both simple and variable products are returned. [Here](https://developer.tillhub.com/tutorials/products/read/#getting-started) you can see the call in tillhub documentation.

```go
// Read all products
products, err := tillhub.ReadProducts("userId", "token")
if err != nil {
    fmt.Println(err)
}

// Print products
fmt.Println(products)
```

### Create product

If you want to create a product, you can do this using the following function call. [Here](https://developer.tillhub.com/tutorials/products/create/#getting-started) you can see the call in tillhub documentation.

```go
// Create product data
body := tillhub.CreateProductBody{
	true,
    "product",
    "J&J Testprodukt",
    "UUIDofTheRevenueAccountObjectInTillhub",
    "UUIDofTheTAXAccountObjectInTillhub",
    "JJTEST2021",
	tillhub.CreateProductBodyAttributes{},
    []tillhub.CreateProductBodyCodes{},
    nil,
    nil,
    nil,
    tillhub.CreateProductBodyImages{},
    nil,
    true,
    nil,
    nil,
    nil,
    tillhub.CreateProductBodyPrices{
        []tillhub.CreateProductBodyPricesDefaultPrices{
            {
			    tillhub.CreateProductBodyPricesDefaultPricesAmount{
                    69,
                    82.11,
                },
                3,
                "EUR",
                20,
                2,
            },
        },
        []tillhub.CreateProductBodyPricesBranchPrices{},
    },
    nil,
    true,
    false,
    nil,
    nil,
    tillhub.CreateProductBodyManufacturer{
        nil,
    },
    tillhub.CreateProductBodySupplier{
        nil,
    },
    nil,
    "simple",
    nil,
    []tillhub.CreateProductBodyTags{},
    false,
    nil,
    tillhub.CreateProductBodyConfiguration{
        true,
		tillhub.CreateProductBodyConfigurationPricing{
            false,
        },
    },
    nil,
    nil}

// Create product function
create, err := tillhub.CreateProduct(body, "userId", "token")
if err != nil {
    fmt.Println(err)
}

// Print product data
fmt.Println(create)
```

### Update product

If you want to update a product, then you can do with the following function. Some data is needed for this. [Here](https://developer.tillhub.com/tutorials/products/update/) you can see the call in tillhub documentation.

```go
// Create product data
body := tillhub.CreateProductBody{
	true,
    "product",
    "J&J Testprodukt",
    "UUIDofTheRevenueAccountObjectInTillhub",
    "UUIDofTheTAXAccountObjectInTillhub",
    "JJTEST2021",
	tillhub.CreateProductBodyAttributes{},
    []tillhub.CreateProductBodyCodes{},
    nil,
    nil,
    nil,
    tillhub.CreateProductBodyImages{},
    nil,
    true,
    nil,
    nil,
    nil,
    tillhub.CreateProductBodyPrices{
        []tillhub.CreateProductBodyPricesDefaultPrices{
            {
			    tillhub.CreateProductBodyPricesDefaultPricesAmount{
                    69,
                    82.11,
                },
                3,
                "EUR",
                20,
                2,
            },
        },
        []tillhub.CreateProductBodyPricesBranchPrices{},
    },
    nil,
    true,
    false,
    nil,
    nil,
    tillhub.CreateProductBodyManufacturer{
        nil,
    },
    tillhub.CreateProductBodySupplier{
        nil,
    },
    nil,
    "simple",
    nil,
    []tillhub.CreateProductBodyTags{},
    false,
    nil,
    tillhub.CreateProductBodyConfiguration{
        true,
		tillhub.CreateProductBodyConfigurationPricing{
            false,
        },
    },
    nil,
    nil}

// Create product function
create, err := tillhub.UpdateProduct(body, "productId", "userId", "token")
if err != nil {
    fmt.Println(err)
}

// Print product data
fmt.Println(create)
```

### Delete product

To remove a product you need the ID of the product and you can remove it with this function. [Here](https://developer.tillhub.com/tutorials/products/delete/#getting-started) you can see the call in tillhub documentation.

```go
// Delete an product whith his dependencies
delete, err := DeleteProduct("productId", "userId", "token")
if err != nil {
    fmt.Println(err)
}

// Print return data
fmt.Println(delete)
```