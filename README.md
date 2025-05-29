# Gooportunities API

API REST desenvolvida em **Go (Golang)**, utilizando:

- [Gin](https://github.com/gin-gonic/gin) como framework web;
- [GORM](https://gorm.io/) para ORM e conexão com banco de dados **SQLite**;
- [Swag](https://github.com/swaggo/swag) para geração de documentação Swagger;
- `Makefile` para automatizar tarefas como geração de docs e execução da aplicação.

---

## 🚀 Tecnologias Utilizadas

- **Go 1.21+**
- **Gin**
- **GORM**
- **SQLite**
- **Swagger**
- **Make**

---

## 🛠️ Instalação e Execução

### Pré-requisitos

- Go instalado: https://go.dev/doc/install
- SQLite (opcional, já embutido via GORM)
- `make` instalado

### Rodando a aplicação

```bash
make run
