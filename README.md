# nulltheory.com

A simple app written with [Go](http://golang.org/) to serve up static content from ``public/``

## Running

Apart from having a working Go installation there's nothing to install, just run the following:

	go run ./nulltheory.go -p 5000

The more adventurous can compile the app first:

	go build ./nulltheory.go
	./nulltheory -p 5000

Then browse to [http://0.0.0.0.0:5000](http://0.0.0.0.0:5000)