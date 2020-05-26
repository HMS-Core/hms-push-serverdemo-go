Huawei Go SDK
Table of Contents
Introduction
Installation Guide
Golang method definition
Configuration parameters
Sample Code
License and Terms

1. Introduction
    Golang sample code encapsulates APIs of the HUAWEI Push Kit server. It provides many sample programs for your reference or usage.
    The following describes packages of Golang sample code.
    
    examples:       Sample code packages. Each package can run independently.
    httpclient:     Common package for sending network requests.
    push:           Package where APIs of the HUAWEI Push Kit server are encapsulated.

2. Installation Guide
    Before using Golang sample code, check whether the Golang environment has been installed. Golang 1.11 or a later version is recommended.
    Decompress the Golang sample code package.
    
    Copy the org.huawei.com package in the decompressed folder to the project vendor directory in the path specified by GOPATH.
    Refresh the project and ensure that the file is successfully copied to the destination directory.
    
3. Golang method definition
    Golang sample code uses the Client structure in the push package as the entry. Each method in the Client structure calls an API of the HUAWEI Push Kit server.
    The following describes methods in the Client structure.
    
    SendMessage:    Sends a message to a device.

4. Configuration parameters
    To use functions provided by packages in examples, you need to set related parameters in pushcommon.go in the common package.
    The following describes parameters in pushcommon.go.
    
    appId:      App ID, which is obtained from app information.
    appSecret:  Secret access key of an app, which is obtained from app information.
    authUrl:    URL for the Huawei OAuth 2.0 service to obtain a token, please refer to Generating an App-Level Access Token.
    pushUrl:    URL for accessing HUAWEI Push Kit, please refer to Sending Messages.
    
    TargetTopic:        Name of a topic to be subscribed to, unsubscribed from, or queried.
    TargetCondition:    Combination of condition expressions for a message.
    TargetToken:        Token of a target device, which is obtained from the device.
    TargetTokenArray:   Tokens of all target devices, which are obtained from the devices.

5. Sample Code
    Download Golang sample code in Downloading Server Sample Code.

    1). Send an Android data message.
    You can obtain the initialized MessageRequest instance of the data message using the NewTransparentMsgRequest method in the push/model package.
    Code location£ºexamples/send_data_message/main.go
    
    2). Send an Android notification message.
    You can obtain the initialized MessageRequest instance of the notification message using the NewNotificationMsgRequest method in the push/model package.
    Code location£ºexamples/send_notify_message/main.go
    
    3). Send a message by topic.
    You can send a notification message or data message to a device by topic. Specify the topic after obtaining the MessageRequest instance.
    Code location£ºexamples/send_topic_message/main.go
    
    4). Send a message by conditions.
    You can send a notification message or data message to a device by conditions. Specify the conditions after obtaining the MessageRequest instance.
    Code location£ºexamples/send_condition_message/main.go
    
    5). Send a message to a Huawei quick app.
    You can send a message to a quick app by setting FastAppTarget.
    Code location£ºexamples/send_instance_app_message/main.go
    
    6). Send a message through the APNs agent.
    You can send a message through the APNs agent by setting Apns of the message.
    Code location£ºexamples/send_apns_message/main.go
    
    7). Send a message through the WebPush agent.
    You can send a message through the WebPush agent by setting WebPush of the message.
    Code location£ºexamples/send_webpush_message/main.go
    
    8). Send a test message.
    Code location£ºexamples/send_test_message/main.go

6.License and Terms
    Huawei Go SDK is licensed under the Apache License, version 2.0.

