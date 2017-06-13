FROM golang:1.8.3 as builder
WORKDIR /golang/src/gcached
ADD . /golang/src/gcached
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gcached .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /golang/src/gcached/gcached .
EXPOSE 3000
CMD ["./gcached"]  