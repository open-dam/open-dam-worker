FROM golang:1.14-alpine AS build

# Install dependencies
RUN apk update && apk add --no-cache \
    # ca-certificates is required to call HTTPS endpoints.
    ca-certificates \
    # exiftool is requred to extract metadata from files
    exiftool \
    # git is required for fetching the dependencies.
    git 

RUN update-ca-certificates

# Copy project files
WORKDIR /go/src
COPY go.mod .
COPY go.sum .
RUN mkdir /dir

# Fetch dependencies
RUN go mod download
COPY . .

# Build project
ENV CGO_ENABLED=0
RUN go build -o open-dam-worker ./cmd/open-dam-worker

FROM alpine AS runtime

# Copy dependent files
COPY --from=build /go/src/open-dam-worker ./
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN apk update && apk add --no-cache \
    # exiftool is requred to extract metadata from files
    exiftool 

RUN mkdir /dir

ENTRYPOINT ["./open-dam-worker"]
