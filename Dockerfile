FROM golang:alpine

RUN apk update && apk upgrade

RUN apk add tzdata \
&& \
	cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime \
&& \
	echo "America/Sao_Paulo" > /etc/timezone \
&& \
	apk del tzdata

WORKDIR /go/src/app

COPY ./go.mod ./go.sum ./

RUN go mod download && go mod verify

COPY ./ ./

RUN go build -ldflags '-s -w' -o /go/bin/app ./cmd/

EXPOSE 8080

CMD [ "app" ]
