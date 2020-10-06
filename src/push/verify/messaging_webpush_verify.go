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
	"github.com/neoscott/hms-push-serverdemo-go/src/push/constant"
	"github.com/neoscott/hms-push-serverdemo-go/src/push/model"
)

func validateWebPushConfig(webPushConfig *model.WebPushConfig) error {
	if webPushConfig == nil {
		return nil
	}

	if err := validateWebPushHeaders(webPushConfig.Headers); err != nil {
		return err
	}

	return validateWebPushNotification(webPushConfig.Notification)
}

func validateWebPushHeaders(headers *model.WebPushHeaders) error {
	if headers == nil {
		return nil
	}

	if headers.TTL != "" && !ttlPattern.MatchString(headers.TTL) {
		return errors.New("malformed ttl")
	}

	if headers.Urgency != "" &&
		headers.Urgency != constant.UrgencyHigh &&
		headers.Urgency != constant.UrgencyNormal &&
		headers.Urgency != constant.UrgencyLow &&
		headers.Urgency != constant.UrgencyVeryLow {
		return errors.New("priority must be 'high', 'normal', 'low' or 'very-low'")
	}
	return nil
}

func validateWebPushNotification(notification *model.WebPushNotification) error {
	if notification == nil {
		return nil
	}

	if err := validateWebPushAction(notification.Actions); err != nil {
		return err
	}

	if err := validateWebPushDirection(notification.Dir); err != nil {
		return err
	}
	return nil
}

func validateWebPushAction(actions []*model.WebPushAction) error {
	if actions == nil {
		return nil
	}

	for _, action := range actions {
		if action.Action == "" {
			return errors.New("web common action can't be empty")
		}
	}
	return nil
}

func validateWebPushDirection(dir string) error {
	if dir != constant.DirAuto && dir != constant.DirLtr && dir != constant.DirRtl {
		return errors.New("web common dir must be 'auto', 'ltr', 'rtl'")
	}
	return nil
}
