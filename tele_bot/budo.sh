#!/bin/bash 

source .telegram_sercrets
export YOUTUBE_DL_BINARY_PATH=`which youtube-dl`

# list of masters to obey comma seperated entries please
export MY_MASTERS="thapakazi"

# compiling binary stuffs
# GO_FILE='yt.go'
# GOOS=linux go build -ldflags="-s -w" ${GO_FILE}
# upx --brute yt
#./yt

# else just run basic 
go run yt.go
