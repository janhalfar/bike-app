##############################
###### STAGE: BUILD     ######
##############################
FROM golang:alpine as builder

WORKDIR /app

COPY ./go.mod  ./

RUN go mod download

COPY . ./

RUN go build -ldflags "-X main.Version=`git rev-parse --short HEAD`" -o /server server/main.go

##############################
###### STAGE: PACKAGE   ######
##############################
FROM alpine

COPY --from=builder /server /bin/server

COPY htdocs /htdocs

EXPOSE      80

ENTRYPOINT  [ "/bin/server" ]