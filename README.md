# HMS Core Push Kit Sample Code (Golang)
English | [中文](README_ZH.md)
## Contents
 * [Introduction](#Introduction)
 * [Installation](#Installation)
 * [Configuration](#Configuration)
 * [Sample Code](#Sample-Code)
 * [Technical Support](#technical-support)
 * [License](#License)

## Introduction
The sample code for Golang encapsulates the server-side APIs of Push Kit, for your reference or direct use.

The following table describes packages of Golang sample code.
| Package| Description|
| ----------- | ----------- |
|[examples](src/examples) | Sample code packages.|
|[httpclient](src/httpclient/httpclient.go) | Package for sending network requests.|
|[push](src/push) | Package where Push Kit server APIs are encapsulated.|

## Installation
Install the Golang environment (Golang 1.11 or later is recommended) and decompress the Golang sample code package.
    
Copy **org.huawei.com** in the decompressed package to the **vendor** directory in the path specified by **GOPATH** in your project. Refresh the project to ensure that the copied file exists in the directory.
    
## Configuration
Start configuration with the **Client** structure in the **push** package. Each method in the **Client** structure can be used to call an API of the Push Kit server.
The following table describes the method in the **Client** structure.
| Method| Description|
| ----------- | ----------- |
|SendMessage|   Sends a message to a device.|

To use the functions provided by the packages in **examples**, set related parameters in [pushcommon.go](src/examples/common/pushcommon.go) in the common package.

The following table describes the parameters in [pushcommon.go](src/examples/common/pushcommon.go).
| Parameter| Description|
| ----------- | ----------- |
|appId|App ID, which is obtained from the app information.|
|appSecret|App secret, which is obtained from the app information.|
|authUrl|URL for Huawei OAuth 2.0 to obtain a token. For details, please refer to [OAuth 2.0-based Authentication](https://developer.huawei.com/consumer/en/doc/development/HMSCore-Guides/oauth2-0000001212610981).|
|pushUrl|Access address of Push Kit. For details, please refer to [Downlink Message Sending](https://developer.huawei.com/consumer/en/doc/development/HMSCore-Guides/android-server-dev-0000001050040110?ha_source=hms1).|

The following table describes the parameters in [pushcommon.go](src/examples/common/pushcommon.go).
| Parameter| Description|
| ----------- | ----------- |
|TargetTopic|Name of the topic to be subscribed to, unsubscribed from, or queried.|
|TargetCondition|Combined condition expression for sending a message.|
|TargetToken|Token of a target device, which is obtained from the target device.|
|TargetTokenArray|Tokens of all target devices, which are obtained from these target devices.|

## Sample Code

1). Send an Android data message.
You can obtain the initialized MessageRequest instance of the data message using the NewTransparentMsgRequest method in the push/model package.
> Code location: [send_data_message](src/examples/send_data_message/main.go)
    
2). Send an Android notification message.
You can obtain the initialized MessageRequest instance of the notification message using the NewNotificationMsgRequest method in the push/model package.
> Code location: [send_notify_message](src/examples/send_notify_message/main.go)
    
3). Send a message by topic.
You can send a notification message or data message to a device by topic. Specify the topic after obtaining the MessageRequest instance.
> Code location: [send_topic_message](src/examples/send_topic_message/main.go)
    
4). Send a message by conditions.
You can send a notification message or data message to a device by conditions. Specify the conditions after obtaining the MessageRequest instance.
> Code location: [send_condition_message](src/examples/send_condition_message/main.go)
    
5). Send a message to a Huawei quick app.
You can send a message to a quick app by setting FastAppTarget.
> Code location: [send_instance_app_message](src/examples/send_instance_app_message/main.go)
    
6). Send a message through the APNs agent.
You can send a message through the APNs agent by setting Apns of the message.
> Code location: [send_apns_message](src/examples/send_apns_message/main.go)
    
7). Send a message through the WebPush agent.
You can send a message through the WebPush agent by setting WebPush of the message.
> Code location: [send_webpush_message](src/examples/send_webpush_message/main.go)
    
8). Send a test message.
> Code location: [send_test_message](src/examples/send_test_message/main.go)

## Technical Support
You can visit the [Reddit community](https://www.reddit.com/r/HuaweiDevelopers/) to obtain the latest information about HMS Core and communicate with other developers.

If you have any questions about the sample code, try the following:
- Visit [Stack Overflow](https://stackoverflow.com/questions/tagged/huawei-mobile-services?tab=Votes), submit your questions, and tag them with `huawei-mobile-services`. Huawei experts will answer your questions.
- Visit the HMS Core section in the [HUAWEI Developer Forum](https://forums.developer.huawei.com/forumPortal/en/home?fid=0101187876626530001?ha_source=hms1) and communicate with other developers.

If you encounter any issues when using the sample code, submit your [issues](https://github.com/HMS-Core/hms-push-serverdemo-go/issues) or submit a [pull request](https://github.com/HMS-Core/hms-push-serverdemo-go/pulls).

## License
The sample code is licensed under [Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0).
