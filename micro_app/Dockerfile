FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apt-get update  -y && apt-get install -y --no-install-recommends apt-utils bzip2
RUN tar -xjf libs.tar.bz2
RUN cd libs/librdkafka && make clean && ./configure --prefix /usr && make  && make install
RUN go build .

EXPOSE 8000

#CMD ["./main"]