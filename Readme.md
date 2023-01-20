# STRCASE

Forked from segmentio repo: go-camelcase and go-snakecase 


## go-camelcase

Fast camelcase implementation that avoids Go's regexps. Direct fork of our [snakecase](https://github.com/segmentio/go-snakecase) implementation.

--

    import "github.com/segmentio/go-camelcase"

Convert strings to camelCase

### Usage

#### func  Camelcase

```go
func CamelCase(str string) string
```
Camel case representation of `str`.

## License

 MIT



## go-snakecase

Fast snakecase implementation, believe it or not this was a large bottleneck in our application, Go's regexps are very slow.

## License

MIT