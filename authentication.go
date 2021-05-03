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

// AuthBody is to send the login data to tillhub
type AuthBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthReturn is to decode json return from api
type AuthReturn struct {
	Status        int          `json:"status"`
	Msg           string       `json:"msg"`
	Request       AuthRequest  `json:"request"`
	User          AuthUser     `json:"user"`
	ValidPassword bool         `json:"valid_password"`
	Token         string       `json:"token"`
	TokenType     string       `json:"token_type"`
	ExpiresAt     string       `json:"expires_at"`
	Features      AuthFeatures `json:"features"`
}

type AuthRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

type AuthUser struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	LegacyId    interface{} `json:"legacy_id"`
	DisplayName string      `json:"display_name"`
	Whitelabel  string      `json:"whitelabel"`
	Scopes      []string    `json:"scopes"`
	Role        string
}

type AuthFeatures struct {
	Crm              bool `json:"crm"`
	Datev            bool `json:"datev"`
	Pagers           bool `json:"pagers"`
	Exports          bool `json:"exports"`
	Invoices         bool `json:"invoices"`
	Vouchers         bool `json:"vouchers"`
	Savedcart        bool `json:"savedcart"`
	Timetracking     bool `json:"timetracking"`
	Fiscalisation    bool `json:"fiscalisation"`
	GastroTables     bool `json:"gastro_tables"`
	DeliveryNotes    bool `json:"delivery_notes"`
	LabelPrinters    bool `json:"label_printers"`
	SafeManagement   bool `json:"safe_management"`
	KitchenPrinters  bool `json:"kitchen_printers"`
	PromotionEngine  bool `json:"promotion_engine"`
	BranchManagement bool `json:"branch_management"`
}

// Auth is to authenticate and get an bearer token back
func Auth(email, password string) (AuthReturn, error) {

	// Url
	url := "https://api.tillhub.com/api/v0/users/login"

	// Define client
	client := &http.Client{}

	// Define json body
	body := AuthBody{email, password}

	// Prepare json body for request
	convert, err := json.Marshal(body)
	if err != nil {
		return AuthReturn{}, err
	}

	// Define request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(convert))
	if err != nil {
		return AuthReturn{}, err
	}

	// Set header
	request.Header.Set("Content-Type", "application/json")

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return AuthReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json return
	var decode AuthReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return AuthReturn{}, err
	}

	// Return data
	return decode, nil

}
