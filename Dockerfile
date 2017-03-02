# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.8.0-alpine

# Update image and install git
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/alecholmez/binary

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get -u github.com/kardianos/govendor
RUN cd /go/src/github.com/alecholmez/binary && govendor sync
RUN go install github.com/alecholmez/binary

# Run the image as a non-root user
RUN adduser -D binuser
USER binuser

# Run the outyet command by default when the container starts.
CMD /go/bin/binary --bind 0.0.0.0:$PORT
