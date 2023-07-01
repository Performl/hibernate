FROM golang:1.19 as builder
WORKDIR /app
COPY ./pkg ./pkg
RUN go mod download
RUN go build -o hibernate ./pkg/core


FROM golang:1.19 as runner
COPY --from=builder /app/hibernate /app/hibernate
ENTRYPOINT ["/app/hibernate"]