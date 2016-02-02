FROM scratch
MAINTAINER Song Gao <song@gao.io>
ADD . /
EXPOSE 80
EXPOSE 22
ENTRYPOINT ["/me"]
