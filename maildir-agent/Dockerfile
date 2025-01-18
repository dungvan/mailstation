FROM golang:1.22 as build

WORKDIR /go/src/app
COPY . .

RUN go build -o maildir-agent .

FROM golang:1.22

COPY --from=build /go/src/app/maildir-agent /usr/local/bin/maildir-agent

ENTRYPOINT [ "/usr/local/bin/maildir-agent" ]