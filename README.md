# Gin-Postgres-API

Gin API for CRUD operations in PostgreSQL.

## 🛠 Built With

<div align="left">
<a href="https://go.dev/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/go.svg" width="36" height="36" alt="Go" /></a>
<a href="https://gin-gonic.com/docs/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/gin.svg" width="36" height="36" alt="Gin" /></a>
<a href="https://www.postgresql.org/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/postgresql.svg" width="36" height="36" alt="PostgreSQL" /></a>
<a href="https://docs.docker.com/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/docker.svg" width="36" height="36" alt="Docker" /></a>
</div>

## ⚙️ Run Locally

Clone the project

```bash
git clone https://github.com/DEMYSTIF/gin-postgres-api.git
cd gin-postgres-api
```

Start the database

```bash
docker compose up -d
```

View the database (optional)

```bash
docker exec -it gin-postgres psql -d gin-postgres -U demystif -W
```

Start the application

```bash
go run .
```

For live reload, install Air (optional)

```bash
go install github.com/cosmtrek/air@latest
```

Run the application with Air

```bash
air
```
