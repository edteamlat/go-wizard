# 🧙 Go Wizard

> This README.md is also available in [Spanish](README_ES.md).

Is a command line interface that provides you with useful commands to help you build your application with the Hexagonal
Architecture. It allows you to generate the initial structure of a project and to generate every layer of a new package.

## 🎯 Features

1. Create the initial structure of a project ✅
2. Create a package on the different layers ✅
3. Add one field on the desired layers (database, model, storage, etc.) WIP 🛠
4. Remove one field on the desired layers (database, model, storage, etc.) WIP 🛠
5. Override a package WIP 🛠

## ⚙ Installation

This command will install the CLI on `$GOPATH/bin` so make sure that you have that on your `$PATH`.

```bash
go install github.com/edteamlat/go-wizard@latest
```

If you want a different version, just replace the `@latest` with the version `@v1.0.0`

## 💻 Usage

### `init` command

To create the initial structure of a new project you just need to run the next command:

```bash
go-wizard init -m github.com/edteamlat/my-app
```

This will create the project on the current directory with the layers to work with the hexagonal architecture, and
it'll automatically init git and the go modules.

```bash
.
└── my-app
    ├── cmd
    │   ├── certificates
    │   ├── logs
    │   ├── config.go
    │   ├── configuration.json.example
    │   ├── database.go
    │   ├── echo.go
    │   ├── logger.go
    │   ├── main.go
    │   └── remoteconfig.go
    ├── domain
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
    │   ├── logger.go
    │   ├── messagehandler.go
    │   ├── model.go
    │   ├── model_test.go
    │   ├── remoteconfig.go
    │   └── router.go
    ├── sqlmigration
    ├── go.mod
    ├── README.md
    └── wizard-config.yaml
```

Once the installation is done, you can open your project folder:

```bash
cd my-app
```

After we init the project the `go mod tidy` command will be executed to download the dependencies.

If you want to get more information about this command, run:

```bash
go-wizard help init
```

### `add package` command

With this command you can create the CRUD on the available layers that will read from a config yaml file that is
generated when you run the init command:

```yaml
# if you don't specify this field, the wizard will automatically use the path of working directory (pwd)
# here you'll have to change the username and set the correct path to your project
project_path: /home/username/Documents/code/

# the module_name is used to create the imports
module_name: github.com/edteamlat/go-wizzard

# Is Use for the name of the structs of the different layers
model: UserRole

# Will be use for the table name and constraints
# also, it'll be converted to UpperCamelCase to be use
# for the Slice model
table: user_roles

# table_comment will be use as the comment for the new table
table_comment: Write your comment here

# choose the type of the dates (valid options: time.Time and int64 for unix)
time_type: time.Time

# choose the type of the primary key (valid options: uint, uuid.UUID)
id_type: uint

# a list of objects where you specify the name, type and if it allows nulls
# by default, the id, created_at and updated_at fields, will be added
fields:
  - name: user_id
    type: uint
    is_null: false
  - name: description
    type: string
    is_null: true
    field_size: 255 # only for string type, if not specified, it'll be 255, if you put -1, it'll be of type TEXT instead of VARCHAR
  - name: is_active
    type: bool
    is_null: false
  - name: begins_at
    type: time.Time
    is_null: false

# The available layers that we can generate
# if you don't want to use one, just remove it
layers:
  - domain
  - handler_echo # we only support the echo framework for now
  - storage_postgres # we don't support other db system for now
  - model
  - sqlmigration_postgres # here we'll save the sql files to modify our db, it only supports postgres syntax for now
```

This config file is added on the root of your project with the name `wizard-config.yaml` when you exec the init command.

To create a new package just fill the config file and exec the next command:

```bash
go-wizard add package
```

With the config file of above this command will generate the following:

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

Now you can enter into every file that was generated to see what's inside.

You can also pass the flag `-c` to indicate the path of the config file:

```bash
go-wizard add package -c /home/username/Documents/code/my-app/wizard-config.yaml
```

If you want to get more information about this command, run:

```bash
go-wizard help add
```

## 🏷 Available types

| Go Type         | Postgres type                  |
|-----------------|--------------------------------|
| string          | VARCHAR(SIZE) and TEXT         |
| int             | INTEGER                        |
| int8            | INTEGER                        |
| int16           | INTEGER                        |
| int32           | INTEGER                        |
| int64           | INTEGER                        |
| uint            | INTEGER                        |
| uint8           | INTEGER                        |
| uint16          | INTEGER                        |
| uint32          | INTEGER                        |
| uint64          | INTEGER                        |
| float32         | NUMERIC()                      |
| float64         | NUMERIC()                      |
| bool            | BOOLEAN                        |
| time.Time       | TIMESTAMP and INTEGER for unix |
| uuid.UUID       | UUID                           |
| json.RawMessage | JSON                           |

If other types are used, the wizard will return set INVALID_TYPE as the type of the field, so if you want to use other
types, you'll have to open an issue or make a PR.

