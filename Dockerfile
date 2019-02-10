FROM golang:alpine

RUN apk add --update tzdata \
    bash wget curl git;

RUN mkdir -p $GOPATH/bin && \
    curl https://glide.sh/get | sh && \
    go get github.com/pilu/fresh

WORKDIR /go/src/go-crud
COPY . ./

RUN GO111MODULE=on go mod vendor
RUN go build

COPY . ./
EXPOSE 8000

CMD glide update && fresh -c runner.conf main.go