FROM docker.io/mbrt/golang-vim-dev

USER root

WORKDIR /go/src/go-buried-point

COPY . /go/src/go-buried-point

ENV TERM=xterm

EXPOSE 8080

ENTRYPOINT ["sh", "/go/src/go-buried-point/docker/entrypoint.sh"]
