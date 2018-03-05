# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/alandwiprasetyo/rest-api

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/alandwiprasetyo/rest-api

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/rest-api

# Document that the service listens on port 8080.
EXPOSE 80