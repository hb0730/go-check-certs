FROM alpine
WORKDIR /app
COPY check-certs .
ENV FLAG ""
ENTRYPOINT /app/check-certs ${FLAG}