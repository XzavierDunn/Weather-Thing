From golang:latest as builder
LABEL maintainer="XD"
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM python:3-alpine
RUN apk --no-cache add ca-certificates
RUN python -m pip install --upgrade pip
RUN pip3 install requests
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
