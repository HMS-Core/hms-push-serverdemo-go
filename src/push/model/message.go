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

package model

import (
	"github.com/neoscott/hms-push-serverdemo-go/src/push/constant"
)

type MessageRequest struct {
	ValidateOnly bool     `json:"validate_only"`
	Message      *Message `json:"message"`
}

type MessageResponse struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}

type Message struct {
	Data         string         `json:"data,omitempty"`
	Notification *Notification  `json:"notification,omitempty"`
	Android      *AndroidConfig `json:"android,omitempty"`
	Apns         *Apns          `json:"apns,omitempty"`
	WebPush      *WebPushConfig `json:"webpush,omitempty"`
	Token        []string       `json:"token,omitempty"`
	Topic        string         `json:"topic,omitempty"`
	Condition    string         `json:"condition,omitempty"`
}

type Notification struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	Image string `json:"image,omitempty"`
}

type Apns struct {
	Headers    *ApnsHeaders           `json:"headers,omitempty"`
	Payload    map[string]interface{} `json:"payload,omitempty"`
	HmsOptions *ApnsHmsOptions        `json:"hms_options,omitempty"`
}

type ApnsHmsOptions struct {
	TargetUserType int `json:"target_user_type,omitempty"`
}

type ApnsHeaders struct {
	Authorization  string `json:"authorization,omitempty"`
	ApnsId         string `json:"apns-id,omitempty"`
	ApnsExpiration int64  `json:"apns-expiration,omitempty"`
	ApnsPriority   string `json:"apns-priority,omitempty"`
	ApnsTopic      string `json:"apns-topic,omitempty"`
	ApnsCollapseId string `json:"apns-collapse-id,omitempty"`
}

type Aps struct {
	Alert            interface{} `json:"alert,omitempty"` // dictionary or string
	Badge            int         `json:"badge,omitempty"`
	Sound            string      `json:"sound,omitempty"`
	ContentAvailable int         `json:"content-available,omitempty"`
	Category         string      `json:"category,omitempty"`
	ThreadId         string      `json:"thread-id,omitempty"`
}

type AlertDictionary struct {
	Title        string   `json:"title,omitempty"`
	Body         string   `json:"body,omitempty"`
	TitleLocKey  string   `json:"title-loc-key,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	ActionLocKey string   `json:"action-loc-key,omitempty"`
	LocKey       string   `json:"loc-key,omitempty"`
	LocArgs      []string `json:"loc-args,omitempty"`
	LaunchImage  string   `json:"launch-image,omitempty"`
}

//NewTransparentMsgRequest will return a new MessageRequest instance with default value to send transparent message.
//developers should set at least on of Message.Token or  Message.Topic or Message.Condition
func NewTransparentMsgRequest() *MessageRequest {
	msgRequest := getDefaultMsgRequest()
	return msgRequest
}

// NewNotificationMsgRequest will return a new MessageRequest instance with default value to send notification message.
//developers should set at least on of Message.Token or  Message.Topic or Message.Condition
func NewNotificationMsgRequest() *MessageRequest {
	msgRequest := getDefaultMsgRequest()
	msgRequest.Message.Notification = getDefaultNotification()
	return msgRequest
}

func getDefaultMsgRequest() *MessageRequest {
	return &MessageRequest{
		ValidateOnly: false,
		Message: &Message{
			Data: "This is a transparent message data",
		},
	}
}

func GetDefaultApns() *Apns {
	return &Apns{
		HmsOptions: &ApnsHmsOptions{
			TargetUserType: constant.TargetUserTypeTest,
		},
		Payload: getDefaultApnsPayload(),
	}
}

func getDefaultApnsPayload() map[string]interface{} {
	payLoad := make(map[string]interface{}, 0)
	aps := &Aps{
		Alert: &AlertDictionary{
			Title: "title",
			Body:  "apns body",
			TitleLocKey : "PLAY",
		},
		Badge: 5,
	}

	payLoad["aps"] = aps
	return payLoad
}

func getDefaultNotification() *Notification {
	return &Notification{
		Title: "notification tittle",
		Body:  "This is a notification message body",
	}
}
