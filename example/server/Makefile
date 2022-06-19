PROJECT_NAME := jayce_server
VERSION := v1

build:
	go build -o ${PROJECT_NAME}

image:
	docker build -t ${PROJECT_NAME}:${VERSION} .

run:
	docker run -d --net=host ${PROJECT_NAME}:${VERSION}

stop:
	docker stop ${PROJECT_NAME}:${VERSION}
