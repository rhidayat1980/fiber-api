FROM golang

ENV PORT 3000

EXPOSE 3000

RUN mkdir /app

ADD . /app

WORKDIR /app/src/

RUN go mod tidy 

RUN go build -o fiber-api .

CMD ["fiber-api"]