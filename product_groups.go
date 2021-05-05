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

// ProductGroupsReturn is to decode the json return
type ProductGroupsReturn struct {
	Status  int                          `json:"status"`
	Msg     string                       `json:"msg"`
	Request ProductGroupsReturnRequest   `json:"request"`
	Count   int                          `json:"count"`
	Results []ProductGroupsReturnResults `json:"results"`
}

type ProductGroupsReturnRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

type ProductGroupsReturnResults struct {
	Id             string                              `json:"id"`
	ProductGroupId string                              `json:"product_group_id"`
	Tax            string                              `json:"tax"`
	Account        string                              `json:"account"`
	CreatedAt      ProductGroupsReturnResultsCreatedAt `json:"created_at"`
	UpdatedAt      ProductGroupsReturnResultsUpdatedAt `json:"updated_at"`
	Images         ProductGroupsReturnResultsImages    `json:"images"`
	Name           string                              `json:"name"`
	CustomIds      interface{}                         `json:"custom_ids"`
	Color          string                              `json:"color"`
	Active         bool                                `json:"active"`
	Deleted        bool                                `json:"deleted"`
	Metadata       interface{}                         `json:"metadata"`
}

type ProductGroupsReturnResultsCreatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type ProductGroupsReturnResultsUpdatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type ProductGroupsReturnResultsImages struct {
	First    string `json:"1x"`
	Second   string `json:"2x"`
	Third    string `json:"3x"`
	Avatar   string `json:"avatar"`
	Original string `json:"original"`
}

// CreateProductGroupBody is to send body data
type CreateProductGroupBody struct {
	Active         bool                         `json:"active"`
	Color          string                       `json:"color"`
	Name           string                       `json:"name"`
	ProductGroupId string                       `json:"product_group_id"`
	Tax            string                       `json:"tax"`
	Account        string                       `json:"account"`
	Images         CreateProductGroupBodyImages `json:"images"`
}

type CreateProductGroupBodyImages struct {
	First    string `json:"1x"`
	Second   string `json:"2x"`
	Third    string `json:"3x"`
	Avatar   string `json:"avatar"`
	Original string `json:"original"`
}

// CreateProductGroupReturn is to decode json data
type CreateProductGroupReturn struct {
	Status  int                               `json:"status"`
	Msg     string                            `json:"msg"`
	Request CreateProductGroupReturnRequest   `json:"request"`
	Count   int                               `json:"count"`
	Results []CreateProductGroupReturnResults `json:"results"`
}

type CreateProductGroupReturnRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

type CreateProductGroupReturnResults struct {
	Id             string                                   `json:"id"`
	ProductGroupId string                                   `json:"product_group_id"`
	Tax            string                                   `json:"tax"`
	Account        string                                   `json:"account"`
	CreatedAt      CreateProductGroupReturnResultsCreatedAt `json:"created_at"`
	UpdatedAt      CreateProductGroupReturnResultsUpdatedAt `json:"updated_at"`
	Images         CreateProductGroupReturnResultsImages    `json:"images"`
	Name           string                                   `json:"name"`
	CustomIds      interface{}                              `json:"custom_ids"`
	Color          string                                   `json:"color"`
	Active         bool                                     `json:"active"`
	Deleted        bool                                     `json:"deleted"`
	Metadata       interface{}                              `json:"metadata"`
}

type CreateProductGroupReturnResultsCreatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type CreateProductGroupReturnResultsUpdatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type CreateProductGroupReturnResultsImages struct {
	First    string `json:"1x"`
	Second   string `json:"2x"`
	Third    string `json:"3x"`
	Avatar   string `json:"avatar"`
	Original string `json:"original"`
}

// DeleteProductGroupReturn is to decode json return
type DeleteProductGroupReturn struct {
	Status  int                             `json:"status"`
	Msg     string                          `json:"msg"`
	Request DeleteProductGroupReturnRequest `json:"request"`
}

type DeleteProductGroupReturnRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

// ProductGroups is to check all product groups
func ProductGroups(accountId, token string) (ProductGroupsReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v0/product_groups/" + accountId

	// Define client
	client := &http.Client{}

	// Define request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ProductGroupsReturn{}, err
	}

	// Set header
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Authorization", "Bearer "+token)

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return ProductGroupsReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json response
	var decode ProductGroupsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductGroupsReturn{}, err
	}

	// Return data
	return decode, nil

}

// CreateProductGroup is to create a new product group
func CreateProductGroup(data CreateProductGroupBody, accountId, token string) (CreateProductGroupReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v0/product_groups/" + accountId

	// Define client
	client := &http.Client{}

	// Prepare body data
	convert, err := json.Marshal(data)
	if err != nil {
		return CreateProductGroupReturn{}, err
	}

	// Define request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(convert))
	if err != nil {
		return CreateProductGroupReturn{}, err
	}

	// Set header
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+token)

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return CreateProductGroupReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json response
	var decode CreateProductGroupReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CreateProductGroupReturn{}, err
	}

	// Return data
	return decode, nil

}

// UpdateProductGroup is to update a product group
func UpdateProductGroup(data CreateProductGroupBody, groupId, accountId, token string) (CreateProductGroupReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v0/product_groups/" + accountId + "/" + groupId

	// Define client
	client := &http.Client{}

	// Prepare body data
	convert, err := json.Marshal(data)
	if err != nil {
		return CreateProductGroupReturn{}, err
	}

	// Define request
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(convert))
	if err != nil {
		return CreateProductGroupReturn{}, err
	}

	// Set header
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+token)

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return CreateProductGroupReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json response
	var decode CreateProductGroupReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CreateProductGroupReturn{}, err
	}

	// Return data
	return decode, nil

}

func DeleteProductGroup(groupId, accountId, token string) (DeleteProductGroupReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v0/product_groups/" + accountId + "/" + groupId

	// Define client
	client := &http.Client{}

	// Define request
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return DeleteProductGroupReturn{}, err
	}

	// Set header
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+token)

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return DeleteProductGroupReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json response
	var decode DeleteProductGroupReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return DeleteProductGroupReturn{}, err
	}

	// Return data
	return decode, nil

}
