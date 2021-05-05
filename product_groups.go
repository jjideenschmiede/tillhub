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
	First   string `json:"1x"`
	Second  string `json:"2x"`
	Third   string `json:"3x"`
	Avatar  string `json:"avatar"`
	Orginal string `json:"orginal"`
}

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
