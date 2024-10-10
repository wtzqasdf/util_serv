@echo off

:: into the current directory
cd %~dp0

:: set compile environments
set GOOS=linux
set GOARCH=arm64
set CGO_ENABLED=0

:: execute go build
go build -o ./build/util_serv