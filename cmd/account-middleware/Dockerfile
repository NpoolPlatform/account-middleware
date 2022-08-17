FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY account-middleware /app

ENTRYPOINT ["/app/account-middleware"]

EXPOSE 50470 50471 50472 50473

CMD ["run"]
