FROM golang:1.21.4 as build
WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

#RUN apt-get update && \
#    apt-get --yes --no-install-recommends install make="4.3-4.1" && \
#    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN go build -o banner-svc ./services/banner/cmd

FROM alpine:latest as production

COPY --from=build /app/banner-svc ./

COPY --from=build /app/.docker.env ./.env
COPY --from=build /app/services/banner/migrations ./migrations
COPY --from=build /app/services/banner/config/config.docker.yml ./services/banner/config/config.local.yml

CMD ["./banner-svc"]

EXPOSE 8080