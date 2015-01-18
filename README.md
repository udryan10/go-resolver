# go-resolver
Simple utility that will resolve domains and output status. Intended to be used against large sets of domains for mass resolving.

<img width="10%" src="https://raw.github.com/golang-samples/gopher-vector/master/gopher.png"/>

## requirement
domains to resolve must be in a file, 1 per line:

```
www.google.com
www.yahoo.com
www.yahoo1.com
www.yahoo23.com
www.slickdeals.com
```

### build
`go build resolver.go`

## running
Run the binary supplying the path of the file that contains your domains:

`./resolver /tmp/hosts.txt`
