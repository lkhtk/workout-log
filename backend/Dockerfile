FROM golang:1.23 AS build-env
WORKDIR /go/src/app/src/
COPY ./src/ ./
RUN go mod download && go get -d -v ./... &&\ 
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend ./ &&\
    mv backend /opt/backend.bin
USER nonroot
FROM scratch
ARG SHA
ENV BUILD_ID=$SHA
WORKDIR /opt
COPY --from=build-env /opt /opt
EXPOSE 8000
CMD ["/opt/backend.bin"]