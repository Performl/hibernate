FROM golang:1.19

WORKDIR /app
COPY ./hibernate ./hibernate

ENTRYPOINT ["./hibernate"]