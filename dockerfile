FROM golang:1.17
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main ./cmd/main.go
CMD ["./main"]




# docker run -it -p 4004:4004 -e API_URL="https://kitsu.io/api/edge/anime" -e API_PORT=":3000" -e BOT_API_KEY="OTEyNTA5NTc1MDg3OTI3MzI2.Gk0dFV.H96ZKQe4zpBqBNlCfirCgUQzniRBe7vh_KUqe8" betamedina/anime_seasson_bot:latest