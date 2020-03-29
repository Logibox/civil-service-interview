# Interview question for DWP

## Building

Requires go 1.13, make, and go-swagger

```
make
.bin/interview-api-server --port 8080
curl localhost:8080/v1/city/London/users?within=50miles
```

### Installing go swagger
`go install github.com/go-swagger/go-swagger/cmd/swagger`

## Building with docker

```
docker build . -t interview
docker run --rm -p 8080:8080 localhost/interview
curl localhost:8080/v1/city/London/users?within=50miles
```
