FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download \
    && go build -o ./out/test-app .

# CMD [ "./bin/test-app" ]