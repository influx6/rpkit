FROM golang:1.8

RUN apt-get update && curl -sL https://deb.nodesource.com/setup_10.x | bash - && apt-get install -y nodejs build-essential gcc g++ make

COPY . /go/src/github.com/gokit/rpkit
WORKDIR /go/src/github.com/gokit/rpkit

RUN chmod +x ./tests.sh && npm install -g mocha chai

CMD ["./tests.sh"]
