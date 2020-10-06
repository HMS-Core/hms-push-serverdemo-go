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

package common

import (
	"fmt"
	"github.com/neoscott/hms-push-serverdemo-go/src/push/config"
	"github.com/neoscott/hms-push-serverdemo-go/src/push/core"

	"sync"
)

const (
	// private data ,it's import,please don't let it out
	appId     = "xxxxxx"
	appSecret = "xxxxxx"
	token     = "xxxxxx"

	// below is public address
	// get token address
	authUrl   = "https://login.cloud.huawei.com/oauth2/v2/token"
	// send push msg address
	pushUrl   = "https://api.push.hicloud.com"
)

var conf = &config.Config{
	AppId:     appId,
	AppSecret: appSecret,
	AuthUrl:   authUrl,
	PushUrl:   pushUrl,
}

var (
	pushClient *core.HttpPushClient
	once       sync.Once
)

var (
	//TargetToken the topic to be subscribed/unsubscribed
	TargetTopic = "topic"

	//TargetCondition the condition of the devices operated
	TargetCondition = "'topic' in topics && ('topic' in topics || 'TopicC' in topics)"

	//TargetToken the token of the device operated
	TargetToken = token

	//TargetTokenArray the collection of the tokens of th devices operated
	TargetTokenArray = []string{TargetToken}
)

func GetPushClient() *core.HttpPushClient {
	once.Do(func() {
		client, err := core.NewHttpClient(conf)
		if err != nil {
			fmt.Printf("Failed to new common client! Error is %s\n", err.Error())
			panic(err)
		}
		pushClient = client
	})

	return pushClient
}

func GetPushConf() *config.Config{
	return conf
}
