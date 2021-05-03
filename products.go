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
	PurchasePrice interface{}                                  `json:"purchase_price"`
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
