FROM golang:latest

RUN go version
ENV GOPATH=/
ENV DB_PASSWORD=6aaf92d903eb3d2aab746d11eeb537e8acb4e707a16cb35536d4790ac9158b4f
ENV PORT=8000

COPY ./ ./

RUN go mod download
RUN go build -o ./bin/main ./main.go

CMD ["./bin/main"]