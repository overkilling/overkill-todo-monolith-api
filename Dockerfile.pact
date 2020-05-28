FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install Pact ruby cli tools
WORKDIR /deps 
RUN apk update \
  && apk --no-cache add ca-certificates wget bash \
  && wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub \
  && wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.31-r0/glibc-2.31-r0.apk \
  && apk add glibc-2.31-r0.apk
RUN wget https://github.com/pact-foundation/pact-ruby-standalone/releases/download/v1.84.0/pact-1.84.0-linux-x86_64.tar.gz \
  && tar xzf pact-1.84.0-linux-x86_64.tar.gz
ENV PATH /deps/pact/bin:$PATH

# Copy project and install dependencies
WORKDIR /build 
COPY . .
RUN go mod download

# Run pact tests
CMD go test -v ./... -tags pact -run TestPactProvider

