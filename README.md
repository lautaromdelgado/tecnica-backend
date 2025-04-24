# Hetmo Technical Test - MVP Backend

## ✨ Descripción General

Este proyecto es un **MVP (Producto Mínimo Viable)** desarrollado como parte de una **prueba técnica para la empresa Hetmo**, con el objetivo de simular la gestión de eventos (creación, inscripción, administración de usuarios y logs del sistema).

La aplicación está desarrollada con **Golang** utilizando el potente y minimalista framework web **Echo**. Su arquitectura está completamente basada en el patrón de diseño **Clean Architecture**, garantizando un alto grado de desacoplamiento, escalabilidad y mantenibilidad.

---

## 🏒 Objetivo del Proyecto

El sistema busca resolver los problemas comunes de administración de eventos, tales como:

- Gestión de usuarios y roles (admin/usuario)
- Inscripción y cancelación a eventos
- Visualización filtrada de eventos según el estado (activo/completado)
- Gestión de eventos solo por parte de administradores
- Registro automático de logs/historial por cada acción

Todo esto implementado de manera segura, robusta y profesional.

---

## 📈 Tecnologías Utilizadas

- **Golang** 1.21+
- **Echo Framework** (web framework)
- **SQLX** (driver SQL robusto con soporte para structs y queries seguras)
- **MySQL 8** como sistema de base de datos
- **Docker & Docker Compose** para facilitar el despliegue
- **Postman** para pruebas de API y documentación

---

## ✨ Características Clave

- Autenticación mediante **JWT** (JSON Web Tokens)
- Protección de rutas según roles:
  - Usuarios con rol `admin` tienen acceso completo
  - Usuarios con rol `user` acceden solo a rutas permitidas
- **Soft Delete** implementado para usuarios y eventos
- **Historial (Logs)** automático ante cada acción realizada sobre los eventos
- Búsqueda y filtros avanzados de eventos y logs por título, acción, organizador y estado
- **Dockerfile** + **docker-compose.yml** para levantar el entorno completo con un solo comando

---

## 🧰 Arquitectura del Proyecto

El proyecto sigue al 100% el patrón **Clean Architecture**, dividiendo el sistema en capas bien definidas:

```
/internal
|-- domain: entidades y contratos (interfaces)
|-- usecase: lógica de negocio (casos de uso)
|-- delivery: controladores HTTP, middlewares y rutas
|-- infrastructure: acceso a datos, implementaciones reales
```

Esta arquitectura facilita:

- Reemplazar MySQL por otro motor sin tocar la lógica de negocio
- Cambiar Echo por otro framework sin tocar los casos de uso
- Escalar el sistema por funcionalidades sin impactar el código existente

---

## 🚀 Requisitos

- **Golang** 1.21 o superior (solo si vas a compilar localmente)
- **Docker** y **Docker Compose** instalados

---

## ⚙️ Instalación & Ejecución

1. Clona el repositorio:

```bash
git clone https://github.com/tuusuario/tecnica-backend-hetmo.git
cd tecnica-backend-hetmo
```

2. Ejecutá el entorno completo con Docker:

```bash
docker-compose up --build
```

Esto levantará:
- Contenedor de MySQL con migraciones y datos de prueba
- Backend Golang escuchando en `http://localhost:8080`

---

## 🔧 Variables de Entorno

**.env.docker:**
```dotenv
JWT_SECRET=hetmo_app_secret_key
PORT=8080
DB_HOST=db
DB_PORT=3306
DB_USER=root
DB_PASS=root
DB_NAME=hetmo_app
```

---

## 🔍 Colección Postman

Podés testear todos los endpoints desde la siguiente colección oficial:

[Acceder a la Colección de Postman 🔗](https://web.postman.co/workspace/fa624846-2b1b-428c-82eb-2c2f2a248207/documentation/36816741-38cdcef7-c0de-4625-8e9c-4a49c314d8f1)

---

## 🌟 Funcionalidades destacadas

- [x] Registro y login de usuarios
- [x] Protección de rutas con JWT y middleware de roles
- [x] CRUD completo de eventos (solo admin)
- [x] Inscripción a eventos solo si están publicados y con fecha futura
- [x] Visualización de eventos completados y activos
- [x] Cancelación (desuscripción) de eventos por parte del usuario
- [x] Logs detallados de cada acción en la base de datos:
  - Creación de eventos
  - Actualizaciones
  - Eliminaciones
  - Restauraciones
  - Publicación / Despublicación
- [x] Acceso a historial por parte de administradores (visualización de logs filtrados)
- [x] Filtros avanzados por título, organizador y acción
- [x] Total desacoplamiento según Clean Architecture

---

## 📅 Estado del Proyecto

✅ Listo para pruebas — Puede extenderse para:
- Enviar emails de confirmación
- Panel administrativo frontend
- Exportación de logs en Excel/PDF

---

## 🌟 Autor

Desarrollado por **Lautaro M. Delgado** como parte de una prueba técnica para Hetmo.

---

> "No solo es código, es arquitectura, escalabilidad y orden."

