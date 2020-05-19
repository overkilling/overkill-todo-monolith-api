# Build stage
FROM golang:alpine as build-stage
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build 
COPY . .
RUN go mod download
RUN go build -o main .

# Run stage
FROM scratch
COPY --from=build-stage /build/main /
EXPOSE 3000
CMD ["/main"]