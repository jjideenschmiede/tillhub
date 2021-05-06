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
	Summary                    string                           `json:"summary"`
	Description                string                           `json:"description"`
	Account                    string                           `json:"account"`
	Tax                        string                           `json:"tax"`
	VatClass                   interface{}                      `json:"vat_class"`
	Category                   interface{}                      `json:"category"`
	Brand                      string                           `json:"brand"`
	Images                     ReadProductsResultsImages        `json:"images,omitempty"`
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
	Sku                        string                           `json:"sku"`
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
	ExternalReferenceId        string                           `json:"external_reference_id"`
	ClientId                   interface{}                      `json:"client_id"`
	DefaultTitleColor          string                           `json:"default_title_color"`
	Discountable               bool                             `json:"discountable"`
	Linkable                   bool                             `json:"linkable"`
	StockConfigurationLocation interface{}                      `json:"stock_configuration_location"`
	StockMode                  string                           `json:"stock_mode"`
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

type ReadProductsResultsImages struct {
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

type ReadProductsResultsManufacturer struct {
	Iln string `json:"iln"`
}

type ReadProductsResultsSupplier struct {
	Sku string `json:"sku"`
}

type ReadProductsResultsPrices struct {
	BasePrices    []interface{}                            `json:"base_prices"`
	BranchPrices  []interface{}                            `json:"branch_prices"`
	scaledPrices  []interface{}                            `json:"scaled_prices"`
	DefaultPrices []ReadProductsResultsPricesDefaultPrices `json:"default_prices,omitempty"`
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
	Images              CreateProductBodyImages        `json:"images,omitempty"`
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
	ScaledPrices  []interface{}                                   `json:"scaled_prices"`
	DefaultPrices []CreateProductReturnResultsPricesDefaultPrices `json:"default_prices,omitempty"`
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

// UpdateProductReturn is to decode the json return
type UpdateProductReturn struct {
	Status  int                          `json:"status"`
	Msg     string                       `json:"msg"`
	Request UpdateProductReturnRequest   `json:"request"`
	Count   int                          `json:"count"`
	Results []UpdateProductReturnResults `json:"results"`
}

type UpdateProductReturnRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

type UpdateProductReturnResults struct {
	Id                         string                                  `json:"id"`
	Sku                        string                                  `json:"sku"`
	Tax                        string                                  `json:"tax"`
	I18n                       interface{}                             `json:"i18n"`
	Name                       string                                  `json:"name"`
	Tags                       []interface{}                           `json:"tags"`
	Type                       string                                  `json:"type"`
	Brand                      string                                  `json:"brand"`
	Codes                      []interface{}                           `json:"codes"`
	Active                     bool                                    `json:"active"`
	Images                     UpdateProductReturnResultsImages        `json:"images,omitempty"`
	Online                     interface{}                             `json:"online"`
	Parent                     interface{}                             `json:"parent"`
	Policy                     interface{}                             `json:"policy"`
	Prices                     UpdateProductReturnResultsPrices        `json:"prices"`
	Season                     interface{}                             `json:"season"`
	Account                    string                                  `json:"account"`
	Barcode                    interface{}                             `json:"barcode"`
	Deleted                    bool                                    `json:"deleted"`
	Seasons                    interface{}                             `json:"seasons"`
	Summary                    string                                  `json:"summary"`
	Category                   interface{}                             `json:"category"`
	Keywords                   []interface{}                           `json:"keywords"`
	Linkable                   bool                                    `json:"linkable"`
	Metadata                   interface{}                             `json:"metadata"`
	Sellable                   bool                                    `json:"sellable"`
	Supplier                   UpdateProductReturnResultsSupplier      `json:"supplier"`
	Audiences                  []interface{}                           `json:"audiences"`
	ClientId                   interface{}                             `json:"client_id"`
	Condition                  interface{}                             `json:"condition"`
	CustomId                   string                                  `json:"custom_id"`
	InsertId                   interface{}                             `json:"insert_id"`
	Locations                  interface{}                             `json:"locations"`
	Revisions                  interface{}                             `json:"revisions"`
	Stockable                  bool                                    `json:"stockable"`
	VatClass                   interface{}                             `json:"vat_class"`
	Attributes                 interface{}                             `json:"attributes"`
	Categories                 interface{}                             `json:"categories"`
	CreatedAt                  UpdateProductReturnResultsCreatedAt     `json:"created_at"`
	Delegatable                bool                                    `json:"delegatable"`
	Description                string                                  `json:"description"`
	ProducedAt                 interface{}                             `json:"produced_at"`
	Purchasable                bool                                    `json:"purchasable"`
	ReleasedAt                 interface{}                             `json:"released_at"`
	ReorderQty                 interface{}                             `json:"reorder_qty"`
	Delegateable               bool                                    `json:"delegateable"`
	DelegatedTo                interface{}                             `json:"delegated_to"`
	Discountable               bool                                    `json:"discountable"`
	ExternalIds                interface{}                             `json:"external_ids"`
	Manufacturer               UpdateProductReturnResultsManufacturer  `json:"manufacturer"`
	PurchasedAt                interface{}                             `json:"purchased_at"`
	Configuration              UpdateProductReturnResultsConfiguration `json:"configuration"`
	Manufacturers              interface{}                             `json:"manufacturers"`
	ProductGroup               interface{}                             `json:"product_group"`
	RefundPolicy               interface{}                             `json:"refund_policy"`
	ReorderPoint               interface{}                             `json:"reorder_point"`
	StockMaximum               interface{}                             `json:"stock_maximum"`
	StockMinimum               interface{}                             `json:"stock_minimum"`
	TaxesOptions               interface{}                             `json:"taxes_options"`
	DelegatableTo              interface{}                             `json:"delegatable_to"`
	DelegatedFrom              interface{}                             `json:"delegated_from"`
	LoyaltyValues              interface{}                             `json:"loyalty_values"`
	DelegateableTo             interface{}                             `json:"delegateable_to"`
	LinkedProducts             interface{}                             `json:"linked_products"`
	WarrantyNotice             interface{}                             `json:"warranty_notice"`
	CustomProperties           interface{}                             `json:"custom_properties"`
	ServiceQuestions           interface{}                             `json:"service_questions"`
	DefaultTitleColor          string                                  `json:"default_title_color"`
	ExternalReferenceId        string                                  `json:"external_reference_id"`
	ServiceQuestionGroups      interface{}                             `json:"service_question_groups"`
	SerialNumberInputRequired  bool                                    `json:"serial_number_input_required"`
	StockConfigurationLocation interface{}                             `json:"stock_configuration_location"`
}

type UpdateProductReturnResultsImages struct {
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

type UpdateProductReturnResultsPrices struct {
	BranchPrices  []interface{}                                   `json:"branch_prices"`
	DefaultPrices []UpdateProductReturnResultsPricesDefaultPrices `json:"default_prices"`
}

type UpdateProductReturnResultsPricesDefaultPrices struct {
	Cost          int                                                 `json:"cost"`
	Amount        UpdateProductReturnResultsPricesDefaultPricesAmount `json:"amount"`
	Margin        int                                                 `json:"margin"`
	Currency      string                                              `json:"currency"`
	PurchasePrice int                                                 `json:"purchase_price"`
}

type UpdateProductReturnResultsPricesDefaultPricesAmount struct {
	Net   float64 `json:"net"`
	Gross float64 `json:"gross"`
}

type UpdateProductReturnResultsSupplier struct {
	Sku string `json:"sku"`
}

type UpdateProductReturnResultsCreatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type UpdateProductReturnResultsManufacturer struct {
	Iln string `json:"iln"`
}

type UpdateProductReturnResultsConfiguration struct {
	Pricing        CreateProductReturnResultsConfigurationPricing `json:"pricing"`
	AllowZeroPrice bool                                           `json:"allow_zero_price"`
}

type UpdateProductReturnResultsConfigurationPricing struct {
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
	request.Header.Set("Authorization", "Bearer "+token)

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
	request.Header.Set("Authorization", "Bearer "+token)

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

// UpdateProduct is to update an existing product
func UpdateProduct(data CreateProductBody, productId, userId, token string) (UpdateProductReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v1/products/" + userId + "/" + productId

	// Define client
	client := &http.Client{}

	// Convert body data
	convert, err := json.Marshal(data)
	if err != nil {
		return UpdateProductReturn{}, err
	}

	// Define request
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(convert))
	if err != nil {
		return UpdateProductReturn{}, err
	}

	// Set header
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+token)

	// Set response and send request
	response, err := client.Do(request)
	if err != nil {
		return UpdateProductReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json response
	var decode UpdateProductReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return UpdateProductReturn{}, err
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
	request.Header.Set("Authorization", "Bearer "+token)

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
