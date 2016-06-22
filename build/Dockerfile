FROM dilgerm/rpi-golang:1.6.2
ENV GO15VENDOREXPERIMENT 1
RUN apt-get update && apt-get install -y libpcap-dev python-requests time file
RUN go get github.com/golang/lint/golint github.com/fzipp/gocyclo github.com/client9/misspell/cmd/misspell
RUN go clean -i net && go install -tags netgo std
RUN go install -tags netgo std
COPY build.sh /
ENTRYPOINT ["sh", "/build.sh"]
