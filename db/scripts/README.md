# Scripts BD test boletia - PostgreSQL

Para crear o recrear la BD test boletia se debe tener:

* Cliente PostgreSQL (psql)

Los pasos de ejecución son los siguientes

`psql -h <host> -p <port> -U <usuario> -f <archivo_sql>`

Ejemplo:

`psql -h localhost -p 5432 -U postgres -f install.sql`

Este comando debe ejecutar dentro de esta carpeta que contiene los scripts.

A continuación para eliminar la BD:

`psql -h localhost -p 5432 -U postgres -f uninstall.sql`

Y finalmente:

`psql -h localhost -p 5432 -U postgres -f install.sql`