FROM golang

RUN go get github.com/gorilla/mux

RUN go get -v golang.org/x/oauth2/google

RUN go get -v google.golang.org/api/calendar/v3

WORKDIR /go/src/github.com/heaptracetechnology/google-calender

ADD . /go/src/github.com/heaptracetechnology/google-calender

RUN go install github.com/heaptracetechnology/google-calender

ENTRYPOINT google-calender

EXPOSE 3000