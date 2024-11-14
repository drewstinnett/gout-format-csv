# Gout Format CSV

Provides a CSV formatter for use with [Gout](https://github.com/drewstinnett/gout)

## Custom Headers

By default, we use the csvutil auto headers. If you provide the Formatter object
with a list of custom headers, the csv output will use that instead

```go
Formatter{
	headers []string{"foo", "bar", "baz"},
}
```