FROM golang:latest
RUN mkdir /build
COPY . /build
WORKDIR /build
RUN go build ./main.go
EXPOSE 8081
ENTRYPOINT [ "/build/main" ]