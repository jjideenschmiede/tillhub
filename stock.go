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

// StockBody is to send body data
type StockBody struct {
	Qty      int    `json:"qty"`
	Location string `json:"location"`
}

// StockReturn is to decode the json response
type StockReturn struct {
	Status  int                  `json:"status"`
	Msg     string               `json:"msg"`
	Request StockReturnRequest   `json:"request"`
	Count   int                  `json:"count"`
	Results []StockReturnResults `json:"results"`
}

type StockReturnRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

type StockReturnResults struct {
	Id                        string                          `json:"id"`
	CustomId                  string                          `json:"custom_id"`
	Metadata                  interface{}                     `json:"metadata"`
	CreatedAt                 StockReturnResultsCreatedAt     `json:"created_at"`
	UpdatedAt                 StockReturnResultsUpdatedAt     `json:"updated_at"`
	Name                      string                          `json:"name"`
	Summary                   string                          `json:"summary"`
	Description               string                          `json:"description"`
	Account                   string                          `json:"account"`
	Tax                       string                          `json:"tax"`
	VarClass                  interface{}                     `json:"var_class"`
	Category                  interface{}                     `json:"category"`
	Brand                     string                          `json:"brand"`
	Images                    StockReturnResultsImages        `json:"images"`
	Condition                 interface{}                     `json:"condition"`
	Manufacturer              StockReturnResultsManufacturer  `json:"manufacturer"`
	ProducedAt                interface{}                     `json:"produced_at"`
	PurchasedAt               interface{}                     `json:"purchased_at"`
	ReleasedAt                interface{}                     `json:"released_at"`
	SimilarTo                 []interface{}                   `json:"similar_to"`
	RelatedTo                 []interface{}                   `json:"related_to"`
	Audiences                 []interface{}                   `json:"audiences"`
	Keywords                  []interface{}                   `json:"keywords"`
	Categories                interface{}                     `json:"categories"`
	CustomIds                 interface{}                     `json:"custom_ids"`
	Active                    bool                            `json:"active"`
	Deleted                   bool                            `json:"deleted"`
	Type                      string                          `json:"type"`
	InsertId                  interface{}                     `json:"insert_id"`
	ProductGroup              interface{}                     `json:"product_group"`
	Supplier                  StockReturnResultsSupplier      `json:"supplier"`
	Attributes                interface{}                     `json:"attributes"`
	Barcode                   interface{}                     `json:"barcode"`
	Prices                    StockReturnResultsPrices        `json:"prices"`
	Sku                       string                          `json:"sku"`
	StockMinimum              interface{}                     `json:"stock_minimum"`
	Stockable                 bool                            `json:"stockable"`
	TaxesOptions              interface{}                     `json:"taxes_options"`
	Parent                    interface{}                     `json:"parent"`
	LinkedProducts            interface{}                     `json:"linked_products"`
	Season                    interface{}                     `json:"season"`
	Seasons                   interface{}                     `json:"seasons"`
	Tags                      []interface{}                   `json:"tags"`
	DelegatedTo               interface{}                     `json:"delegated_to"`
	StockMaximum              interface{}                     `json:"stock_maximum"`
	Codes                     []interface{}                   `json:"codes"`
	ReorderPoint              interface{}                     `json:"reorder_point"`
	ReorderQty                interface{}                     `json:"reorder_qty"`
	ExternalReferenceId       string                          `json:"external_reference_id"`
	ClientId                  interface{}                     `json:"client_id"`
	DefaultTitleColor         string                          `json:"default_title_color"`
	StockMode                 string                          `json:"stock_mode"`
	Locations                 interface{}                     `json:"locations"`
	IsService                 bool                            `json:"is_service"`
	ServiceQuestions          interface{}                     `json:"service_questions"`
	ServiceQuestionGroups     interface{}                     `json:"service_question_groups"`
	Configuration             StockReturnResultsConfiguration `json:"configuration"`
	LoyaltyValues             interface{}                     `json:"loyalty_values"`
	Manufacturers             interface{}                     `json:"manufacturers"`
	DelegatedFrom             interface{}                     `json:"delegated_from"`
	DelegateableTo            interface{}                     `json:"delegateable_to"`
	Delegateable              bool                            `json:"delegateable"`
	Delegatable               bool                            `json:"delegatable"`
	DelegatableTo             interface{}                     `json:"delegatable_to"`
	WarrantyNotice            interface{}                     `json:"warranty_notice"`
	RefundPolicy              interface{}                     `json:"refund_policy"`
	Disclaimer                interface{}                     `json:"disclaimer"`
	Policy                    interface{}                     `json:"policy"`
	Online                    interface{}                     `json:"online"`
	ShippingRequired          interface{}                     `json:"shipping_required"`
	ExternalIds               interface{}                     `json:"external_ids"`
	CustomProperties          interface{}                     `json:"custom_properties"`
	Sellable                  bool                            `json:"sellable"`
	Purchasable               bool                            `json:"purchasable"`
	SerialNumberInputRequired bool                            `json:"serial_number_input_required"`
	Stock                     StockReturnResultsStock         `json:"stock"`
	Reason                    interface{}                     `json:"reason"`
	Comments                  interface{}                     `json:"comments"`
}

type StockReturnResultsCreatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type StockReturnResultsUpdatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type StockReturnResultsImages struct {
	First        string `json:"1x"`
	Second       string `json:"2x"`
	Third        string `json:"3x"`
	LiaFirst     string `json:"lia_1x"`
	LiaSecond    string `json:"lia_2x"`
	LiaThird     string `json:"lia_3x"`
	Original     string `json:"original"`
	SquareFirst  string `json:"square_1x"`
	SquareSecond string `json:"square_2x"`
	SquareThird  string `json:"square_3x"`
}

type StockReturnResultsManufacturer struct {
	Iln string `json:"iln"`
}

type StockReturnResultsSupplier struct {
	Sku string `json:"sku"`
}

type StockReturnResultsPrices struct {
	BranchPrices  []interface{}                                   `json:"branch_prices"`
	DefaultPrices []CreateProductReturnResultsPricesDefaultPrices `json:"default_prices,omitempty"`
}

type StockReturnResultsPricesDefaultPrices struct {
	Cost          int                                         `json:"cost"`
	Amount        StockReturnResultsPricesDefaultPricesAmount `json:"amount"`
	Margin        int                                         `json:"margin"`
	Currency      string                                      `json:"currency"`
	PurchasePrice int                                         `json:"purchase_price"`
}

type StockReturnResultsPricesDefaultPricesAmount struct {
	Net   float64 `json:"net"`
	Gross float64 `json:"gross"`
}

type StockReturnResultsConfiguration struct {
	Pricing         StockReturnResultsConfigurationPricing `json:"pricing"`
	AllowZeroPrices bool                                   `json:"allow_zero_prices"`
}

type StockReturnResultsConfigurationPricing struct {
	AllowIsFree bool `json:"allow_is_free"`
}

type StockReturnResultsStock struct {
	Id           string      `json:"id"`
	CreatedAt    string      `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
	Article      interface{} `json:"article"`
	Product      string      `json:"product"`
	Qty          int         `json:"qty"`
	Branch       interface{} `json:"branch"`
	Status       string      `json:"status"`
	Metadata     interface{} `json:"metadata"`
	LocationType interface{} `json:"location_type"`
	Location     string      `json:"location"`
}

// Stock is to update stock of an product
func Stock(data StockBody, productId, accountId, token string) (StockReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v1/products/" + accountId + "/" + productId + "/stock/book"

	// Define client
	client := &http.Client{}

	// Prepare body
	convert, err := json.Marshal(data)
	if err != nil {
		return StockReturn{}, err
	}

	// Set request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(convert))
	if err != nil {
		return StockReturn{}, err
	}

	// Set header
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+token)

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return StockReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json return
	var decode StockReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return StockReturn{}, err
	}

	// Return data
	return decode, nil

}
