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

package constant

const (
	// unspecified visibility
	VisibilityUnspecified = "VISIBILITY_UNSPECIFIED"
	// private visibility
	VisibilityPrivate = "PRIVATE"
	// public visibility
	VisibilityPublic = "PUBLIC"
	// secret visibility
	VisibilitySecret = "SECRET"
)

const (
	// high priority
	DeliveryPriorityHigh = "HIGH"
	// normal priority
	DeliveryPriorityNormal = "NORMAL"
)

const (
	// high priority
	NotificationPriorityHigh = "HIGH"
	// default priority
	NotificationPriorityDefault = "NORMAL"
	// low priority
	NotificationPriorityLow = "LOW"
)

const (
	// very low urgency
	UrgencyVeryLow = "very-low"
	// low urgency
	UrgencyLow = "low"
	// normal urgency
	UrgencyNormal = "normal"
	// high urgency
	UrgencyHigh = "high"
)

const (
	// webPush text direction auto
	DirAuto = "auto"
	// webPush text direction ltr
	DirLtr = "ltr"
	// webPush text direction rtl
	DirRtl = "rtl"
)

const (
	// success code from push server
	Success = "80000000"
	// parameter invalid code from push server
	ParameterError = "80100001"
	// token invalid code from push server
	TokenFailedErr = "80200001"
	//token timeout code from push server
	TokenTimeoutErr = "80200003"
)

const (
	StyleBigText = iota + 1
)

const (
	TypeIntentOrAction = iota + 1
	TypeUrl
	TypeApp
	TypeRichResource
)

const (
	FastAppTargetDevelop = iota + 1
	FastAppTargetProduct
)

const (
	// test user
	TargetUserTypeTest = iota + 1
	// formal user
	TargetUserTypeFormal
	// VoIP user
	TargetUserTypeVoIP
)
