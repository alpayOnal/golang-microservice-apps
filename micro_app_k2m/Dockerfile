FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN apt-get update  -y && apt-get install -y --no-install-recommends apt-utils bzip2
RUN tar xjf libs.tar.bz2
RUN cd libs/librdkafka && make clean && ./configure --prefix /usr && make  && make install

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main