# Build stage
FROM golang:alpine as build-stage

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update && apk --no-cache add make

WORKDIR /build 
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN make build

# Run stage
FROM scratch
COPY --from=build-stage /build/todoapi /
COPY --from=build-stage /build/postgres/migrations /postgres/migrations
EXPOSE 3000
CMD ["/todoapi"]