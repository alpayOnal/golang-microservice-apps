FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN apt-get update  -y && apt-get install -y --no-install-recommends apt-utils bzip2
RUN cp libs.tar.bz2 /tmp
RUN cd /tmp && tar xjf libs.tar.bz2
RUN cd /tmp/libs/librdkafka && make clean && ./configure --prefix /usr && make  && make install

EXPOSE 8000

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main