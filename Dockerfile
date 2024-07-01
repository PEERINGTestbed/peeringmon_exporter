FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o /peeringmon_exporter

EXPOSE 2112

CMD [ "/peeringmon_exporter" ]
