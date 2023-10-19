FROM golang:latest

WORKDIR /go/app

# lib necessário para rodar o kafka com o go
RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null"]
