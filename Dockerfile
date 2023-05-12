# FROM golang:1.20.1-alpine3.17

# WORKDIR /app

# # COPY go.mod ./
# # COPY go.sum ./
# COPY . .

# RUN go build -o main main.go

# EXPOSE 8080

# CMD [ "/app/main" ]


FROM golang:1.20.1-alpine3.17 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /

COPY --from=build-stage /app /
EXPOSE 8080

# CMD [ "/docker-gs-ping/main.go" ]
#USER nonroot:nonroot
CMD ["/main"]