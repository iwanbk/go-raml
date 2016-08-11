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

## Server

Generate server code by using this command
```
$ go-raml server --ramlfile ../api.raml --import-path examples.com/goramldir --dir $GOPATH/src/examples.com/goramldir
```

You can find all server files in `$GOPATH/src/examples.com/goramldir` directory.


### Server side itsyou.online integration

We need to write/modify some code for this integration

**JWT decoder**

You can find JWT decoder code in [iyo.go](server/iyo.go) and do these:

- modify the package from `iyo` to `main`
- copy `iyo.go` to `$GOPATH/src/examples.com/goramldir` directory

**Oauth2 middleware**

Need to modify ` oauth2_itsyouonline_middleware.go` in `$GOPATH/src/examples.com/goramldir` directory.

Find this line in `func (om *Oauth2itsyouonlineMiddleware) Handler(next http.Handler) http.Handler {`
```
var scopes []string
```
replace it with
```
scopes, err := getIyoUserScope(accessToken)
if err != nil {
    w.WriteHeader(403)
    return
}
```

In `getIyoUserScope` function, you can find code to:

- decode itsyou.online JWT token which require itsyou.online public key
- check if the token issued by itsyou.online
- get the `scopes`

## Client

**itsyou.online client library**

We need the library to get oauth2 token and generate JWT token

`go get -u -v github.com/itsyouonline/identityserver/clients/go/itsyouonline`
