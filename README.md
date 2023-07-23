# test-boletia

## Requisitos

* Go 1.20.4
* PostgreSQL 15.x ( docker pull postgres:latest )
## Ambiente Local ( BD basado en Docker )

```
docker run --name test-boletia-db -e POSTGRES_PASSWORD=123456 -d -p 5432:5432 postgres:latest
```

## Creación Esquema y Tablas - Carga datos iniciales ( Basado en Docker)

Copiar scripts dentro del contenedor : 
```
docker cp ./db/scripts/ test-boletia-db:/tmp/
```
Eliminación y creación de esquema y tablas :
```
docker exec -it test-boletia-db psql -U postgres -d test-boletia-db -f /tmp/scripts/create-db.sql
```


## Compilación y Ejecución

```
make run
make test
make lint
make build
make clean
make all
make             # default is make all
```

This has been created using go modules; to run the tests, just execute:

```bash
go test -mod vendor -race -cover -coverprofile=coverage.txt -covermode=atomic ./...
```

or (using make):

```bash
make test
```

The Makefile also supports other commands, such as:

```bash
make lint
```

## Docker

Comandos para generación de contenedor de API. No es necesario para ambiente local.
```bash
docker build -t test-boletia:1.0 .
docker run -p 1323:1323 --name test-boletia test-boletia:1.0
```