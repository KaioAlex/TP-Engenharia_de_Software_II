FROM golang:alpine 


WORKDIR /app

COPY ./ /app

RUN apk add inotify-tools
RUN apk add git
ENV CGO_ENABLED 0
COPY startScript.sh /build/startScript.sh

RUN git clone https://github.com/go-delve/delve.git && \
    cd delve && \
    go install github.com/go-delve/delve/cmd/dlv

RUN go mod tidy 

RUN go build -o /server -gcflags -N -gcflags -l -buildvcs=false

EXPOSE 3000
EXPOSE 30000

ENTRYPOINT sh startScript.sh