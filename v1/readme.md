# 🧪 guarapo_lab_test

**guarapo_lab_test** es una API REST simple para la gestión de tareas de usuario, construida con Go (Golang). El proyecto está diseñado para ser fácil de ejecutar tanto de forma local con Go como usando Docker Compose.

---

## 🚀 Tecnologías utilizadas

- [Go (Golang)](https://golang.org/)
- [Gin Gonic](https://github.com/gin-gonic/gin) - Framework HTTP
- [GORM](https://gorm.io/) - ORM para Go
- [SQLite](https://www.sqlite.org/index.html) - Base de datos ligera

---

## ⚙️ Requisitos

- Go 1.20 o superior (si lo ejecutas localmente)
- Docker y Docker Compose (si prefieres usar contenedores)

---

## ▶️ Cómo ejecutar el proyecto

### 🔧 Opción 1: Ejecutar con Go

El proyecto puede ejecutarse con persistencia en memoria o en base de datos sqlite, para seleccionar una
de estas dos opciones modifique la variable PERSISTENCE_TYPE en el archivo env_local.sh

El servicio cuenta con endpoinds protejidos por https como no protejidos http, para activar los endpoinds protegidos debe generar las claves ssl haciendo uso del scrip ssl-generator.sh
```bash
cd ./dev_helper/ && ./ssl-generator.sh && cd ..
```
Esto creara las claves en el directorio certs, tenga en cuenta que debe agregar permisos de ejecucion al script. Si desea correr el servicio sin enpoinds https comente las variables PATH_CERT_HTTPS y PATH_KEY_HTTPS en el archivo dev_helper/env_local.sh

```bash
git clone https://github.com/tu-usuario/guarapo_lab_test.git
cd guarapo_lab_test
go mod tidy
source dev_helper/env_local.sh
go run cmd/server/main.go
```


---

### 🔧 Opción 2: Ejecutar con docker

---
Primero contruya la imagen ejecutando:
```bash
cd Docker && ./build.sh
```
Posteriormente ejecute
```bash
cd docker-compose && docker-compose up
```
puede modificar tambien el valor de la variable PERSISTENCE_TYPE en el archivo docker-compose.yml

