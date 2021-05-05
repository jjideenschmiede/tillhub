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

type TaxReturn struct {
	Status  int                `json:"status"`
	Msg     string             `json:"msg"`
	Request TaxReturnRequest   `json:"request"`
	Count   int                `json:"count"`
	Results []TaxReturnResults `json:"results"`
}

type TaxReturnRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

type TaxReturnResults struct {
	Id                string                    `json:"id"`
	FaAccountNumber   string                    `json:"fa_account_number"`
	CreatedAt         TaxReturnResultsCreatedAt `json:"created_at"`
	UpdatedAt         TaxReturnResultsUpdatedAt `json:"updated_at"`
	Active            bool                      `json:"active"`
	Type              string                    `json:"type"`
	Metadata          interface{}               `json:"metadata"`
	Name              string                    `json:"name"`
	Deleted           bool                      `json:"deleted"`
	Accounts          interface{}               `json:"accounts"`
	FinancialAccounts interface{}               `json:"financial_accounts"`
	Description       interface{}               `json:"description"`
}

type TaxReturnResultsCreatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

type TaxReturnResultsUpdatedAt struct {
	Iso  string `json:"iso"`
	Unix int    `json:"unix"`
}

// Tax is to find the taxes and their ids
func Tax(accountId, token string) (TaxReturn, error) {

	// Set url
	url := "https://api.tillhub.com/api/v0/accounts/" + accountId + "?type=revenue&deleted=false"

	// Define client
	client := &http.Client{}

	// Set request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return TaxReturn{}, err
	}

	// Set header
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Authorization", "Bearer "+token)

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return TaxReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json return
	var decode TaxReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return TaxReturn{}, err
	}

	// Return data
	return decode, nil

}
