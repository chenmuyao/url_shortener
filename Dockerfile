FROM alpine
COPY url_shortener /app/url_shortener
WORKDIR /app
CMD [ "/app/url_shortener" ]
