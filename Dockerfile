FROM scratch
COPY go /go
COPY tmp /tmp
CMD ["/go/bin/go", "run", "/go/src/begin"]
