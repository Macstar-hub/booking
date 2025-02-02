package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type ResourceAccess struct {
	Account Account `json:"account"`
}

type Account struct {
	Roles []string `json:"roles"`
}

type RealmAccess struct {
	Roles []string `json:"roles"`
}

type userEmail struct {
	Email          string         `json:"email"`
	EXP            int            `json:"exp"`
	AuthTime       int            `json:"auth_time"`
	AllowedOrigins []string       `json:"allowed-origins"`
	RealmAccess    RealmAccess    `json:"realm_access"`
	ResourceAccess ResourceAccess `json:"resource_access"`
}

func jsonLoad(jsonFileName string, jsonToken string) {
	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		fmt.Println("Cannot open json file with error: ", err)
	}

	defer jsonFile.Close()

	byteJsonFile, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Cannot make byte from json file with error: ", err)
	}
	fmt.Println(byteJsonFile)

	var users userEmail
	err = json.Unmarshal([]byte(jsonToken), &users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users.RealmAccess.Roles)
	fmt.Println(users.ResourceAccess.Account.Roles)
}

func decodeAccessToken(token string) []byte {
	part := strings.Split(token, ".")
	payload, err := base64.RawURLEncoding.DecodeString(part[1])
	if err != nil {
		log.Println("Cannot decode Access token with error: ", err)
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		fmt.Println("Failed to parse payload as JSON:", err)
	}

	// Print the decoded claims
	fmt.Println("Decoded Access Token Claims:")
	jsonClaims, _ := json.MarshalIndent(claims, "", "    ")
	fmt.Println(string(jsonClaims))
	jsonLoad("sample.json", string(jsonClaims))
	return jsonClaims
}

func main() {
	configURL := "http://192.168.1.100:8090/realms/admin-area"
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, configURL)
	if err != nil {
		log.Println("Cannot connection to keycloak: ", err)
	}

	clientID := "admin-booking"
	clientSecret := "LOmb3mUS4YiRWlT6s9be7Uk3NHdSGBYU"

	redirectURL := "http://192.168.1.100:80/api/v1/getusers"

	ouath2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	state := "someState"

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}

	verifier := provider.Verifier(oidcConfig)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		uriForwarded := r.URL.Query().Get("uri")
		redirectURL := "http://192.168.1.100" + uriForwarded
		fmt.Println(redirectURL)

		ouath2Config := oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
		}

		fmt.Println("Debug URI forwarded: ", uriForwarded)

		// Get all headers
		// for k, v := range r.Header {
		// 	fmt.Printf("%v: %v\n", k, v)
		// }
		//

		rawAccessToken := r.Header.Get("Authorization")
		if rawAccessToken == "" {
			http.Redirect(w, r, ouath2Config.AuthCodeURL(state), http.StatusFound)
			return
		}
		parts := strings.Split(rawAccessToken, " ")
		if len(parts) != 2 {
			w.WriteHeader(400)
			return
		}
		_, err := verifier.Verify(ctx, parts[1])
		if err != nil {
			http.Redirect(w, r, ouath2Config.AuthCodeURL(state), http.StatusFound)
			return
		}

		w.Write([]byte("Hello from pervios token get."))
	})

	http.HandleFunc("/booking/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "State did not match", http.StatusBadRequest)
		}
		oauth2Token, err := ouath2Config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			fmt.Println("Cannot exchange oauth2.", err)
		}
		rawIDToken, _ := oauth2Token.Extra("id_token").(string)
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			fmt.Println("Cannot verify token: ", err)
		}

		// decodeAccessToken(oauth2Token.AccessToken)

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage
		}{oauth2Token, new(json.RawMessage)}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(data)
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:8091", nil))
}
