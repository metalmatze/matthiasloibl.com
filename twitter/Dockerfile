FROM alpine:latest
EXPOSE 8080
ENTRYPOINT ["/usr/bin/twitter"]

RUN apk add --update ca-certificates
ADD twitter/twitter /usr/bin/twitter
