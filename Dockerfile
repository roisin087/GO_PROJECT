# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /GO_TEST/src/github.com/roisin087/

# Build the servercommand inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/roisin087/server/main

# Run the server command by default when the container starts.
ENTRYPOINT /GO_TEST/bin/

# Document that the service listens on port 8181.
EXPOSE 8181