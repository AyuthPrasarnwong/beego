FROM golang

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

RUN mkdir -p /go/src/api
COPY . /go/src/api
WORKDIR /go/src/api

RUN go get github.com/beego/bee
RUN go get github.com/astaxie/beego
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/smartystreets/goconvey
RUN go get github.com/dgrijalva/jwt-go

CMD bee run -downdoc=true -gendoc=true
