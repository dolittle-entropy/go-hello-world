# Start server

## Without Prefix
```sh
PORT=1234 go run main.go
```

## With Prefix
```sh
URL_PREFIX="/yo" PORT=1234 go run main.go
```

# Query
## Without Prefix

```sh
PORT=1234 curl "localhost:1234/"
```

```sh
PORT=1234 curl "localhost:1234/hi"
```

```sh
PORT=1234 curl "localhost:1234/hello-world"
```

# With Prefix
```sh
PORT=1234 curl "localhost:1234/yo"
```

```sh
PORT=1234 curl "localhost:1234/yo/"
```

```sh
PORT=1234 curl "localhost:1234/yo/hi"
```

```sh
PORT=1234 curl "localhost:1234/yo/hello-world"
```


# Docker

## Build
```sh
docker build -f ./Dockerfile -t hello-world .

docker tag hello-world dolittle/hello-world:latest
docker push dolittle/hello-world:latest
```