## Build
FROM golang:1.20-alpine AS buildenv

ADD go.mod /
RUN go mod download

COPY / /app/
WORKDIR /app

RUN go build -o bin ./cmd/main/

## Deploy
FROM scratch

WORKDIR /

COPY --from=buildenv /app/bin /app/

######
######
## Change here which input data to copy and use in program
COPY /test/csv_reader/db.csv /
COPY /test/json_reader/db.json /
ENV FILE_PATH="db.json"
######
######

CMD ["/app/bin"]