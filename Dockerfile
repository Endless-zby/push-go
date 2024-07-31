FROM ubuntu:latest
LABEL authors="byzhao"

ENTRYPOINT ["top", "-b"]