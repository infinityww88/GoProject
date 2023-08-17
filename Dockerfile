FROM golang
RUN mkdir /app
COPY ./HttpGo /app
WORKDIR /app
CMD go run HttpGo

