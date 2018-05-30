FROM golang:1.8

RUN apt-get update && curl -sL https://deb.nodesource.com/setup_10.x | bash - && apt-get install -y nodejs build-essential gcc g++ make

COPY . /rpkit
WORKDIR /rpkit

RUN chmod +x ./scripts/tests.sh && npm install -g mocha chai

ENTRYPOINT ["./scripts/tests.sh"]
