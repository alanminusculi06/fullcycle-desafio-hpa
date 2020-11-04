FROM golang:1.15-alpine as builder

WORKDIR /src/

COPY main.go .

RUN go build main.go \
    && chmod +x main

FROM scratch

COPY --from=builder ./src/main ./main

EXPOSE 8000

CMD [ "./main" ]