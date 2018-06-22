FROM golang as build

ADD . /go/src/github.com/smartsiva/filehandler
RUN go get -d -v github.com/smartsiva/filehandler
WORKDIR /go/src/github.com/smartsiva/filehandler
RUN go build -o filehandler .

FROM ubuntu
COPY --from=build /go/src/github.com/smartsiva/filehandler/filehandler .
EXPOSE 8087
ENTRYPOINT ["./filehandler"]
