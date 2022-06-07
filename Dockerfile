FROM golang:1.18

WORKDIR /usr/src/app

RUN ls $GOPATH

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/link-collector ./main.go

EXPOSE 8420

CMD ["link-collector"]
