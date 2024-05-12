FROM golang:1.22.3

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

ENV AIR_FLAGS=-buildvcs=false
ENV GOFLAGS=-buildvcs=false
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin


# CMD ["go", "run", "main.go"]
CMD ["air"]