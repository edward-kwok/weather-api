FROM golang:alpine AS build-env
COPY . /go/src/github.com/edward-kwok/weather-api
WORKDIR /go/src/github.com/edward-kwok/weather-api
RUN apk add --update --no-cache git 
#RUN go get -u github.com/gocolly/colly/... && go get -u github.com/gorilla/mux/...
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./weather

FROM alpine
WORKDIR /apps
COPY --from=build-env /go/src/github.com/edward-kwok/weather-api/app /apps
CMD [ "./app" ]
