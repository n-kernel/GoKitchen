FROM golang

COPY . /go/src/github.com/Jeroenimoo/GoKitchen
WORKDIR "/go/src/github.com/Jeroenimoo/GoKitchen"
RUN go get -d -v
RUN go install -v

EXPOSE 8080

CMD ["go-wrapper", "run"]