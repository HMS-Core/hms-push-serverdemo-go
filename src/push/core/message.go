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
	"fmt"
	"github.com/neoscott/hms-push-serverdemo-go/src/httpclient"
	"github.com/neoscott/hms-push-serverdemo-go/src/push/constant"
	"github.com/neoscott/hms-push-serverdemo-go/src/push/model"
	"github.com/neoscott/hms-push-serverdemo-go/src/push/verify"
	"net/http"


)

// SendMessage sends a message to huawei cloud common
// One of Token, Topic and Condition fields must be invoked in message
// If validationOnly is set to true, the message can be verified by not sent to users
func (c *HttpPushClient) SendMessage(ctx context.Context, msgRequest *model.MessageRequest) (*model.MessageResponse, error) {
	result := &model.MessageResponse{}

	err := verify.ValidateMessage(msgRequest.Message)
	if err != nil {
		return nil, err
	}

	request, err := c.getSendMsgRequest(msgRequest)
	if err != nil {
		return nil, err
	}

	err = c.executeApiOperation(ctx, request, result)
	if err != nil {
		return result, err
	}
	return result, err
}

func (c *HttpPushClient) getSendMsgRequest(msgRequest *model.MessageRequest) (*httpclient.PushRequest, error) {
	body, err := json.Marshal(msgRequest)
	if err != nil {
		return nil, err
	}

	request := &httpclient.PushRequest{
		Method: http.MethodPost,
		URL:    fmt.Sprintf(constant.SendMessageFmt, c.endpoint, c.appId),
		Body:   body,
		Header: []httpclient.HTTPOption{
			httpclient.SetHeader("Content-Type", "application/json;charset=utf-8"),
			httpclient.SetHeader("Authorization", "Bearer "+c.token),
		},
	}
	return request, nil
}
