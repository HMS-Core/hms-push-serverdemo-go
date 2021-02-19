# 华为推送服务服务端Golang示例代码

## 目录
 * [简介](#简介)
 * [安装](#安装)
 * [配置](#配置)
 * [示例代码](#示例代码)
 * [授权许可](#授权许可)

## 简介
Golang示例代码对华为推送服务（HUAWEI Push Kit）服务端接口进行封装，包含丰富的示例程序，方便您参考或直接使用。

示例代码主要包括以下组成：
| 包名   | 说明 |
| ----------- | ----------- |  
|examples|示例代码包，每个包都可以独立运行|
|httpclient|发送网络请求的公共包|
|push|推送服务的服务端接口封装包|

## 安装
使用Golang示例代码前，请确认您已安装Golang开发环境（推荐使用Golang 1.11或以上版本），并解压Golang示例代码包。
    
将解压包中的org.huawei.com包复制到GOPATH指定的工程vendor目录下。刷新该工程，确保文件出现在对应的目录下。
    
## 配置 
Golang示例代码以push包中的Client结构体为入口。Client结构体中的每个方法都可以调用推送服务服务端的一个接口。
Client结构体包含以下方法：
| 方法   | 说明 |
| ----------- | ----------- |    
|SendMessage|   向设备发送消息|

如需使用examples中的功能，请在公共包中的pushcommon.go中设置相关参数。

pushcommon.go中配置如下参数：
| 参数   | 说明 |
| ----------- | ----------- |    
|appId|应用ID，从应用消息中获取|
|appSecret|应用访问密钥，从应用信息中获取|
|authUrl|华为OAuth 2.0获取token的地址。具体请参考[基于OAuth 2.0开放鉴权-客户端模式](https://developer.huawei.com/consumer/cn/doc/development/parts-Guides/generating_app_level_access_token).|
|pushUrl|推送服务的访问地址。具体请参考[推送服务-下行消息](https://developer.huawei.com/consumer/cn/doc/development/HMS-References/push-sendapi).|

target.go中配置如下参数：
| 参数   | 说明 |
| ----------- | ----------- | 
|TargetTopic|订阅、退订或查询的主题名称|
|TargetCondition|消息的条件表达式组合|
|TargetToken|目标设备token，从目标设备上获取|
|TargetTokenArray|所有目标设备tokens，从目标设备上获取|

## 示例代码
请在[服务端示例代码](https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Examples-V5/server-sample-code-0000001050986079-V5)下载Golang示例代码.

1). 发送Android透传消息。
通过push/model包中的NewTransparentMsgRequest方法获取初始化的透传消息的MessageRequest实例。
> 代码位置: examples/send_data_message/main.go
    
2). 发送Android通知栏消息。
通过push/model包中的NewNotificationMsgRequest方法获取初始化的透传消息的MessageRequest实例。
> 代码位置: examples/send_notify_message/main.go
    
3). 基于主题发送消息。基于主题发送通知栏消息或透传消息。
获取MessageRequest实例后可定制主题。
> 代码位置: examples/send_topic_message/main.go
    
4). 基于条件发送消息。
基于条件发送通知栏消息或透传消息。获取MessageRequest实例后可定制条件。
> 代码位置: examples/send_condition_message/main.go
    
5). 向华为快应用发送消息。
通过设置FastAppTarget属性来实现。
> 代码位置: examples/send_instance_app_message/main.go
    
6). 基于APNs代理发送消息。
通过设置消息的Apns属性来实现。
> 代码位置: examples/send_apns_message/main.go
    
7). 基于WebPush代理发送消息。
通过设置消息的WebPush属性来实现。
> 代码位置: examples/send_webpush_message/main.go
    
8). 发送测试消息。
> 代码位置: examples/send_test_message/main.go

## 技术支持
如果您对HMS Core还处于评估阶段，可在[Reddit社区](https://www.reddit.com/r/HuaweiDevelopers/)获取关于HMS Core的最新讯息，并与其他开发者交流见解。

如果您对使用HMS示例代码有疑问，请尝试：
- 开发过程遇到问题上[Stack Overflow](https://stackoverflow.com/questions/tagged/huawei-mobile-services)，在`huawei-mobile-services`标签下提问，有华为研发专家在线一对一解决您的问题。
- 到[华为开发者论坛](https://developer.huawei.com/consumer/cn/forum/blockdisplay?fid=18) HMS Core板块与其他开发者进行交流。

如果您在尝试示例代码中遇到问题，请向仓库提交[issue](https://github.com/HMS-Core/hms-push-serverdemo-go/issues)，也欢迎您提交[Pull Request](https://github.com/HMS-Core/hms-push-serverdemo-go/pulls)。

## 授权许可
华为推送服务Golang示例代码经过[Apache License, version 2.0](http://www.apache.org/licenses/LICENSE-2.0)授权许可。

