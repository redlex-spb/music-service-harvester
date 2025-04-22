FROM golang:1.22-alpine AS build
WORKDIR /src
COPY . .
RUN go mod download && go build -o /harvester ./cmd/harvester
FROM alpine:3.19
COPY --from=build /harvester /harvester
EXPOSE 8080
ENTRYPOINT ["/harvester"]
