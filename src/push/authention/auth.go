/*
Copyright 2020. Huawei Technologies Co., Ltd. All rights reserved.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/neoscott/hms-push-serverdemo-go/src/httpclient"
	"github.com/neoscott/hms-push-serverdemo-go/src/push/config"
	"net/http"
	"net/url"

)

type AuthClient struct {
	endpoint  string
	appId     string
	appSecret string
	client    *httpclient.HTTPClient
}

type TokenMsg struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// NewClient creates a instance of the huawei cloud auth client
// It's contained in huawei cloud app and provides service through huawei cloud app
// If AuthUrl is null using default auth url address
func NewAuthClient(conf *config.Config) (*AuthClient, error) {
	if conf.AppId == "" || conf.AppSecret == "" {
		return nil, errors.New("appId or appSecret is null")
	}

	c, err := httpclient.NewHTTPClient()
	if err != nil {
		return nil, errors.New("failed to get http client")
	}

	if conf.AuthUrl == "" {
		return nil, errors.New("authUrl can't be empty")
	}

	return &AuthClient{
		endpoint:  conf.AuthUrl,
		appId:     conf.AppId,
		appSecret: conf.AppSecret,
		client:    c,
	}, nil
}

// GetAuthToken gets token from huawei cloud
// the developer can access the app by using this token
func (ac *AuthClient) GetAuthToken(ctx context.Context) (string, error) {
	u, _ := url.Parse(ac.appSecret)
	body := fmt.Sprintf("grant_type=client_credentials&client_secret=%s&client_id=%s", u.String(), ac.appId)

	request := &httpclient.PushRequest{
		Method: http.MethodPost,
		URL:    ac.endpoint,
		Body:   []byte(body),
		Header: []httpclient.HTTPOption{httpclient.SetHeader("Content-Type", "application/x-www-form-urlencoded")},
	}

	resp, err := ac.client.DoHttpRequest(ctx, request)
	if err != nil {
		return "", err
	}

	var token TokenMsg
	if resp.Status == 200 {
		err = json.Unmarshal(resp.Body, &token)
		if err != nil {
			return "", err
		}
		return token.AccessToken, nil
	}
	return "", nil
}
