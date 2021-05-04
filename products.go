//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede UG (haftungsbeschränkt) <info@jj-ideenschmiede.de>
//
// This file is part of tillhub.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package tillhub

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// ReadProductsReturn is to decode read product data return
type ReadProductsReturn struct {
	Status  int                   `json:"status"`
	Msg     string                `json:"msg"`
	Request ReadProductsRequest   `json:"request"`
	Cursor  ReadProductsCursor    `json:"cursor"`
	Count   int                   `json:"count"`
	Results []ReadProductsResults `json:"results"`
}

type ReadProductsRequest struct {
	Host   string                    `json:"host"`
	Id     string                    `json:"id"`
	Cursor ReadProductsRequestCursor `json:"cursor"`
}

type ReadProductsRequestCursor struct {
	First ReadProductsRequestCursorFirst `json:"first"`
	Page  int                            `json:"page"`
}

type ReadProductsRequestCursorFirst struct {
	Id          string `json:"id"`
	UpdatedAt   string `json:"updated_at"`
	Page        int    `json:"page"`
	CursorField string `json:"cursor_field"`
}

type ReadProductsCursor struct {
	First string `json:"first"`
	Self  string `json:"self"`
}

type ReadProductsResults struct {
	Children                   []ReadProductsResultsChildren    `json:"children,omitempty"`
	PriceBookEntries           interface{}                      `json:"price_book_entries"`
	Id                         string                           `json:"id"`
	CustomId                   string                           `json:"custom_id"`
	Metadata                   interface{}                      `json:"metadata"`
	CreatedAt                  ReadProductsResultsCreatedAt     `json:"created_at"`
	UpdatedAt                  ReadProductsResultsUpdatedAt     `json:"updated_at"`
	Name                       string                           `json:"name"`
	Summary                    interface{}                      `json:"summary"`
	Description                interface{}                      `json:"description"`
	Account                    string                           `json:"account"`
	Tax                        string                           `json:"tax"`
	VatClass                   interface{}                      `json:"vat_class"`
	Category                   interface{}                      `json:"category"`
	Brand                      interface{}                      `json:"brand"`
	Images                     interface{}                      `json:"images"`
	Condition                  interface{}                      `json:"condition"`
	Manufacturer               ReadProductsResultsManufacturer  `json:"manufacturer"`
	ProducedAt                 interface{}                      `json:"produced_at"`
	PurchasedAt                interface{}                      `json:"purchased_at"`
	ReleasedAt                 interface{}                      `json:"released_at"`
	SimilarTo                  []interface{}                    `json:"similar_to"`
	RelatedTo                  []interface{}                    `json:"related_to"`
	Audiences                  []interface{}                    `json:"audiences"`
	Keywords                   []interface{}                    `json:"keywords"`
	Categories                 []interface{}                    `json:"categories"`
	CustomIds                  interface{}                      `json:"custom_ids"`
	Active                     bool                             `json:"active"`
	Deleted                    bool                             `json:"deleted"`
	Type                       string                           `json:"type"`
	InsertId                   interface{}                      `json:"insert_id"`
	ProductGroup               interface{}                      `json:"product_group"`
	Supplier                   ReadProductsResultsSupplier      `json:"supplier"`
	Attributes                 interface{}                      `json:"attributes"`
	Barcode                    interface{}                      `json:"barcode"`
	Prices                     ReadProductsResultsPrices        `json:"prices"`
	Sku                        interface{}                      `json:"sku"`
	StockMinimum               interface{}                      `json:"stock_minimum"`
	Stockable                  bool                             `json:"stockable"`
	TaxesOptions               interface{}                      `json:"taxes_options"`
	Parent                     string                           `json:"parent"`
	LinkedProducts             interface{}                      `json:"linked_products"`
	Season                     interface{}                      `json:"season"`
	Seasons                    interface{}                      `json:"seasons"`
	Tags                       interface{}                      `json:"tags"`
	DelegatedTo                interface{}                      `json:"delegated_to"`
	StockMaximum               interface{}                      `json:"stock_maximum"`
	Codes                      interface{}                      `json:"codes"`
	ReorderPoint               interface{}                      `json:"reorder_point"`
	ReorderQty                 interface{}                      `json:"reorder_qty"`
	ExternalReferenceId        interface{}                      `json:"external_reference_id"`
	ClientId                   interface{}                      `json:"client_id"`
	DefaultTitleColor          interface{}                      `json:"default_title_color"`
	Discountable               bool                             `json:"discountable"`
	Linkable                   bool                             `json:"linkable"`
	StockConfigurationLocation interface{}                      `json:"stock_configuration_location"`
	StockMode                  interface{}                      `json:"stock_mode"`
	Locations                  interface{}                      `json:"locations"`
	IsService                  bool                             `json:"is_service"`
	ServiceQuestions           interface{}                      `json:"service_questions"`
	ServiceQuestionGroup       interface{}                      `json:"service_question_group"`
	Configuration              ReadProductsResultsConfiguration `json:"configuration"`
	LoyaltyValues              interface{}                      `json:"loyalty_values"`
	Manufactures               interface{}                      `json:"manufactures"`
	DelegatedFrom              interface{}                      `json:"delegated_from"`
	DelegateableTo             interface{}                      `json:"delegateable_to"`
	Delegateable               bool                             `json:"delegateable"`
	Delegatable                bool                             `json:"delegatable"`
	DelegatableTo              interface{}                      `json:"delegatable_to"`
	WarrantyNotice             interface{}                      `json:"warranty_notice"`
	RefundPolicy               interface{}                      `json:"refund_policy"`
	Disclaimer                 interface{}                      `json:"disclaimer"`
	Policy                     interface{}                      `json:"policy"`
	Online                     bool                             `json:"online"`
	ShippingRequired           bool                             `json:"shipping_required"`
	ExternalIds                interface{}                      `json:"external_ids"`
	CustomProperties           interface{}                      `json:"custom_properties"`
	Sellable                   bool                             `json:"sellable"`
	Purchasable                bool                             `json:"purchasable"`
	SerialNumberInputRequired  bool                             `json:"serial_number_input_required"`
	Options                    []interface{}                    `json:"options"`
}

type ReadProductsResultsChildren struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	CustomId   string      `json:"custom_id"`
	Attributes interface{} `json:"attributes"`
	Locations  interface{} `json:"locations"`
	Images     interface{} `json:"images"`
}

type ReadProductsResultsCreatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type ReadProductsResultsUpdatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type ReadProductsResultsManufacturer struct {
	Iln interface{} `json:"iln"`
}

type ReadProductsResultsSupplier struct {
	Sku interface{} `json:"sku"`
}

type ReadProductsResultsPrices struct {
	BasePrices    []interface{}                            `json:"base_prices"`
	BranchPrices  []interface{}                            `json:"branch_prices"`
	scaledPrices  []interface{}                            `json:"scaled_prices"`
	DefaultPrices []ReadProductsResultsPricesDefaultPrices `json:"default_prices"`
}

type ReadProductsResultsPricesDefaultPrices struct {
	Cost          int                                          `json:"cost"`
	Amount        ReadProductsResultsPricesDefaultPricesAmount `json:"amount"`
	Margin        int                                          `json:"margin"`
	Currency      string                                       `json:"currency"`
	PurchasePrice int                                          `json:"purchase_price"`
}

type ReadProductsResultsPricesDefaultPricesAmount struct {
	Net   float64 `json:"net"`
	Gross float64 `json:"gross"`
}

type ReadProductsResultsConfiguration struct {
	Pricing        ReadProductsResultsConfigurationPricing `json:"pricing"`
	AllowZeroPrice bool                                    `json:"allow_zero_price"`
}

type ReadProductsResultsConfigurationPricing struct {
	AllowIsFree bool `json:"allow_is_free"`
}

// CreateProductBody is to send json product data
type CreateProductBody struct {
	Active              bool                           `json:"active"`
	Type                string                         `json:"type"`
	Name                string                         `json:"name"`
	Account             string                         `json:"account"`
	Tax                 string                         `json:"tax"`
	CustomId            string                         `json:"custom_id"`
	Attributes          CreateProductBodyAttributes    `json:"attributes"`
	Codes               []CreateProductBodyCodes       `json:"codes"`
	Barcode             interface{}                    `json:"barcode"`
	ProductGroup        interface{}                    `json:"product_group"`
	Categories          interface{}                    `json:"categories"`
	Images              CreateProductBodyImages        `json:"images"`
	StockMinimum        interface{}                    `json:"stock_minimum"`
	Stockable           bool                           `json:"stockable"`
	Sku                 interface{}                    `json:"sku"`
	Description         interface{}                    `json:"description"`
	Summary             interface{}                    `json:"summary"`
	Prices              CreateProductBodyPrices        `json:"prices"`
	ExternalReferenceId interface{}                    `json:"external_reference_id"`
	Discountable        bool                           `json:"discountable"`
	Linkable            bool                           `json:"linkable"`
	LinkedProducts      interface{}                    `json:"linked_products"`
	Brand               interface{}                    `json:"brand"`
	Manufacturer        CreateProductBodyManufacturer  `json:"manufacturer"`
	Supplier            CreateProductBodySupplier      `json:"supplier"`
	Locations           interface{}                    `json:"locations"`
	StockMode           string                         `json:"stock_mode"`
	ReorderPoint        interface{}                    `json:"reorder_point"`
	Tags                []CreateProductBodyTags        `json:"tags"`
	IsService           bool                           `json:"is_service"`
	ServiceQuestions    interface{}                    `json:"service_questions"`
	Configuration       CreateProductBodyConfiguration `json:"configuration"`
	Online              interface{}                    `json:"online"`
	ShippingRequired    interface{}                    `json:"shipping_required"`
}

type CreateProductBodyAttributes struct {
}

type CreateProductBodyCodes struct {
}

type CreateProductBodyImages struct {
}

type CreateProductBodyPrices struct {
	DefaultPrices []CreateProductBodyPricesDefaultPrices `json:"default_prices"`
	BranchPrices  []CreateProductBodyPricesBranchPrices  `json:"branch_prices"`
}

type CreateProductBodyPricesDefaultPrices struct {
	Amount        CreateProductBodyPricesDefaultPricesAmount `json:"amount"`
	PurchasePrice int                                        `json:"purchase_price"`
	Currency      string                                     `json:"currency"`
	Cost          int                                        `json:"cost"`
	Margin        int                                        `json:"margin"`
}

type CreateProductBodyPricesBranchPrices struct {
}

type CreateProductBodyPricesDefaultPricesAmount struct {
	Net   float64 `json:"net"`
	Gross float64 `json:"gross"`
}

type CreateProductBodyManufacturer struct {
	Iln interface{} `json:"iln"`
}

type CreateProductBodySupplier struct {
	Sku interface{} `json:"sku"`
}

type CreateProductBodyTags struct {
}

type CreateProductBodyConfiguration struct {
	AllowZeroPrices bool                                  `json:"allow_zero_prices"`
	Pricing         CreateProductBodyConfigurationPricing `json:"pricing"`
}

type CreateProductBodyConfigurationPricing struct {
	AllowIsFree bool `json:"allow_is_free"`
}

// CreateProductReturn is to decode json return
type CreateProductReturn struct {
	Status  int                          `json:"status"`
	Msg     string                       `json:"msg"`
	Errors  []interface{}                `json:"errors"`
	Request CreateProductReturnRequest   `json:"request"`
	Count   int                          `json:"count"`
	Results []CreateProductReturnResults `json:"results"`
}

type CreateProductReturnRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

type CreateProductReturnResults struct {
	Id                         string                                  `json:"id"`
	CustomId                   string                                  `json:"custom_id"`
	Metadata                   interface{}                             `json:"metadata"`
	CreatedAt                  CreateProductReturnResultsCreatedAt     `json:"created_at"`
	UpdatedAt                  CreateProductReturnResultsUpdatedAt     `json:"updated_at"`
	Name                       string                                  `json:"name"`
	Summary                    interface{}                             `json:"summary"`
	Description                interface{}                             `json:"description"`
	Account                    string                                  `json:"account"`
	Tax                        string                                  `json:"tax"`
	VatClass                   interface{}                             `json:"vat_class"`
	Category                   interface{}                             `json:"category"`
	Brand                      interface{}                             `json:"brand"`
	Images                     interface{}                             `json:"images"`
	Condition                  interface{}                             `json:"condition"`
	Manufacturer               CreateProductReturnResultsManufactor    `json:"manufacturer"`
	ProducedAt                 interface{}                             `json:"produced_at"`
	PurchasedAt                interface{}                             `json:"purchased_at"`
	ReleasedAt                 interface{}                             `json:"released_at"`
	SimilarTo                  []interface{}                           `json:"similar_to"`
	RelatedTo                  []interface{}                           `json:"related_to"`
	Audiences                  []interface{}                           `json:"audiences"`
	Keywords                   []interface{}                           `json:"keywords"`
	Categories                 interface{}                             `json:"categories"`
	CustomIds                  interface{}                             `json:"custom_ids"`
	Active                     bool                                    `json:"active"`
	Deleted                    bool                                    `json:"deleted"`
	Type                       string                                  `json:"type"`
	InsertId                   interface{}                             `json:"insert_id"`
	ProductGroup               interface{}                             `json:"product_group"`
	Supplier                   CreateProductReturnResultsSupplier      `json:"supplier"`
	Attributes                 interface{}                             `json:"attributes"`
	Barcode                    interface{}                             `json:"barcode"`
	Prices                     CreateProductReturnResultsPrices        `json:"prices"`
	Sku                        interface{}                             `json:"sku"`
	StockMinimum               interface{}                             `json:"stock_minimum"`
	Stockable                  interface{}                             `json:"stockable"`
	TaxesOptions               interface{}                             `json:"taxes_options"`
	Parent                     string                                  `json:"parent"`
	LinkedProducts             interface{}                             `json:"linked_products"`
	Season                     interface{}                             `json:"season"`
	Seasons                    interface{}                             `json:"seasons"`
	Tags                       []interface{}                           `json:"tags"`
	DelegatedTo                interface{}                             `json:"delegated_to"`
	StockMaximum               interface{}                             `json:"stock_maximum"`
	Codes                      []interface{}                           `json:"codes"`
	ReorderPoint               interface{}                             `json:"reorder_point"`
	ReorderQty                 interface{}                             `json:"reorder_qty"`
	ExternalReferenceId        interface{}                             `json:"external_reference_id"`
	ClientId                   interface{}                             `json:"client_id"`
	DefaultTitleColor          interface{}                             `json:"default_title_color"`
	Discountable               bool                                    `json:"discountable"`
	Linkable                   bool                                    `json:"linkable"`
	StockConfigurationLocation interface{}                             `json:"stock_configuration_location"`
	StockMode                  interface{}                             `json:"stock_mode"`
	Locations                  interface{}                             `json:"locations"`
	IsService                  bool                                    `json:"is_service"`
	ServiceQuestions           interface{}                             `json:"service_questions"`
	ServiceQuestionGroup       interface{}                             `json:"service_question_group"`
	Configuration              CreateProductReturnResultsConfiguration `json:"configuration"`
	LoyaltyValues              interface{}                             `json:"loyalty_values"`
	Manufactures               interface{}                             `json:"manufactures"`
	DelegatedFrom              interface{}                             `json:"delegated_from"`
	DelegateableTo             interface{}                             `json:"delegateable_to"`
	Delegateable               bool                                    `json:"delegateable"`
	Delegatable                bool                                    `json:"delegatable"`
	DelegatableTo              interface{}                             `json:"delegatable_to"`
	WarrantyNotice             interface{}                             `json:"warranty_notice"`
	RefundPolicy               interface{}                             `json:"refund_policy"`
	Disclaimer                 interface{}                             `json:"disclaimer"`
	Policy                     interface{}                             `json:"policy"`
	Online                     interface{}                             `json:"online"`
	ShippingRequired           interface{}                             `json:"shipping_required"`
	ExternalIds                interface{}                             `json:"external_ids"`
	CustomProperties           interface{}                             `json:"custom_properties"`
	Sellable                   bool                                    `json:"sellable"`
	Purchasable                bool                                    `json:"purchasable"`
	SerialNumberInputRequired  bool                                    `json:"serial_number_input_required"`
	StockInfo                  interface{}                             `json:"stock_info"`
	I18n                       interface{}                             `json:"i18n"`
	Revisions                  interface{}                             `json:"revisions"`
}

type CreateProductReturnResultsCreatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type CreateProductReturnResultsUpdatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type CreateProductReturnResultsManufactor struct {
	Iln interface{} `json:"iln"`
}

type CreateProductReturnResultsSupplier struct {
	Sku interface{} `json:"sku"`
}

type CreateProductReturnResultsPrices struct {
	BasePrices    []interface{}                                   `json:"base_prices"`
	BranchPrices  []interface{}                                   `json:"branch_prices"`
	scaledPrices  []interface{}                                   `json:"scaled_prices"`
	DefaultPrices []CreateProductReturnResultsPricesDefaultPrices `json:"default_prices"`
}

type CreateProductReturnResultsPricesDefaultPrices struct {
	Cost          int                                                 `json:"cost"`
	Amount        CreateProductReturnResultsPricesDefaultPricesAmount `json:"amount"`
	Margin        int                                                 `json:"margin"`
	Currency      string                                              `json:"currency"`
	PurchasePrice int                                                 `json:"purchase_price"`
}

type CreateProductReturnResultsPricesDefaultPricesAmount struct {
	Net   float64 `json:"net"`
	Gross float64 `json:"gross"`
}

type CreateProductReturnResultsConfiguration struct {
	Pricing        CreateProductReturnResultsConfigurationPricing `json:"pricing"`
	AllowZeroPrice bool                                           `json:"allow_zero_price"`
}

type CreateProductReturnResultsConfigurationPricing struct {
	AllowIsFree bool `json:"allow_is_free"`
}

// DeleteProductReturn is to decode json response
type DeleteProductReturn struct {
	Status  int                        `json:"status"`
	Msg     string                     `json:"msg"`
	Request DeleteProductReturnRequest `json:"request"`
}

type DeleteProductReturnRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

// ReadProducts is to read products
func ReadProducts(userId, token string) (ReadProductsReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v1/products/" + userId + "/?deleted=false"

	// Define client
	client := &http.Client{}

	// Define request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ReadProductsReturn{}, err
	}

	// Set header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", token)

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return ReadProductsReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode response
	var decode ReadProductsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ReadProductsReturn{}, err
	}

	// Return data
	return decode, nil

}

// CreateProduct is to create an new product
func CreateProduct(data CreateProductBody, userId, token string) (CreateProductReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v1/products/" + userId

	// Define client
	client := &http.Client{}

	// Convert body data
	convert, err := json.Marshal(data)
	if err != nil {
		return CreateProductReturn{}, err
	}

	// Define request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(convert))
	if err != nil {
		return CreateProductReturn{}, err
	}

	// Set header
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", token)

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return CreateProductReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode response
	var decode CreateProductReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CreateProductReturn{}, err
	}

	// Return data
	return decode, nil

}

// DeleteProduct is to delete an product and his dependencies
func DeleteProduct(productId, userId, token string) (DeleteProductReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v1/products/" + userId + "/" + productId + "?delete_dependencies=true"

	// Define client
	client := &http.Client{}

	// Define request
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return DeleteProductReturn{}, err
	}

	// Set header
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", token)

	// Set response and send request
	response, err := client.Do(request)
	if err != nil {
		return DeleteProductReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json response
	var decode DeleteProductReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return DeleteProductReturn{}, err
	}

	// Return data
	return decode, nil

}
