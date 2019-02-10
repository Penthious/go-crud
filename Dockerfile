FROM golang

WORKDIR /go-crud

RUN GO111MODULE=on go build -mod=vendor -o go-crud

COPY . ./

# Building using -mod=vendor, which will utilize the v

CMD ["./go-crud"]