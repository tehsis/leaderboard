FROM golang:1.8

WORKDIR /go/src/github.com/tehsis/leaderboard
ADD . /go/src/github.com/tehsis/leaderboard
RUN go get  gopkg.in/redis.v5

