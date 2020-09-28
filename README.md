# HMS Pushkit Go Severdemo
## Table of Contents
* [Introduction](#introduction)
* [Installation](#installation)
* [Configuration](#configuration)
* [Sample Code](#sample-code)
* [License](#license)


## Introduction
Golang sample code encapsulates APIs of the HUAWEI Push Kit server. It provides many sample programs for your reference or usage.

The following describes packages of Golang sample code.
| Package   | Description |
| ----------- | ----------- |  
|examples|Sample code packages. Each package can run independently.|
|httpclient|Common package for sending network requests.|
|push|Package where APIs of the HUAWEI Push Kit server are encapsulated.|

## Installation
Before using Golang sample code, check whether the Golang environment has been installed. Golang 1.11 or a later version is recommended.
Decompress the Golang sample code package.
    
Copy the org.huawei.com package in the decompressed folder to the project vendor directory in the path specified by GOPATH.
Refresh the project and ensure that the file is successfully copied to the destination directory.
    
## Configuration 
Golang sample code uses the Client structure in the push package as the entry. Each method in the Client structure calls an API of the HUAWEI Push Kit server.
The following describes methods in the Client structure.
| Method   | Description |
| ----------- | ----------- |    
|SendMessage|   Sends a message to a device.|

To use functions provided by packages in examples, you need to set related parameters in pushcommon.go in the common package.

The following describes parameters in pushcommon.go.
| Parameter   | Description |
| ----------- | ----------- |    
|appId|App ID, which is obtained from app information.|
|appSecret|Secret access key of an app, which is obtained from app information.|
|authUrl|URL for the Huawei OAuth 2.0 service to obtain a token, please refer to [Generating an App-Level Access Token](https://developer.huawei.com/consumer/en/doc/development/parts-Guides/generating_app_level_access_token).|
|pushUrl|URL for accessing HUAWEI Push Kit, please refer to [Sending Messages](https://developer.huawei.com/consumer/en/doc/development/HMS-References/push-sendapi).|

The following table describes parameters in target.go. 
| Parameter   | Description |
| ----------- | ----------- | 
|TargetTopic|Name of a topic to be subscribed to, unsubscribed from, or queried.|
|TargetCondition|Combination of condition expressions for a message.|
|TargetToken|Token of a target device, which is obtained from the device.|
|TargetTokenArray|Tokens of all target devices, which are obtained from the devices.|

## Sample Code
Download Golang sample code in Downloading Server Sample Code.

1). Send an Android data message.
You can obtain the initialized MessageRequest instance of the data message using the NewTransparentMsgRequest method in the push/model package.
> Code location: examples/send_data_message/main.go
    
2). Send an Android notification message.
You can obtain the initialized MessageRequest instance of the notification message using the NewNotificationMsgRequest method in the push/model package.
> Code location: examples/send_notify_message/main.go
    
3). Send a message by topic.
You can send a notification message or data message to a device by topic. Specify the topic after obtaining the MessageRequest instance.
> Code location: examples/send_topic_message/main.go
    
4). Send a message by conditions.
You can send a notification message or data message to a device by conditions. Specify the conditions after obtaining the MessageRequest instance.
> Code location: examples/send_condition_message/main.go
    
5). Send a message to a Huawei quick app.
You can send a message to a quick app by setting FastAppTarget.
> Code location: examples/send_instance_app_message/main.go
    
6). Send a message through the APNs agent.
You can send a message through the APNs agent by setting Apns of the message.
> Code location: examples/send_apns_message/main.go
    
7). Send a message through the WebPush agent.
You can send a message through the WebPush agent by setting WebPush of the message.
> Code location: examples/send_webpush_message/main.go
    
8). Send a test message.
> Code location: examples/send_test_message/main.go

## Question or issues
If you want to evaluate more about HMS Core,
[r/HMSCore on Reddit](https://www.reddit.com/r/HuaweiDevelopers/) is for you to keep up with latest news about HMS Core, and to exchange insights with other developers.

If you have questions about how to use HMS samples, try the following options:
- [Stack Overflow](https://stackoverflow.com/questions/tagged/huawei-mobile-services) is the best place for any programming questions. Be sure to tag your question with 
`huawei-mobile-services`.
- [Huawei Developer Forum](https://forums.developer.huawei.com/forumPortal/en/home?fid=0101187876626530001) HMS Core Module is great for general questions, or seeking recommendations and opinions.

If you run into a bug in our samples, please submit an [issue](https://github.com/HMS-Core/hms-push-serverdemo-go/issues) to the Repository. Even better you can submit a [Pull Request](https://github.com/HMS-Core/hms-push-serverdemo-go/pulls) with a fix.

## License
Pushkit Go sample is licensed under the [Apache License, version 2.0](http://www.apache.org/licenses/LICENSE-2.0).

