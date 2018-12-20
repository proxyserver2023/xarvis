# We use the official golang image, which contains all the 
# correct build tools and libraries. Notice `as builder`,
# this gives this container a name that we can reference later on. 
FROM golang:1.11-alpine as builder
RUN apk add git

# Set our workdir to our current service in the gopath
WORKDIR /go/src/github.com/alamin-mahamud/gapi

# Copy the current code into our workdir
COPY go.mod .
COPY go.sum .

# get deps
ENV GO111MODULE on
RUN go mod download

COPY . .

# Build the binary, with a few flags which will allow
# us to run this binary in Alpine. 
RUN     CGO_ENABLED=0 \
        GOOS=linux    \
        GOARCH=amd64  \
        go build -a -installsuffix cgo -o app cmd/api/main.go

# Here we're using a second FROM statement, which is strange,
# but this tells Docker to start a new build process with this
# image.
FROM alpine:latest

# Security related package, good to have.
RUN apk add --update --no-cache ca-certificates git

# Here, instead of copying the binary from our host machine,
# we pull the binary from the container named `builder`, within
# this build context. This reaches into our previous image, finds
# the binary we built, and pulls it into this container. Amazing!
COPY --from=builder /go/src/github.com/alamin-mahamud/gapi/app /bin/app

# Run the binary as per usual! This time with a binary build in a
# separate container, with all of the correct dependencies and
# run time libraries.
ENTRYPOINT ["/bin/app"]