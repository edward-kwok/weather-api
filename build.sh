#!/bin/bash


go get -t -v -u github.com/gorilla/mux
go get -t -v -u github.com/gocolly/colly

go build -o app weather/weather.go