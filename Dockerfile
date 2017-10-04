FROM golang as builder

RUN wget "http://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-1.3.0.tar.gz" && \
   tar -zxv -C /usr/local/ -f libtensorflow-cpu-linux-x86_64-1.3.0.tar.gz && \
   rm -f libtensorflow-cpu-linux-x86_64-1.3.0.tar.gz

ENV LD_LIBRARY_PATH=/usr/local/lib/

ADD . /go/src/github.com/jnummelin/go-inception-client

WORKDIR  /go/src/github.com/jnummelin/go-inception-client

RUN go get github.com/golang/protobuf/ptypes/wrappers && \
    go get github.com/tensorflow/tensorflow/tensorflow/go && \
    go get google.golang.org/grpc && \
    go get github.com/gorilla/mux

RUN go build app.go inception_client.go


FROM ubuntu

RUN apt-get update && apt-get install -y wget && \
    wget "http://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-1.3.0.tar.gz" && \
    tar -zxv -C /usr/local/ -f libtensorflow-cpu-linux-x86_64-1.3.0.tar.gz && \
    rm -f libtensorflow-cpu-linux-x86_64-1.3.0.tar.gz

ENV LD_LIBRARY_PATH=/usr/local/lib/

COPY --from=builder /go/src/github.com/jnummelin/go-inception-client/app .

CMD ["./app"]