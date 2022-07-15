FROM golang:1.18

COPY . .

ENV GOPATH=/SysTray
ENV GO111MODULE=on

RUN apt-get update && \
     apt-get -y install gcc libgtk-3-dev libwebkit2gtk-4.0-dev libayatana-appindicator3-dev
RUN go build -ldflags="-linkmode internal"

CMD ["SysTray/SysTray"]



