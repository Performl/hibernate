FROM golang:1.19 as builder
WORKDIR /app
COPY ./pkg ./pkg
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download
RUN go build -o hibernate ./pkg/core


FROM golang:1.19 as runner
COPY --from=builder /app/hibernate /app/hibernate
ENTRYPOINT ["/app/hibernate"]