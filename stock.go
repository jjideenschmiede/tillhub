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

// CreateStockBody is to send body data
type CreateStockBody struct {
	Qty      int    `json:"qty"`
	Location string `json:"location"`
}

// CreateStockReturn is to decode the json response
type CreateStockReturn struct {
	Status  int                        `json:"status"`
	Msg     string                     `json:"msg"`
	Request CreateStockReturnRequest   `json:"request"`
	Count   int                        `json:"count"`
	Results []CreateStockReturnResults `json:"results"`
}

type CreateStockReturnRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

type CreateStockReturnResults struct {
	Id                        string                                `json:"id"`
	CustomId                  string                                `json:"custom_id"`
	Metadata                  interface{}                           `json:"metadata"`
	CreatedAt                 CreateStockReturnResultsCreatedAt     `json:"created_at"`
	UpdatedAt                 CreateStockReturnResultsUpdatedAt     `json:"updated_at"`
	Name                      string                                `json:"name"`
	Summary                   string                                `json:"summary"`
	Description               string                                `json:"description"`
	Account                   string                                `json:"account"`
	Tax                       string                                `json:"tax"`
	VarClass                  interface{}                           `json:"var_class"`
	Category                  interface{}                           `json:"category"`
	Brand                     string                                `json:"brand"`
	Images                    CreateStockReturnResultsImages        `json:"images"`
	Condition                 interface{}                           `json:"condition"`
	Manufacturer              CreateStockReturnResultsManufacturer  `json:"manufacturer"`
	ProducedAt                interface{}                           `json:"produced_at"`
	PurchasedAt               interface{}                           `json:"purchased_at"`
	ReleasedAt                interface{}                           `json:"released_at"`
	SimilarTo                 []interface{}                         `json:"similar_to"`
	RelatedTo                 []interface{}                         `json:"related_to"`
	Audiences                 []interface{}                         `json:"audiences"`
	Keywords                  []interface{}                         `json:"keywords"`
	Categories                interface{}                           `json:"categories"`
	CustomIds                 interface{}                           `json:"custom_ids"`
	Active                    bool                                  `json:"active"`
	Deleted                   bool                                  `json:"deleted"`
	Type                      string                                `json:"type"`
	InsertId                  interface{}                           `json:"insert_id"`
	ProductGroup              interface{}                           `json:"product_group"`
	Supplier                  CreateStockReturnResultsSupplier      `json:"supplier"`
	Attributes                interface{}                           `json:"attributes"`
	Barcode                   interface{}                           `json:"barcode"`
	Prices                    CreateStockReturnResultsPrices        `json:"prices"`
	Sku                       string                                `json:"sku"`
	StockMinimum              interface{}                           `json:"stock_minimum"`
	Stockable                 bool                                  `json:"stockable"`
	TaxesOptions              interface{}                           `json:"taxes_options"`
	Parent                    interface{}                           `json:"parent"`
	LinkedProducts            interface{}                           `json:"linked_products"`
	Season                    interface{}                           `json:"season"`
	Seasons                   interface{}                           `json:"seasons"`
	Tags                      []interface{}                         `json:"tags"`
	DelegatedTo               interface{}                           `json:"delegated_to"`
	StockMaximum              interface{}                           `json:"stock_maximum"`
	Codes                     []interface{}                         `json:"codes"`
	ReorderPoint              interface{}                           `json:"reorder_point"`
	ReorderQty                interface{}                           `json:"reorder_qty"`
	ExternalReferenceId       string                                `json:"external_reference_id"`
	ClientId                  interface{}                           `json:"client_id"`
	DefaultTitleColor         string                                `json:"default_title_color"`
	StockMode                 string                                `json:"stock_mode"`
	Locations                 interface{}                           `json:"locations"`
	IsService                 bool                                  `json:"is_service"`
	ServiceQuestions          interface{}                           `json:"service_questions"`
	ServiceQuestionGroups     interface{}                           `json:"service_question_groups"`
	Configuration             CreateStockReturnResultsConfiguration `json:"configuration"`
	LoyaltyValues             interface{}                           `json:"loyalty_values"`
	Manufacturers             interface{}                           `json:"manufacturers"`
	DelegatedFrom             interface{}                           `json:"delegated_from"`
	DelegateableTo            interface{}                           `json:"delegateable_to"`
	Delegateable              bool                                  `json:"delegateable"`
	Delegatable               bool                                  `json:"delegatable"`
	DelegatableTo             interface{}                           `json:"delegatable_to"`
	WarrantyNotice            interface{}                           `json:"warranty_notice"`
	RefundPolicy              interface{}                           `json:"refund_policy"`
	Disclaimer                interface{}                           `json:"disclaimer"`
	Policy                    interface{}                           `json:"policy"`
	Online                    interface{}                           `json:"online"`
	ShippingRequired          interface{}                           `json:"shipping_required"`
	ExternalIds               interface{}                           `json:"external_ids"`
	CustomProperties          interface{}                           `json:"custom_properties"`
	Sellable                  bool                                  `json:"sellable"`
	Purchasable               bool                                  `json:"purchasable"`
	SerialNumberInputRequired bool                                  `json:"serial_number_input_required"`
	Stock                     CreateStockReturnResultsStock         `json:"stock"`
	Reason                    interface{}                           `json:"reason"`
	Comments                  interface{}                           `json:"comments"`
}

type CreateStockReturnResultsCreatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type CreateStockReturnResultsUpdatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type CreateStockReturnResultsImages struct {
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

type CreateStockReturnResultsManufacturer struct {
	Iln string `json:"iln"`
}

type CreateStockReturnResultsSupplier struct {
	Sku string `json:"sku"`
}

type CreateStockReturnResultsPrices struct {
	BranchPrices  []interface{}                                 `json:"branch_prices"`
	DefaultPrices []CreateStockReturnResultsPricesDefaultPrices `json:"default_prices,omitempty"`
}

type CreateStockReturnResultsPricesDefaultPrices struct {
	Cost          int                                               `json:"cost"`
	Amount        CreateStockReturnResultsPricesDefaultPricesAmount `json:"amount"`
	Margin        int                                               `json:"margin"`
	Currency      string                                            `json:"currency"`
	PurchasePrice int                                               `json:"purchase_price"`
}

type CreateStockReturnResultsPricesDefaultPricesAmount struct {
	Net   float64 `json:"net"`
	Gross float64 `json:"gross"`
}

type CreateStockReturnResultsConfiguration struct {
	Pricing         CreateStockReturnResultsConfigurationPricing `json:"pricing"`
	AllowZeroPrices bool                                         `json:"allow_zero_prices"`
}

type CreateStockReturnResultsConfigurationPricing struct {
	AllowIsFree bool `json:"allow_is_free"`
}

type CreateStockReturnResultsStock struct {
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

// DeleteStockReturn is to decode json data
type DeleteStockReturn struct {
	Status  int                      `json:"status"`
	Msg     string                   `json:"msg"`
	Request DeleteStockReturnRequest `json:"request"`
}

type DeleteStockReturnRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

// CreateStock is to update stock of an product
func CreateStock(data CreateStockBody, productId, accountId, token string) (CreateStockReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v1/products/" + accountId + "/" + productId + "/stock/book"

	// Define client
	client := &http.Client{}

	// Prepare body
	convert, err := json.Marshal(data)
	if err != nil {
		return CreateStockReturn{}, err
	}

	// Set request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(convert))
	if err != nil {
		return CreateStockReturn{}, err
	}

	// Set header
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+token)

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return CreateStockReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json return
	var decode CreateStockReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CreateStockReturn{}, err
	}

	// Return data
	return decode, nil

}

// DeleteStock is to clear the stock
func DeleteStock(stockId, accountId, token string) (DeleteStockReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v0/stock/" + accountId + "/" + stockId

	// Define client
	client := &http.Client{}

	// Set request
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return DeleteStockReturn{}, err
	}

	// Set header
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+token)

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return DeleteStockReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json return
	var decode DeleteStockReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return DeleteStockReturn{}, err
	}

	// Return data
	return decode, nil

}
