package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

func main() {
	configURL := "http://192.168.1.100:8090/realms/admin-area"
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, configURL)
	if err != nil {
		log.Println("Cannot connection to keycloak: ", err)
	}

	clientID := "admin-booking"
	clientSecret := "LOmb3mUS4YiRWlT6s9be7Uk3NHdSGBYU"

	// redirectURL := "http://localhost:8091/booking/callback"
	redirectURL := "http://192.168.100:80/"

	ouath2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	state := "someState"

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}

	verifier := provider.Verifier(oidcConfig)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
