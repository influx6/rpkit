FROM golang:1.8

RUN sudo apt-get update && curl -sL https://deb.nodesource.com/setup_10.x | sudo -E bash - && apt-get install -y nodejs build-essential

COPY . /rpkit
WORKDIR /rpkit

RUN chmod +x ./scripts/tests.sh && npm install -g mocha chai

ENTRYPOINT ["bash ./scripts/tests.sh"]
