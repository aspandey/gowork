FROM golang

#RUN go get -u github.com/spf13/cobra
#RUN go get go.mongodb.org/mongo-driver/mongo

RUN mkdir -p /usr/local/go/src/employee

COPY . /usr/local/go/src/employee

WORKDIR /usr/local/go/src/employee

RUN go build

ENTRYPOINT ["tail", "-f", "/dev/null"]