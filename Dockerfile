FROM golang:1.22.4

WORKDIR /app

COPY . .

EXPOSE 8080

CMD ["go", "run", "."]