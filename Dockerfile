FROM golang:alpine AS build
# create a working directory
WORKDIR /go/src/app
#COPY go.mod go.sum ./
#RUN go mod download
# add source code
COPY . .
# run main.go
ENV CGO_ENABLED=0
RUN go build -o main
#CMD ["go", "run", "main.go"]

FROM scratch

WORKDIR /go/src/app
COPY --from=build /go/src/app/main .
CMD ["./main"]
