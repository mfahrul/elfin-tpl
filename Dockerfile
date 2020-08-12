FROM alpine
ADD elfin-tpl-service /elfin-tpl-service
ENTRYPOINT [ "/elfin-tpl-service" ]
