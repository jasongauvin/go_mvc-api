# Start from golang base image
FROM golang:1.15-alpine

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Jason Gauvin <contact.jason.gauvin@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /api

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download -x
RUN go get github.com/githubnemo/CompileDaemon

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

#Command to build the Go app with hot reload
ENTRYPOINT CompileDaemon -exclude-dir=.git --build="go build api/main.go" --command=./main