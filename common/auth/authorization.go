package auth

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

func VerifyIDToken(ctx context.Context, idpName string, token *oauth2.Token) (*oidc.IDToken, error) {
	idp := getIDP(idpName)
	if idp == nil {
		return nil, fmt.Errorf("IDP name %s is not registered", idpName)
	}

	return idp.idTokenVerifier.Verify(ctx, token.Extra("id_token").(string))
}

// UserInfo get user info
func UserInfo(idp string, token *oauth2.Token) (map[string]interface{}, error) {
	return nil, nil
}

// RevokeToken revoke token
func RevokeToken(token, idp string) error {
	return nil
}

// RefreshToken refresh token
func RefreshToken(token, idp string) (string, error) {
	return "", nil
}

// AuthCodeURL returning idp auth url
func AuthCodeURL(idpName, state string) (string, error) {
	idp := getIDP(idpName)
	if idp == nil {
		return "", fmt.Errorf("IDP name %s is not registered", idpName)
	}
	return idp.oauth2Conf.AuthCodeURL(state), nil
}

// ExchangeToken exchange code to token
func ExchangeToken(ctx context.Context, idpName, code string) (*oauth2.Token, error) {
	idp := getIDP(idpName)
	if idp == nil {
		return nil, fmt.Errorf("IDP name %s is not registered", idpName)
	}

	return idp.oauth2Conf.Exchange(ctx, code)
}
