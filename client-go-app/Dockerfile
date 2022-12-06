FROM golang:alpine3.16

COPY . /k8s-client-go-app

WORKDIR /k8s-client-go-app

RUN go build -o main && \
    chmod -R 777 /k8s-client-go-app

CMD ["./main"]