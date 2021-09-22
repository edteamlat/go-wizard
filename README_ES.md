# 🧙 Go Wizard

> Este README.md también está disponible en [Ingles](README.md).

Es una aplicación de consola que nos provee una serie de comandos para crear tu proyecto con la Arquitectura Hexagonal.
Nos permitirá generar la estructura inicial de un proyecto y también generar cada capa de un nuevo paquete.

## 🎯 Funcionalidades

1. Crear la estructura inicial de un proyecto ✅
2. Crear un paquete en las capas requeridas ✅
3. Agregar un campo en las capas deseadas (database, model, storage, etc.) WIP 🛠
4. Remover un campo en las capas deseadas (database, model, storage, etc.) WIP 🛠
5. Sobreescribir un paquete WIP 🛠

## ⚙ Installation

Este comando va a instalar la aplicación en la ruta `$GOPATH/bin`, asi que hay que asegurarse que esa ruta esté en
nuestra variable de entorno `$PATH`

```bash
go install github.com/edteamlat/go-wizard@latest
```

Si deseas utilizar una versión diferente, solo reemplaza el `@latest` por la versión, ej `@v1.0.0`

## 💻 Guia de Uso

### `init` command

Para crear la estructura inicial de un nuevo proyecto, solo necesitas ejecutar el siguiente comando:

```bash
go-wizard init -m github.com/edteamlat/my-app
```

Esto va a crear el proyecto en el directorio actual donde se esté ejecutando el comando con las capas para empezar a
trabajar con la arquitectura hexagonal, y automáticamente va a inicializar git y los módulos de go.

```bash
my-app
├── cmd
│   ├── config.go
│   ├── configuration.json
│   ├── database.go
│   ├── echo.go
│   ├── logger.go
│   ├── main.go
│   └── remoteconfig.go
├── domain
├── go.mod
├── infrastructure
│   ├── handler
│   │   ├── request
│   │   │   ├── fields.go
│   │   │   ├── parameter.go
│   │   │   └── token.go
│   │   ├── response
│   │   │   ├── message.go
│   │   │   └── response.go
│   │   └── router.go
│   └── postgres
├── model
│   ├── config.go
│   ├── error.go
│   ├── filter.go
│   ├── logger.go
│   ├── messagehandler.go
│   ├── model.go
│   ├── model_test.go
│   ├── remoteconfig.go
│   └── router.go
├── README.md
├── sqlmigration
└── wizard-config.yaml
```

Una vez la instalación está terminada, puedes abrir tu proyecto:

```bash
cd my-app
```

Dentro del proyecto creado, debes ejecutar el siguiente comando para instalar las dependencias:

```bash
go mod init
```

Si quieres obtener más información de este comando, ejecuta:

```bash
go-wizard help init
```

### `add package` command

Con este comando podrás crear un CRUD en las capas disponibles que se leerán de un archivo de configuración yaml que es
generado cuando ejecutamos el comando `init`With this command you can create the CRUD on the available layers that will
read from a config yaml file that is generated when you run the init command:

```yaml
# if you don't specify this field, the wizard will automatically use the path of working directory (pwd)
# Si no indicas este campo, el wizard automáticamente va a usar el directorio actual (pwd)
project_path: /home/username/Documents/code/

# El nombre del módulo es usado para crear los imports
module_name: github.com/edteamlat/go-wizzard

# Es usado para el nombre de las estructuras de las diferentes capas
model: UserRole

# Va a ser utilizado para el nombre de la tabla y constrains
# también, el nombre se va a convertir a UpperCamelCase para ser usado
# en el nombre del Slice
table: user_roles

# Esta description será usada como comentario de la nueva tabla
table_comment: Write your comment here

# Una lista de objetos donde se debera indicar el nombre, tipo y si permite nulos
# por defecto, los campos id, created_at y updated_at sarán agregados
fields:
  - name: user_id
    type: uint
    is_null: false
  - name: description
    type: string
    is_null: true
  - name: is_active
    type: bool
    is_null: false
  - name: begins_at
    type: time.Time
    is_null: false

# Las capas disponibles que se pueden generar.
# Si no deseas usar uno, solo remuévela de la lista
layers:
  - domain
  - handler_echo # por ahora solo soportamos el framework `echo`
  - storage_postgres # por ahora solo soportamos postgres
  - model
  - sqlmigration_postgres # por ahora solo soportamos
```

Este archivo de configuración es agregado en la raíz de tu proyecto con el nombre `wizard-config.yaml` cuando ejecutas
el comando `init`.

Para crear un nuevo paquete solo llena el archivo de configuración y ejecuta el comando:

```bash
go-wizard add package
```

Con el archivo de configuration de arriba, este comando va a generar lo siguiente:

```bash
├── cmd
├── domain
│   └── userrole
│       ├── usecase.go
│       └── userrole.go
├── infrastructure
│   ├── handler
│   │   ├── request
│   │   ├── response
│   │   ├── userrole
│   │   │   ├── handler.go
│   │   │   └── route.go
│   └── postgres
│       └── userrole
│           └── userrole.go
├── model
│   └── userrole.go
├── sqlmigration
│   └── 20210908_063849_create_user_roles_table.sql
```

Ahora puedes acceder a cada uno de los archivos generados para ver que hay dentro.

Si quieres saber más sobre este comando, ejecuta:
```bash
go-wizard help add
```
