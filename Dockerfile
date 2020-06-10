FROM golang:1.14

RUN mkdir /marsbot

ADD . /marsbot

WORKDIR /marsbot

CMD ["/marsbot/cmd/cmd"]
