# Notification-center
> Notification-center is one entry point for services in local network which want to send notification to telegram channel. Realised for Grafana, Redash. And ofcouse any servicec can send message by API.

___
## Build and run application
> Build binary file with all dependency and generate swagger documentation.

```Bash
go mod vendor
swag init 
go build -o notification-center
./notification-center
```
---
## Required variables

> Required enviroment variables: 
+ BotTokenDevops - Token for devops bot			
+ BotTokenMarkirovka - Token for markirovka bot	
+ TelegramURL -	URL of telegram			
+ ChatIdMarkirovka - Chat id of Markirovka
+ ChatIdEcomm - Chat id of Ecomm
+ ChatIdDevops - Chat id of Devops 
+ ChatIdOVSSupport - - Chat id of OVSSupport
+ ChatIdOVSDBA - Chat id of OVSDBA
+ Authorization - Token for autorization in notification-center
+ GrafanaUser -	Grafana user 
+ GrafanaPassword -	Password of grafana user  	
+ RedashUser - Redash user 
+ RedashPassword - Password of redash user  
---
## Docker

> Build Docker image from a [Dockerfile](https://github.com/7visij7/notification-center/blob/master/Dockerfile)
```
docker build -t IMAGENAME
```
> Start application
```
docker run -it --rm IMAGENAME -p 8080:8080
```
---
## Kubernetes

> Deploy to kubernetes cluster with [helm](https://github.com/7visij7/notification-center/tree/master/helm).
```Bash
helm upgrade --atomic notification-center notification-center/ --namespace devops  --debug --timeout 2m --wait
```
___
## How it works:
To send message to telegram chart for devops:
```Bash
curl -X POST "http://notification-center.infra.local/api/v1/telegram/devops" -H  "Content-Type: application/json" -H "Authorization:$TOKEN" -d  '{"text":"Hello world"}```