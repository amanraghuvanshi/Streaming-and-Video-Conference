FROM golang:1.19.4-alpine3.17
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY . . 

RUN CGO_ENABLED=0 go build -o /bin/app ./cmd

FROM alpine
WORKDIR /src

COPY --from=0 /bin/app /bin/app
COPY --from=0 /src/views /src/views
COPY --from=0 /src/assests /src/assests

ENTRYPOINT [ "/bin/app" ]