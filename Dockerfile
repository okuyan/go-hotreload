FROM golang:1.12
EXPOSE 8080
WORKDIR /app
ENV GO111MODULE=on
COPY . .
RUN go get github.com/pilu/fresh
CMD ["fresh"]
