#Build Stage 
FROM golang:1.17-alpine as builder

RUN mkdir /build
WORKDIR /build
COPY . .
RUN go build -o main cmd/main.go


#Run Stage 
FROM alpine
# ENV Vars
ARG ARG_PASS
ENV DB_PASS=${ARG_PASS}
WORKDIR /build
COPY --from=builder /build/main .
EXPOSE 8000

ENTRYPOINT [ "/build/main" ]