FROM golang:1.23.1
WORKDIR /go/src/app
COPY . ./
RUN go mod tidy \
    && go mod verify
RUN go build -o /main
CMD ["/main"]