# go make it

A walking skeleton of a go project using docker compose

## building

- You can use the standard Go tools if you wish, see `build-app.sh`
- In Jenkins this will be built using `test.sh` which uses Docker. `run-locally.sh` also uses Docker so that you can use all sorts of databases or whatever inside `docker-compose.yml`

## dependencies! go! wtf! 

Basically your external lib deps live in /vendor

You *could* manage this manually, but this project has [govendor](https://github.com/kardianos/govendor) installed so you should probably use that.
 
## i've never written go before

- [Install it](https://golang.org/dl/)
- [Next, must-read guide](https://golang.org/doc/code.html)
- [Effective Go](https://golang.org/doc/effective_go.html)

From there, there's tons of resources out there. Notable things:

- [The official go blog](https://blog.golang.org/) Lots of in-depth articles
- [Ben Johnson](https://medium.com/@benbjohnson) has an ongoing series of in-depth guides to the standard library. [The post on the io package](https://medium.com/@benbjohnson/go-walkthrough-io-package-8ac5e95a9fbd#.18heiybyt) is lovely and you will end up using it a lot so worth a read.
- [Go by example](https://gobyexample.com/)
 
## some general tips

- Want the documentation of all of your local packages including stdlib? `godoc -http=:6060`