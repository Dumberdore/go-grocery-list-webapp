FROM golang:1.23-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/api/main.go

FROM alpine:3.20.1 AS prod
WORKDIR /app
COPY --from=build /app/main /app/main
ENV PORT=8080
EXPOSE 8080
CMD ["./main"]

FROM node:20 AS frontend
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install
COPY frontend .
EXPOSE 5173
CMD ["npm", "run", "dev"]
