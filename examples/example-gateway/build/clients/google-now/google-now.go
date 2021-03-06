// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package googlenowClient

import (
	"context"
	"strconv"

	clientsGooglenowGooglenow "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/googlenow/googlenow"
	"github.com/uber/zanzibar/runtime"
)

// Client defines google-now client interface.
type Client interface {
	HTTPClient() *zanzibar.HTTPClient
	AddCredentials(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsGooglenowGooglenow.GoogleNowService_AddCredentials_Args,
	) (map[string]string, error)
	CheckCredentials(
		ctx context.Context,
		reqHeaders map[string]string,
	) (map[string]string, error)
}

// googleNowClient is the http client.
type googleNowClient struct {
	clientID   string
	httpClient *zanzibar.HTTPClient
}

// NewClient returns a new http client.
func NewClient(gateway *zanzibar.Gateway) Client {
	ip := gateway.Config.MustGetString("clients.google-now.ip")
	port := gateway.Config.MustGetInt("clients.google-now.port")

	baseURL := "http://" + ip + ":" + strconv.Itoa(int(port))
	return &googleNowClient{
		clientID:   "google-now",
		httpClient: zanzibar.NewHTTPClient(gateway, baseURL),
	}
}

// HTTPClient returns the underlying HTTP client, should only be
// used for internal testing.
func (c *googleNowClient) HTTPClient() *zanzibar.HTTPClient {
	return c.httpClient
}

// AddCredentials calls "/add-credentials" endpoint.
func (c *googleNowClient) AddCredentials(
	ctx context.Context,
	headers map[string]string,
	r *clientsGooglenowGooglenow.GoogleNowService_AddCredentials_Args,
) (map[string]string, error) {
	req := zanzibar.NewClientHTTPRequest(
		c.clientID, "addCredentials", c.httpClient,
	)
	// TODO(jakev): Ensure we validate mandatory headers

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/add-credentials"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}
	// TODO(jakev): verify mandatory response headers

	res.CheckOKResponse([]int{202})

	switch res.StatusCode {
	case 202:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}
		return respHeaders, nil
	}

	return respHeaders, &zanzibar.UnexpectedHTTPError{
		StatusCode: res.StatusCode,
		RawBody:    res.GetRawBody(),
	}
}

// CheckCredentials calls "/check-credentials" endpoint.
func (c *googleNowClient) CheckCredentials(
	ctx context.Context,
	headers map[string]string,
) (map[string]string, error) {
	req := zanzibar.NewClientHTTPRequest(
		c.clientID, "checkCredentials", c.httpClient,
	)
	// TODO(jakev): Ensure we validate mandatory headers

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/check-credentials"

	err := req.WriteJSON("POST", fullURL, headers, nil)
	if err != nil {
		return nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}
	// TODO(jakev): verify mandatory response headers

	res.CheckOKResponse([]int{202})

	switch res.StatusCode {
	case 202:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}
		return respHeaders, nil
	}

	return respHeaders, &zanzibar.UnexpectedHTTPError{
		StatusCode: res.StatusCode,
		RawBody:    res.GetRawBody(),
	}
}
