FROM golang:1.22@sha256:c4fb952e712efd8f787bcd8e53fd66d1d83b7dc26adabc218e9eac1dbf776bdf as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go get -u golang.org/x/lint/golint
RUN go install golang.org/x/lint/golint

COPY . ./

RUN ${GOPATH}/bin/golint -set_exit_status ./...
RUN go test -v ./...

RUN CGO_ENABLED=0 go build -o /app/calc ./cmd

FROM golang:1.22@sha256:c4fb952e712efd8f787bcd8e53fd66d1d83b7dc26adabc218e9eac1dbf776bdf

WORKDIR /app
COPY --from=builder /app/calc /calc

EXPOSE 8080

COPY config /app/config
ENTRYPOINT ["/calc"]
