FROM golang
WORKDIR /app
COPY server.go /app/
COPY web /app/web/
RUN go build server.go
EXPOSE 8080
CMD ["./server"]