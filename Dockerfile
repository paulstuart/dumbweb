FROM scratch

ADD dumbweb /go/bin/dumbweb

ENTRYPOINT ["/go/bin/dumbweb"]
#CMD ["/go/bin/dumbweb"]

