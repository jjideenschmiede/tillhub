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
data, err := tillhub.AuthKey("accountId", "apiKey")
if err != nil {
    fmt.Println(err)
}

// Print token
fmt.Println(data.token)
```

### Tax

To find out the tax ids and get more information about the tax you can use the following function. [Here](https://developer.tillhub.com/tutorials/products/create/#prerequisits) you can see the call in tillhub documentation.

```go
// Get tax information
tax, err := Tax("accountId", "token")
if err != nil {
    fmt.Println(err)
}

// Print tax information
fmt.Println(tax)
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
    tillhub.CreateProductBodyImages{
        "https://jj-ideenschmiede.de/test.png",
		"https://jj-ideenschmiede.de/test.png",
		"https://jj-ideenschmiede.de/test.png",
		"https://jj-ideenschmiede.de/test.png",
		"https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
    },
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
    tillhub.CreateProductBodyImages{
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
        "https://jj-ideenschmiede.de/test.png",
    },
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
delete, err := tillhub.DeleteProduct("productId", "userId", "token")
if err != nil {
    fmt.Println(err)
}

// Print return data
fmt.Println(delete)
```

### Read all product groups

If you want to read out all product groups, you can do it with this function. Both simple and variable products are returned. [Here](https://api.tillhub.com/api/docs#tag/product_groups/paths/~1api~1v0~1product_groups~1{clientAccountID}/get) you can see the call in tillhub documentation.

```go
// Get all product groups
groups, err := tillhub.ProductGroups("accountId", "token")
if err != nil {
fmt.Println(err)
}

// Print all product groups
fmt.Println(groups)
```

### Create new product group

To define a new product group the following function can be used. You can find the documentation [here](https://api.tillhub.com/api/docs#tag/product_groups/paths/~1api~1v0~1product_groups~1{clientAccountID}/post).

```go
// Create product body data
body := tillhub.CreateProductGroupBody{
    true,
    "#87cef9",
    "newGroup",
    "00001",
    "f85910a6-94c4-40bd-a0ac-386ef880c857",
    "673ba8d1-f544-48be-b432-122633e92715",
	tillhub.CreateProductGroupBodyImages{
        "https://storage.googleapis.com/tillhub-api-images/dd5a5b03-69fd-4747-bd6c-2ecf22ec78c1/productGroups/1a743bf7-10af-4897-891e-13c9924441de_1x.png",
        "https://storage.googleapis.com/tillhub-api-images/dd5a5b03-69fd-4747-bd6c-2ecf22ec78c1/productGroups/1a743bf7-10af-4897-891e-13c9924441de_1x.png",
        "https://storage.googleapis.com/tillhub-api-images/dd5a5b03-69fd-4747-bd6c-2ecf22ec78c1/productGroups/1a743bf7-10af-4897-891e-13c9924441de_1x.png",
        "https://storage.googleapis.com/tillhub-api-images/dd5a5b03-69fd-4747-bd6c-2ecf22ec78c1/productGroups/1a743bf7-10af-4897-891e-13c9924441de_1x.png",
        "https://storage.googleapis.com/tillhub-api-images/dd5a5b03-69fd-4747-bd6c-2ecf22ec78c1/productGroups/1a743bf7-10af-4897-891e-13c9924441de_1x.png",
    },
}

// Save new product group
group, err := tillhub.CreateProductGroup(body, "accountId", "token")
if err != nil {
    fmt.Println(err)
}

// Print return
fmt.Println(group)
```

### Update a product group

To update a product group the following function can be used. You can find the documentation [here](https://api.tillhub.com/api/docs#tag/product_groups/paths/~1api~1v0~1product_groups~1{clientAccountID}~1{productGroupID}/put).

```go
// Create product body data
body := tillhub.CreateProductGroupBody{
    true,
    "#87cef9",
    "newGroup",
    "00001",
    "f85910a6-94c4-40bd-a0ac-386ef880c857",
    "673ba8d1-f544-48be-b432-122633e92715",
	tillhub.CreateProductGroupBodyImages{
        "https://storage.googleapis.com/tillhub-api-images/dd5a5b03-69fd-4747-bd6c-2ecf22ec78c1/productGroups/1a743bf7-10af-4897-891e-13c9924441de_1x.png",
        "https://storage.googleapis.com/tillhub-api-images/dd5a5b03-69fd-4747-bd6c-2ecf22ec78c1/productGroups/1a743bf7-10af-4897-891e-13c9924441de_1x.png",
        "https://storage.googleapis.com/tillhub-api-images/dd5a5b03-69fd-4747-bd6c-2ecf22ec78c1/productGroups/1a743bf7-10af-4897-891e-13c9924441de_1x.png",
        "https://storage.googleapis.com/tillhub-api-images/dd5a5b03-69fd-4747-bd6c-2ecf22ec78c1/productGroups/1a743bf7-10af-4897-891e-13c9924441de_1x.png",
        "https://storage.googleapis.com/tillhub-api-images/dd5a5b03-69fd-4747-bd6c-2ecf22ec78c1/productGroups/1a743bf7-10af-4897-891e-13c9924441de_1x.png",
    },
}

// Update product group
group, err := tillhub.UpdateProductGroup(body, "groupId", "accountId", "token")
if err != nil {
    fmt.Println(err)
}

// Print return
fmt.Println(group)
```

### Delete product group

To delete a product group the following function can be used. You can find the documentation [here](https://api.tillhub.com/api/docs#tag/product_groups/paths/~1api~1v0~1product_groups~1{clientAccountID}~1{product_groupID}/delete).

```go
// Update product group
delete, err := tillhub.DeleteProductGroup("groupId", "accountId", "token")
if err != nil {
    fmt.Println(err)
}

// Print return
fmt.Println(delete)
```