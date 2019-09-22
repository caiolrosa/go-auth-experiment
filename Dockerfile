FROM golang

WORKDIR /app

COPY . /app

RUN go build -o /bin/guardian-api .

CMD [ "/bin/guardian-api" ]

EXPOSE 3000
