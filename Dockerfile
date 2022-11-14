FROM golang:1.16-alpineAS builder
WORKDIR /app
COPY .environment /tmp/.environment
COPY app/ .
RUN set -a; source /tmp/.environment; set +a
RUN apk update && apk add git openssh-client ; \
    go mod download && go get -u github.com/swaggo/swag/cmd/swag ; \
    go get github.com/alecthomas/template@v0.0.0-20190718012654-fb15b899a751; \
    go mod download golang.org/x/crypto
RUN swag init 
RUN env GOOS=linux GOARCH=amd64 go build main.go
CMD /app/main

############################
FROM scratch
WORKDIR /app
COPY --from=builder /app/notification-center /app/notification-center
ENTRYPOINT ["/app/notification-center"]
EXPOSE 8080