## We specify the base image we need for our
## go application

FROM golang:1.20


## We specify that we now wish to execute 
## any further commands inside our /app
## directory

WORKDIR /go/src/personel-resume-api

COPY . .

## Add this go mod download command to pull in any dependencies

RUN go get -v

## we run go build to compile the binary
## executable of our Go program

RUN go build -o main .

#EXPOSE the port
EXPOSE 8000

## Our start command which kicks off
## our newly created binary executable

CMD ["./main"]

