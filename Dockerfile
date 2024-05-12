FROM golang:1.22.3

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# ENV AIR_FLAGS=-buildvcs=false
ENV GOFLAGS=-buildvcs=false
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]