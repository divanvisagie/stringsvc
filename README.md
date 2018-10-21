# stringsvc
An implementation of the stringsvc from the [go-kit](https://gokit.io) tutorial

### Testing
```sh
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
# expect {"v":"HELLO, WORLD","err":null}
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
# expect {"v":12}
```