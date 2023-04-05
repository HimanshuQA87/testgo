FROM golang:1.20.2-alpine3.16
RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest
#COPY go.mod go.sum ./
WORKDIR /go/src
RUN mkdir testgo

#COPY *.go /go/src
#COPY ./ /go/src/testgo
COPY ./ /go/src/testgo
WORKDIR /go/src/testgo
CMD ["ginkgo"]