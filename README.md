# is-it-hacked

[![Godoc][godoc-image]][godoc-url]

A tool using the [IsItHacked.com](https://www.isithacked.com/sites) service that checks if your site is disguised for the GoogleBot, contains spam links, funny redirects, or otherwise looks hacked. Powered by Joomla, Drupal, WordPress, Magento, Prestashop, Laravel, Codeigniter, Symfony or any other server-side technology.

## Installing
##### *Go >= 1.16*

To start using IsItHacked, install Go and run `go get`:

```sh
$ go get -u github.com/fabelx/isithacked
```

Or install from source using `git`:
```shell script
$ git clone https://github.com/fabelx/isithacked
$ cd isithacked 
$ go build ./cmd/isithacked
```

## Usage

```go
// Retrieve data from service
data, err := isithacked.IsItHacked("example.com")

// Error handling
if err != nil {
    // do something ...
}

// Print received result
d, _ := json.Marshal(data)
fmt.Println(string(d))
```
###Cli

```shell script
# specify the target (domain), the result, if any, 
# will be written to the output.json file 
$ isithacked -target example.com

# specify the target (domain) and result file name
$ isithacked -target example.com -output results.json

# specify the target (ip_v4 address)
$ isithacked -target 127.0.0.1 -ip
```

[godoc-image]: https://godoc.org/github.com/akyoto/cache?status.svg
[godoc-url]: https://godoc.org/github.com/akyoto/cache
