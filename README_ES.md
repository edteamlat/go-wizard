# π§ Go Wizard

> Este README.md tambiΓ©n estΓ‘ disponible en [Ingles](README.md).

Es una aplicaciΓ³n de consola que nos provee una serie de comandos para crear tu proyecto con la Arquitectura Hexagonal.
Nos permitirΓ‘ generar la estructura inicial de un proyecto y tambiΓ©n generar cada capa de un nuevo paquete.

## π― Funcionalidades

1. Crear la estructura inicial de un proyecto β
2. Crear un paquete en las capas requeridas β
3. Agregar un campo en las capas deseadas (database, model, storage, etc.) WIP π 
4. Remover un campo en las capas deseadas (database, model, storage, etc.) WIP π 
5. Sobreescribir un paquete WIP π 

## β Installation

Este comando va a instalar la aplicaciΓ³n en la ruta `$GOPATH/bin`, asi que hay que asegurarse que esa ruta estΓ© en
nuestra variable de entorno `$PATH`

```bash
go install github.com/edteamlat/go-wizard@latest
```

Si deseas utilizar una versiΓ³n diferente, solo reemplaza el `@latest` por la versiΓ³n, ej `@v1.0.0`

## π» Guia de Uso

### `init` command

Para crear la estructura inicial de un nuevo proyecto, solo necesitas ejecutar el siguiente comando:

```bash
go-wizard init -m github.com/edteamlat/my-app
```

Esto va a crear el proyecto en el directorio actual donde se estΓ© ejecutando el comando con las capas para empezar a
trabajar con la arquitectura hexagonal, y automΓ‘ticamente va a inicializar git y los mΓ³dulos de go.

```bash
.
βββ my-app
    βββ cmd
    β   βββ certificates
    β   βββ logs
    β   βββ config.go
    β   βββ configuration.json.example
    β   βββ database.go
    β   βββ echo.go
    β   βββ logger.go
    β   βββ main.go
    β   βββ remoteconfig.go
    βββ domain
    βββ infrastructure
    β   βββ handler
    β   β   βββ request
    β   β   β   βββ fields.go
    β   β   β   βββ parameter.go
    β   β   β   βββ token.go
    β   β   βββ response
    β   β   β   βββ message.go
    β   β   β   βββ response.go
    β   β   βββ router.go
    β   βββ postgres
    βββ model
    β   βββ config.go
    β   βββ error.go
    β   βββ logger.go
    β   βββ messagehandler.go
    β   βββ model.go
    β   βββ model_test.go
    β   βββ remoteconfig.go
    β   βββ router.go
    βββ sqlmigration
    βββ go.mod
    βββ README.md
    βββ wizard-config.yaml
```

Una vez la instalaciΓ³n estΓ‘ terminada, puedes abrir tu proyecto:

```bash
cd my-app
```

Dentro del proyecto creado, debes ejecutar el siguiente comando para instalar las dependencias:

```bash
go mod tidy
```

Si quieres obtener mΓ‘s informaciΓ³n de este comando, ejecuta:

```bash
go-wizard help init
```

### `add package` command

Con este comando podrΓ‘s crear un CRUD en las capas disponibles que se leerΓ‘n de un archivo de configuraciΓ³n yaml que es
generado cuando ejecutamos el comando `init`With this command you can create the CRUD on the available layers that will
read from a config yaml file that is generated when you run the init command:

```yaml
# if you don't specify this field, the wizard will automatically use the path of working directory (pwd)
# Si no indicas este campo, el wizard automΓ‘ticamente va a usar el directorio actual (pwd)
project_path: /home/username/Documents/code/

# El nombre del mΓ³dulo es usado para crear los imports
module_name: github.com/edteamlat/go-wizzard

# Es usado para el nombre de las estructuras de las diferentes capas
model: UserRole

# Va a ser utilizado para el nombre de la tabla y constrains
# tambiΓ©n, el nombre se va a convertir a UpperCamelCase para ser usado
# en el nombre del Slice
table: user_roles

# Esta description serΓ‘ usada como comentario de la nueva tabla
table_comment: Write your comment here

# Una lista de objetos donde se debera indicar el nombre, tipo y si permite nulos
# por defecto, los campos id, created_at y updated_at sarΓ‘n agregados
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
# Si no deseas usar uno, solo remuΓ©vela de la lista
layers:
  - domain
  - handler_echo # por ahora solo soportamos el framework `echo`
  - storage_postgres # por ahora solo soportamos postgres
  - model
  - sqlmigration_postgres # por ahora solo soportamos
```

Este archivo de configuraciΓ³n es agregado en la raΓ­z de tu proyecto con el nombre `wizard-config.yaml` cuando ejecutas
el comando `init`.

Para crear un nuevo paquete solo llena el archivo de configuraciΓ³n y ejecuta el comando:

```bash
go-wizard add package
```

Con el archivo de configuration de arriba, este comando va a generar lo siguiente:

```bash
βββ cmd
βββ domain
β   βββ userrole
β       βββ usecase.go
β       βββ userrole.go
βββ infrastructure
β   βββ handler
β   β   βββ request
β   β   βββ response
β   β   βββ userrole
β   β   β   βββ handler.go
β   β   β   βββ route.go
β   βββ postgres
β       βββ userrole
β           βββ userrole.go
βββ model
β   βββ userrole.go
βββ sqlmigration
β   βββ 20210908_063849_create_user_roles_table.sql
```

Ahora puedes acceder a cada uno de los archivos generados para ver que hay dentro.

Si quieres saber mΓ‘s sobre este comando, ejecuta:
```bash
go-wizard help add
```
