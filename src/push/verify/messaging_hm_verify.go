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

func validateAndroidConfig(androidConfig *model.AndroidConfig) error {
	if androidConfig == nil {
		return nil
	}

	if androidConfig.CollapseKey < -1 || androidConfig.CollapseKey > 100 {
		return errors.New("collapse_key must be in interval [-1 - 100]")
	}

	if androidConfig.Urgency != "" &&
		androidConfig.Urgency != constant.DeliveryPriorityHigh &&
		androidConfig.Urgency != constant.DeliveryPriorityNormal {
		return errors.New("delivery_priority must be 'HIGH' or 'NORMAL'")
	}

	if androidConfig.TTL != "" && !ttlPattern.MatchString(androidConfig.TTL) {
		return errors.New("malformed ttl")
	}

	if androidConfig.FastAppTarget != 0 &&
		androidConfig.FastAppTarget != constant.FastAppTargetDevelop &&
		androidConfig.FastAppTarget != constant.FastAppTargetProduct {
		return errors.New("invalid fast_app_target")
	}

	// validate android notification
	return validateAndroidNotification(androidConfig.Notification)
}

func validateAndroidNotification(notification *model.AndroidNotification) error {
	if notification == nil {
		return nil
	}

	if notification.Sound == "" && notification.DefaultSound == false {
		return errors.New("sound must not be empty when default_sound is false")
	}

	if err := validateAndroidNotifyStyle(notification); err != nil {
		return err
	}

	if err := validateAndroidNotifyPriority(notification); err != nil {
		return err
	}

	if err := validateVibrateTimings(notification); err != nil {
		return err
	}

	if err := validateVisibility(notification); err != nil {
		return err
	}

	if err := validateLightSetting(notification); err != nil {
		return err
	}

	if notification.Color != "" && !colorPattern.MatchString(notification.Color) {
		return errors.New("color must be in the form #RRGGBB")
	}

	// validate click action
	return validateClickAction(notification.ClickAction)
}

func validateAndroidNotifyStyle(notification *model.AndroidNotification) error {
	switch notification.Style {
	case constant.StyleBigText:
		if notification.BigTitle == "" {
			return errors.New("big_title must not be empty when style is 1")
		}

		if notification.BigBody == "" {
			return errors.New("big_body must not be empty when style is 1")
		}

	}
	return nil
}

func validateAndroidNotifyPriority(notification *model.AndroidNotification) error {
	if notification.Importance != "" &&
		notification.Importance != constant.NotificationPriorityHigh &&
		notification.Importance != constant.NotificationPriorityDefault &&
		notification.Importance != constant.NotificationPriorityLow  {
		return errors.New("Importance must be 'HIGH', 'NORMAL' or 'LOW'")
	}
	return nil
}

func validateVibrateTimings(notification *model.AndroidNotification) error {
	if notification.VibrateConfig != nil {
		if len(notification.VibrateConfig) > 10 {
			return errors.New("vibrate_timings can't be more than 10 elements")
		}
		for _, vibrateTiming := range notification.VibrateConfig {
			if !ttlPattern.MatchString(vibrateTiming) {
				return errors.New("malformed vibrate_timings")
			}
		}
	}
	return nil
}

func validateVisibility(notification *model.AndroidNotification) error {
	if notification.Visibility == "" {
		notification.Visibility = constant.VisibilityPrivate
		return nil
	}
	if notification.Visibility != constant.VisibilityUnspecified && notification.Visibility != constant.VisibilityPrivate &&
		notification.Visibility != constant.VisibilityPublic && notification.Visibility != constant.VisibilitySecret {
		return errors.New("visibility must be VISIBILITY_UNSPECIFIED, PRIVATE, PUBLIC or SECRET")
	}
	return nil
}

func validateLightSetting(notification *model.AndroidNotification) error {
	if notification.LightSettings == nil {
		return nil
	}

	if notification.LightSettings.Color == nil {
		return errors.New("light_settings.color can't be nil")
	}

	if notification.LightSettings.LightOnDuration == "" ||
		!ttlPattern.MatchString(notification.LightSettings.LightOnDuration) {
		return errors.New("light_settings.light_on_duration is empty or malformed")
	}

	if notification.LightSettings.LightOffDuration == "" ||
		!ttlPattern.MatchString(notification.LightSettings.LightOffDuration) {
		return errors.New("light_settings.light_off_duration is empty or malformed")
	}
	return nil
}

func validateClickAction(clickAction *model.ClickAction) error {
	if clickAction == nil {
		return errors.New("click_action object must not be null")
	}

	switch clickAction.Type {
	case constant.TypeIntentOrAction:
		if clickAction.Intent == "" && clickAction.Action == "" {
			return errors.New("at least one of intent and action is not empty when type is 1")
		}
	case constant.TypeUrl:
		if clickAction.Url == "" {
			return errors.New("url must not be empty when type is 2")
		}
	case constant.TypeApp:
	case constant.TypeRichResource:
		if clickAction.RichResource == "" {
			return errors.New("rich_resource must not be empty when type is 4")
		}
	default:
		return errors.New("type must be in the interval [1 - 4]")
	}
	return nil
}
