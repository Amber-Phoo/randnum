ARG VERSION=1.18

FROM golang:${VERSION} as builder

WORKDIR /app

ADD go.mod .
ADD go.sum .
ADD main.go .

RUN CGO_ENABLED=0 go build -o /app/main .

FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=builder /app/main .
ADD static static
ADD views views

ENV PORT="8080" 

EXPOSE ${PORT}

VOLUME /app/static
VOLUME /app/views

CMD [ "/app/main", "--mode=release", "--staticDir=/app/static" ]
