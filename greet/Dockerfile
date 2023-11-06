FROM golang:1.20

WORKDIR /usr/src/app 

COPY . .

EXPOSE 8888

CMD [ "go", "run", "greet.go" ]


# # syntax=docker/dockerfile:1

# FROM golang:1.20

# # Set destination for COPY
# WORKDIR /app

# # Download Go modules
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./

# # Build
# RUN go build greet.go

# # Optional:
# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can document in the Dockerfile what ports
# # the application is going to listen on by default.
# # https://docs.docker.com/engine/reference/builder/#expose
# EXPOSE 8888

# # Run
# CMD ["/greet"]