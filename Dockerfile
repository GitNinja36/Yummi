# start
FROM golang:1.24.3-alpine

# working directory
WORKDIR /app

# copy go.mod and go.sum first 
COPY go.mod go.sum ./

# download dependencies
RUN go mod download

# copy
COPY . .

# build Go app
RUN go build -o main .

# port
EXPOSE 8080

# command to run 
CMD ["./main"]