# First stage, build the executable
FROM golang:1.17

ENV GOPROXY=https://goproxy.cn
ENV HOME=/opt/app

WORKDIR $HOME

COPY go.mod $HOME
COPY go.sum $HOME
RUN go mod download

COPY . $HOME
# Use static linking to get rid of the error below
# exec user process caused "no such file or directory"
RUN GOOS=linux GOARCH=amd64 go build -a -ldflags "-linkmode external -extldflags '-static' -s -w"

EXPOSE 80

CMD ["/opt/app/dianasdog", "-host", "0.0.0.0", "-port", "80"]
