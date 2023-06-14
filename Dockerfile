########################
# BUILD CONTAINER
########################
FROM golang:alpine AS builder

RUN mkdir calc
WORKDIR /calc
COPY . . 
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o calculator-server ./server/server.go

########################
# RUN CONTAINER
########################
FROM alpine
WORKDIR /calc
COPY --from=builder /calc/calculator-server .
EXPOSE 50052

CMD [ "./calculator-server" ]