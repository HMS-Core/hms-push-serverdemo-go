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

package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/neoscott/hms-push-serverdemo-go/src/httpclient"
	auth "github.com/neoscott/hms-push-serverdemo-go/src/push/authention"
	"github.com/neoscott/hms-push-serverdemo-go/src/push/config"
	"github.com/neoscott/hms-push-serverdemo-go/src/push/constant"

	"reflect"


)

type HttpPushClient struct {
	endpoint   string
	appId      string
	token      string
	authClient *auth.AuthClient
	client     *httpclient.HTTPClient
}

// NewClient creates a instance of the huawei cloud common client
// It's contained in huawei cloud app and provides service through huawei cloud app
func NewHttpClient(c *config.Config) (*HttpPushClient, error) {
	if c.AppId == "" {
		return nil, errors.New("appId can't be empty")
	}

	if c.PushUrl == "" {
		return nil, errors.New("pushUrl can't be empty")
	}

	client, err := httpclient.NewHTTPClient()
	if err != nil {
		return nil, errors.New("failed to get http client")
	}

	authClient, err := auth.NewAuthClient(c)
	if err != nil {
		return nil, err
	}

	token, err := authClient.GetAuthToken(context.Background())
	if err != nil {
		return nil, errors.New("refresh token fail")
	}

	return &HttpPushClient{
		endpoint:   c.PushUrl,
		appId:      c.AppId,
		token:      token,
		authClient: authClient,
		client:     client,
	}, nil
}

func (c *HttpPushClient) refreshToken() error {
	if c.authClient == nil {
		return errors.New("can't refresh token because getting auth client fail")
	}

	token, err := c.authClient.GetAuthToken(context.Background())
	if err != nil {
		return errors.New("refresh token fail")
	}

	c.token = token
	return nil
}

func (c *HttpPushClient) executeApiOperation(ctx context.Context, request *httpclient.PushRequest, responsePointer interface{}) error {
	err := c.sendHttpRequest(ctx, request, responsePointer)
	if err != nil {
		return err
	}

	// if need to retry for token timeout or other reasons
	retry, err := c.isNeedRetry(responsePointer)
	if err != nil {
		return err
	}

	if retry {
		err = c.sendHttpRequest(ctx, request, responsePointer)
		return err
	}
	return err
}

func (c *HttpPushClient) sendHttpRequest(ctx context.Context, request *httpclient.PushRequest, responsePointer interface{}) error {
	resp, err := c.client.DoHttpRequest(ctx, request)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(resp.Body, responsePointer); err != nil {
		return err
	}
	return nil
}

// if token is timeout or error or other reason, need to refresh token and send again
func (c *HttpPushClient) isNeedRetry(responsePointer interface{}) (bool, error) {
	tokenError, err := isTokenError(responsePointer)
	if err != nil {
		return false, err
	}

	if !tokenError {
		return false, nil
	}

	err = c.refreshToken()
	if err != nil {
		return false, err
	}
	return true, nil
}

// if token is timeout or error, refresh token and send again
func isTokenError(responsePointer interface{}) (bool, error) {
	//the responsePointer must be point of struct
	val, _, ok := checkParamStructPtr(responsePointer)
	if !ok {
		return false, errors.New("the parameter should be pointer of the struct")
	}

	code := val.Elem().FieldByName("Code").String()
	if code == constant.TokenTimeoutErr || code == constant.TokenFailedErr {
		return true, nil
	}
	return false, nil
}

func checkParamStructPtr(structPtr interface{}) (reflect.Value, reflect.Type, bool) {
	val := reflect.ValueOf(structPtr)
	if val.Kind() != reflect.Ptr {
		fmt.Println("The Parameter should be Pointer of Struct!")
		return reflect.Value{}, nil, false
	}

	t := reflect.Indirect(val).Type()
	if t.Kind() != reflect.Struct {
		fmt.Println("The Parameter should be Pointer of Struct!")
		return reflect.Value{}, nil, false
	}
	return val, t, true
}
