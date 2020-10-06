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

package verify

import (
	"errors"
	"github.com/neoscott/hms-push-serverdemo-go/src/push/model"
	"regexp"


)

var (
	ttlPattern   = regexp.MustCompile("\\d+|\\d+[sS]|\\d+.\\d{1,9}|\\d+.\\d{1,9}[sS]")
	colorPattern = regexp.MustCompile("^#[0-9a-fA-F]{6}$")
)

func ValidateMessage(message *model.Message) error {
	if message == nil {
		return errors.New("message must not be null")
	}

	// validate field target, one of Token, Topic and Condition must be invoked
	if err := validateFieldTarget(message.Token, message.Topic, message.Condition); err != nil {
		return err
	}

	// validate android config
	if err := validateAndroidConfig(message.Android); err != nil {
		return err
	}

	// validate web common config
	if err := validateWebPushConfig(message.WebPush); err != nil {
		return err
	}
	return nil
}

func validateFieldTarget(token []string, strings ...string) error {
	count := 0
	if token != nil {
		count++
	}

	for _, s := range strings {
		if s != "" {
			count++
		}
	}

	if count == 1 {
		return nil
	}
	return errors.New("token, topics or condition must be choice one")
}
