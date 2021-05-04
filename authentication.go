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

// AuthBasicBody is to send the login data to tillhub
type AuthBasicBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthBasicReturn is to decode json return from api
type AuthBasicReturn struct {
	Status        int               `json:"status"`
	Msg           string            `json:"msg"`
	Request       AuthBasicRequest  `json:"request"`
	User          AuthBasicUser     `json:"user"`
	ValidPassword bool              `json:"valid_password"`
	Token         string            `json:"token"`
	TokenType     string            `json:"token_type"`
	ExpiresAt     string            `json:"expires_at"`
	Features      AuthBasicFeatures `json:"features"`
}

type AuthBasicRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

type AuthBasicUser struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	LegacyId    interface{} `json:"legacy_id"`
	DisplayName string      `json:"display_name"`
	Whitelabel  string      `json:"whitelabel"`
	Scopes      []string    `json:"scopes"`
	Role        string      `json:"role"`
}

type AuthBasicFeatures struct {
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

// AuthKeyBody is to send the login data to tillhub
type AuthKeyBody struct {
	Id     string `json:"id"`
	ApiKey string `json:"api_key"`
}

// AuthKeyReturn is to decode json response
type AuthKeyReturn struct {
	Status    int                  `json:"status"`
	Msg       string               `json:"msg"`
	Request   AuthKeyReturnRequest `json:"request"`
	User      AuthKeyReturnUser    `json:"user"`
	Token     string               `json:"token"`
	TokenType string               `json:"token_type"`
	SubUser   AuthKeyReturnSubUser `json:"sub_user"`
	ExpiresAt string               `json:"expires_at"`
}

type AuthKeyReturnRequest struct {
	Host string `json:"host"`
	Id   string `json:"id"`
}

type AuthKeyReturnUser struct {
	Id       string      `json:"id"`
	Name     string      `json:"name"`
	LegacyId interface{} `json:"legacy_id"`
}

type AuthKeyReturnSubUser struct {
	Id                       string                   `json:"id"`
	CreatedAt                string                   `json:"created_at"`
	UpdatedAt                string                   `json:"updated_at"`
	Metadata                 interface{}              `json:"metadata"`
	Groups                   interface{}              `json:"groups"`
	Scopes                   interface{}              `json:"scopes"`
	Attributes               interface{}              `json:"attributes"`
	User                     AuthKeyReturnSubUserUser `json:"user"`
	Description              interface{}              `json:"description"`
	Active                   bool                     `json:"active"`
	Role                     string                   `json:"role"`
	Parents                  interface{}              `json:"parents"`
	Children                 interface{}              `json:"children"`
	Blocked                  bool                     `json:"blocked"`
	ConfigurationId          string                   `json:"configuration_id"`
	UserId                   interface{}              `json:"user_id"`
	Deleted                  bool                     `json:"deleted"`
	ApiKey                   string                   `json:"api_key"`
	Key                      interface{}              `json:"key"`
	Username                 interface{}              `json:"username"`
	Name                     string                   `json:"name"`
	Locations                interface{}              `json:"locations"`
	Firstname                interface{}              `json:"firstname"`
	Lastname                 interface{}              `json:"lastname"`
	UserPermissionTemplateId interface{}              `json:"user_permission_template_id"`
}

type AuthKeyReturnSubUserUser struct {
	Id    interface{} `json:"id"`
	Email interface{} `json:"email"`
}

// AuthBasic is to authenticate and get an bearer token back
func AuthBasic(email, password string) (AuthBasicReturn, error) {

	// Url
	url := "https://api.tillhub.com/api/v0/users/login"

	// Define client
	client := &http.Client{}

	// Define json body
	body := AuthBasicBody{email, password}

	// Prepare json body for request
	convert, err := json.Marshal(body)
	if err != nil {
		return AuthBasicReturn{}, err
	}

	// Define request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(convert))
	if err != nil {
		return AuthBasicReturn{}, err
	}

	// Set header
	request.Header.Set("Content-Type", "application/json")

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return AuthBasicReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json return
	var decode AuthBasicReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return AuthBasicReturn{}, err
	}

	// Return data
	return decode, nil

}

func AuthKey(accountId, apiKey string) (AuthKeyReturn, error) {

	// Url
	url := "https://api.tillhub.com/api/v1/users/auth/key"

	// Define client
	client := &http.Client{}

	// Define json body
	body := AuthKeyBody{accountId, apiKey}

	// Prepare json for request
	convert, err := json.Marshal(body)
	if err != nil {
		return AuthKeyReturn{}, err
	}

	// Define request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(convert))
	if err != nil {
		return AuthKeyReturn{}, err
	}

	// Set header
	request.Header.Set("Content-Type", "application/json")

	// Define response & send request
	response, err := client.Do(request)
	if err != nil {
		return AuthKeyReturn{}, err
	}

	// Close body after function ends
	defer response.Body.Close()

	// Decode json return
	var decode AuthKeyReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return AuthKeyReturn{}, err
	}

	// Return data
	return decode, nil

}
