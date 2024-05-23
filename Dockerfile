FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./cmd/lambda

FROM public.ecr.aws/lambda/go:1

COPY --from=builder /app/main ${LAMBDA_TASK_ROOT}

CMD ["main"]
