FROM golang
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

ENV CGO_ENABLED=1
RUN go build -v -o ./app ./

CMD ["/app/app"]