FROM centos:7

ADD ./resources/images /images
ADD ./resources/css /css
ADD ./bin/godocker_linux /

CMD ["/godocker_linux"]
EXPOSE 8888