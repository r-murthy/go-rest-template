FROM golang:1.17rc1-alpine

WORKDIR /app/

COPY . /app/

RUN go mod download
RUN go mod vendor

RUN CGO_ENABLED=0 go build -o /credit

EXPOSE 8001

CMD [ "/credit" ]