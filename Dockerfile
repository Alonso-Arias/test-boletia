############################
# STEP 1 build executable binary
############################
FROM golang:1.20-alpine3.17 as builder

CMD mkdir /app
COPY . /app

WORKDIR /app

# Sincronizar el directorio vendor con las dependencias actuales
RUN go mod vendor
RUN go mod tidy

# Build the binary.
RUN CGO_ENABLED=0 go build -mod vendor -ldflags="-s -w" -o ./bin/server ./cmd/server/

############################
# STEP 2 build a small image
############################
FROM alpine:3.14.2

# Crear un usuario no privilegiado llamado "appuser"
RUN adduser -D -g '' appuser

USER appuser

# Crear el directorio de la aplicación
RUN mkdir -p /home/appuser/app

WORKDIR /home/appuser/app

# Copiar el binario desde el builder
COPY --from=builder /app/bin/server /home/appuser/app/bin/server

# Ejecutar la aplicación
CMD ["/home/appuser/app/bin/server"]
