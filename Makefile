APPNAME ?= dumbweb
TAG ?= latest
REPO ?= pstuart
IMAGE ?= ${REPO}/${APPNAME}:${TAG}

.phony:	build docker docker upx all push redo

# note that the build disables CGO to ensure a static binary with zero dependencies
# which will be necessary when using a scratch image

build:
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ${APPNAME}

docker:
	docker build -t ${IMAGE} .

upx:
	upx dumbweb

run:
	docker run -it --rm -p 8080:8080 ${IMAGE}

all:	build upx docker

push:
	docker push ${IMAGE}

redo:	all push
