FROM golang:1.13 AS buildenv
ENV CGO_ENABLED=0
RUN mkdir /src
ADD . /src
WORKDIR /src
RUN go install github.com/go-swagger/go-swagger/cmd/swagger
RUN make

FROM gcr.io/distroless/base-debian10
COPY --from=buildenv /src/.bin/interview-api-server /
EXPOSE 8080
ENTRYPOINT ["/interview-api-server", "--port", "8080"]
