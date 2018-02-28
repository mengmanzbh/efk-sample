# base image
FROM golang:alpine

# set entry work dir
WORKDIR /go/src/app

# copy program to the image
COPY ./helloworld.go /go/src/app

# run command
CMD [ "go", "run", "helloworld.go" ]
