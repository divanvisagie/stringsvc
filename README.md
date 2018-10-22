# stringsvc
An implementation of the [stringsvc](https://gokit.io/examples/stringsvc.html#first-principles) from the [go-kit](https://gokit.io) tutorial

### Testing
```sh
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
# expect {"v":"HELLO, WORLD","err":null}
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
# expect {"v":12}
```

### Running monitoring
```sh
docker-compose -f docker/docker-compose-prometheus-grafana.yml up -d
```

Prometheus: [localhost:9090](localhost:9090)

Grafana:    [localhost:3000](localhost:3000)