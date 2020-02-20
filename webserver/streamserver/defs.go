package main

const(
	VIDEO_DIR="./videos/"
	MAX_UPLOAD_SIZE=1024*1024*400 //400mb
	AV_DIR="./av/"
)

var M=NewConnLimiter(2)