// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "auth": login Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-auth/design
// --notool=true
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-jenkins-idler/internal/auth
// --pkg=client
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// LoginLoginPath computes a request path to the login action of login.
func LoginLoginPath() string {

	return fmt.Sprintf("/api/login")
}

// Login user
func (c *Client) LoginLogin(ctx context.Context, path string, apiClient *string, redirect *string, scope *string) (*http.Response, error) {
	req, err := c.NewLoginLoginRequest(ctx, path, apiClient, redirect, scope)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginLoginRequest create the request corresponding to the login action endpoint of the login resource.
func (c *Client) NewLoginLoginRequest(ctx context.Context, path string, apiClient *string, redirect *string, scope *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if apiClient != nil {
		values.Set("api_client", *apiClient)
	}
	if redirect != nil {
		values.Set("redirect", *redirect)
	}
	if scope != nil {
		values.Set("scope", *scope)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
