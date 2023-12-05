FROM golang:1.21

WORKDIR /app

#COPY go.mod .
#COPY main.go .
COPY . .

#RUN go get
RUN go install
RUN go build -o bin .

EXPOSE 3333

ENTRYPOINT [ "/app/bin"]