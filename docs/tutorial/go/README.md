# Go Tutorial

In this tutorial we will generate a simple Go client & server from [an RAML file](../api.raml) and then integrate it
with [itsyou.online](https://www.itsyou.online/) authorization server.

This tutorial use [jumpscale/ubuntu1604_golang](https://hub.docker.com/r/jumpscale/ubuntu1604_golang/) docker image.
But you can use any system with Go installed and configured.

## Create docker instances

You can skip this step if you want to use your own enviroment.

```
sudo docker pull jumpscale/ubuntu1604_golang
sudo docker run --rm -t -i  --name=goraml jumpscale/ubuntu1604_golang
```

## Generate server

```
$ go-raml server --ramlfile ../api.raml --import-path examples.com/goramldir --dir $GOPATH/src/examples.com/goramldir
```

Then you can find all server files in `$GOPATH/src/examples.com/goramldir` directory.

## Install itsyouonline Go client library

We will use [itsyouonline](https://www.itsyou.online/) as Oauth2 Authorization server.

`go get -u -v github.com/itsyouonline/identityserver/clients/go/itsyouonline`

