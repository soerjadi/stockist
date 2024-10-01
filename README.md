# STOCKIST

Simple e-commerce application

## how to run
- copy and adjust config file locate in `files/config-example.ini`
```
$ cp files/config-example.ini files/config.ini
```
- adjust database, redis and please set token
- run go mod
```
$ go mod tidy && go mod vendor
```
- run the migrations. sql schema locate in `script/sql/schema.sql`
- running the application

```
$ cd cmd/rest && go run app.go
```

### how to run in docker
- please adjust database and redis configuration in config.development.ini
- run docker build
```
$ docker build --tag=stockist:1.0.0 .
```
- after docker build finished. run docker image
```
$ docker run -e ENV=PRODUCTION stockist:1.0.0
```