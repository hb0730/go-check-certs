FROM alpine
WORKDIR /app
ENV FLAG ""
COPY check-certs .
ENTRYPOINT /app/check-certs ${FLAG}