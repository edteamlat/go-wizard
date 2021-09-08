# 🧙 Go Wizard

Is a command line interface that provides you with useful commands to help you build your application with the Hexagonal
Architecture. It allows you to generate the initial structure of a project and to generate every layer of a new package.

## 🎯 Features

1. Create the initial structure of a project ✅
2. Create a package on the different layers ✅
3. Add one field on the desired layers (database, model, storage, etc.) WIP 🛠
4. Remove one field on the desired layers (database, model, storage, etc.) WIP 🛠
5. Override a package WIP 🛠

## ⚙ Installation

This command will install the CLI on $GOPATH/bin so make sure that you have that on your $PATH.

```bash
go install github.com/edteamlat/go-wizard@latest
```

If you want a different version, just replace the `@latest` with the version `@v1.0.0`

## 💻 Usage

### Init command

To create the initial structure of a new project you just need to run the next command:

```bash
go-wizard init -m github.com/edteamlat/my-app
```

This will create the project on the current directory with the layers to work with the hexagonal architecture, and
it'll automatically init git and the go modules.

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
Once the installation is done, you can open your project folder:
```bash
cd my-app
```

Inside the created project, you must run the next command to install the dependencies:
```bash
go mod init
```

If you want to get more information about this command, run:
```bash
go-wizard help init
```

### Add package command
With this command you can create the CRUD on the available layers that will read from a config yaml file:
```yaml
# if you don't specify this field, the wizard will automatically use the path of working directory (pwd)
project_path: /home/username/Documents/code/

# the module_name is used to create the imports
module_name: github.com/edteamlat/go-wizzard

# Use for the name of the structs of the different layers
model: UserRole

# Will be use for the table name and constraints
# also, it'll be converted to camel case to be use
# for the Slice model
table: user_roles

# table_comment will be use as the comment for the new table
table_comment: Write your comment here

# a list of objects where you specify the name, type and if it allows nulls
# by default, the id, created_at and updated_at fields, will be added
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

# The available layers that we can generate
# if you don't want to use one, just remove it
layers:
  - domain
  - handler_echo # we only support the echo framework for now
  - storage_postgres # we don't support other db system for now
  - model
  - sqlmigration_postgres # here we'll save the sql files to modify our db, it only supports postgres syntax for now
```

This config file is added on the root of your project  with the name `wizard-config.yaml` when you exec the init command.

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


If you want to get more information about this command, run:
```bash
go-wizard help add
```
