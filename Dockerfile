FROM golang:1.7

RUN mkdir -p /opt/printpot
WORKDIR /opt/printpot

ADD printpot.go printpot.go

RUN go build printpot.go

EXPOSE 9100

CMD /opt/printpot/printpot
