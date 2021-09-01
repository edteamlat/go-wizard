package texttemplate

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"

	"github.com/edteamlat/go-wizard/domain/stringparser"
	"github.com/edteamlat/go-wizard/model"
)

const edhexTemplatesPath = "../../cmd/templates/edhex"

const moduleName = "github.com/edteamlat/go-wizard"

type fields struct {
	tpl *template.Template
}

type args struct {
	templateName string
	data         model.Layer
}

type testTable struct {
	name    string
	fields  fields
	args    args
	wantWr  string
	wantErr bool
}

type testTables []testTable

func TestTemplate_Create(t1 *testing.T) {
	tests := testTables{}
	tests = append(tests, getDomainLayerTests()...)
	tests = append(tests, getDomainUseCaseLayerTests()...)
	tests = append(tests, getModelLayerTests()...)
	tests = append(tests, getSQLMigrationLayerTests()...)
	tests = append(tests, getHandlerLayerTests()...)
	tests = append(tests, getHandlerRouteLayerTests()...)
	tests = append(tests, getHandlerLayerTests()...)

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := textTemplate{
				tpl: tt.fields.tpl,
			}
			wr := &bytes.Buffer{}
			err := t.Create(wr, tt.args.templateName, tt.args.data)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWr := wr.String(); gotWr != tt.wantWr {
				t1.Errorf("Create() gotWr = %v, want %v", gotWr, tt.wantWr)
			}
		})
	}
}

func getDomainLayerTests() testTables {
	path := fmt.Sprintf("%s/domain/package/domain.gotpl", edhexTemplatesPath)
	tpl := template.Must(template.New("domain.gotpl").Funcs(stringparser.GetTemplateFunctions()).ParseFiles(path))

	return testTables{
		{
			name: "domain one word package name",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "domain.gotpl",
				data: model.Layer{
					Model:      "User",
					Table:      "users",
					Fields:     nil,
					ModuleName: moduleName,
				},
			},
			wantWr: `package user

import (
	"github.com/edteamlat/go-wizard/model"
)

type UseCase interface {
	Create(m *model.User) error
	Update(m *model.User) error
	Delete(ID uint) error

	GetWhere(specification model.FiltersSpecification) (model.User, error)
	GetAllWhere(specification model.FiltersSpecification) (model.Users, error)
}

type Storage interface {
	GetTx() (model.Transaction, error)

	Create(m *model.User) error
	Update(m *model.User) error
	Delete(ID uint) error

	GetWhere(specification model.FiltersSpecification) (model.User, error)
	GetAllWhere(specification model.FiltersSpecification) (model.Users, error)
}
`,
			wantErr: false,
		},
		{
			name: "domain two words package name",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "domain.gotpl",
				data: model.Layer{
					Model:      "UserLogin",
					Table:      "user_logins",
					Fields:     nil,
					ModuleName: moduleName,
				},
			},
			wantWr: `package userlogin

import (
	"github.com/edteamlat/go-wizard/model"
)

type UseCase interface {
	Create(m *model.UserLogin) error
	Update(m *model.UserLogin) error
	Delete(ID uint) error

	GetWhere(specification model.FiltersSpecification) (model.UserLogin, error)
	GetAllWhere(specification model.FiltersSpecification) (model.UserLogins, error)
}

type Storage interface {
	GetTx() (model.Transaction, error)

	Create(m *model.UserLogin) error
	Update(m *model.UserLogin) error
	Delete(ID uint) error

	GetWhere(specification model.FiltersSpecification) (model.UserLogin, error)
	GetAllWhere(specification model.FiltersSpecification) (model.UserLogins, error)
}
`,
			wantErr: false,
		},
		{
			name: "domain two words package name",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "domain.gotpl",
				data: model.Layer{
					Model:      "UserRole",
					Table:      "user_roles",
					Fields:     nil,
					ModuleName: moduleName,
				},
			},
			wantWr: `package userrole

import (
	"github.com/edteamlat/go-wizard/model"
)

type UseCase interface {
	Create(m *model.UserRole) error
	Update(m *model.UserRole) error
	Delete(ID uint) error

	GetWhere(specification model.FiltersSpecification) (model.UserRole, error)
	GetAllWhere(specification model.FiltersSpecification) (model.UserRoles, error)
}

type Storage interface {
	GetTx() (model.Transaction, error)

	Create(m *model.UserRole) error
	Update(m *model.UserRole) error
	Delete(ID uint) error

	GetWhere(specification model.FiltersSpecification) (model.UserRole, error)
	GetAllWhere(specification model.FiltersSpecification) (model.UserRoles, error)
}
`,
			wantErr: false,
		},
		{
			name: "domain one word package name",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "domain.gotpl",
				data: model.Layer{
					Model:      "Role",
					Table:      "roles",
					Fields:     nil,
					ModuleName: moduleName,
				},
			},
			wantWr: `package role

import (
	"github.com/edteamlat/go-wizard/model"
)

type UseCase interface {
	Create(m *model.Role) error
	Update(m *model.Role) error
	Delete(ID uint) error

	GetWhere(specification model.FiltersSpecification) (model.Role, error)
	GetAllWhere(specification model.FiltersSpecification) (model.Roles, error)
}

type Storage interface {
	GetTx() (model.Transaction, error)

	Create(m *model.Role) error
	Update(m *model.Role) error
	Delete(ID uint) error

	GetWhere(specification model.FiltersSpecification) (model.Role, error)
	GetAllWhere(specification model.FiltersSpecification) (model.Roles, error)
}
`,
			wantErr: false,
		},
	}
}

func getDomainUseCaseLayerTests() testTables {
	path := fmt.Sprintf("%s/domain/package/usecase.gotpl", edhexTemplatesPath)
	tpl := template.Must(template.New("usecase.gotpl").Funcs(stringparser.GetTemplateFunctions()).ParseFiles(path))

	return testTables{
		{
			name: "",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "usecase.gotpl",
				data: model.Layer{
					Model: "User",
					Table: "users",
					Fields: []model.Field{
						{
							Name: "id",
						},
						{
							Name: "created_at",
						},
					},
					ModuleName: moduleName,
				},
			},
			wantWr: `package user

import (
	"fmt"

	"github.com/edteamlat/go-wizard/model"
)

var allowedFieldsForQuery = []string{
	"id","created_at",
}

// User implements UseCase
type User struct {
	storage Storage
}

// New returns a new User
func New(s Storage) User {
	return User{storage: s}
}

// Create creates a model.User
func (u User) Create(m *model.User) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("user: %w", model.ErrNilPointer)
	}

	if err := m.Validate(); err != nil {
		return fmt.Errorf("user: %w", err)
	}

	err := u.storage.Create(m)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// Update updates a model.User by id
func (u User) Update(m *model.User) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("user: %w", model.ErrNilPointer)
	}

	if !m.HasID() {
		return fmt.Errorf("user: %w", model.ErrInvalidID)
	}

	if err := m.Validate(); err != nil {
		return fmt.Errorf("user: %w", err)
	}

	err := u.storage.Update(m)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// Delete deletes a model.User by id
func (u User) Delete(ID uint) error {
	err := u.storage.Delete(ID)
	if err != nil {
		return handleStorageErr(err)
	}

	return nil
}

// GetWhere returns a model.User according to filters and sorts
func (u User) GetWhere(specification model.FiltersSpecification) (model.User, error) {
	if err := specification.Fields.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	user, err := u.storage.GetWhere(specification)
	if err != nil {
		return model.User{}, fmt.Errorf("user: %w", err)
	}

	return user, nil
}

// GetAllWhere returns a model.Users according to filters and sorts
func (u User) GetAllWhere(specification model.FiltersSpecification) (model.Users, error) {
	if err := specification.Fields.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	if err := specification.Sorts.ValidateNames(allowedFieldsForQuery); err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	users, err := u.storage.GetAllWhere(specification)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	return users, nil
}

// handleStorageErr handles errors from storage layer
func handleStorageErr(err error) error {
	e := model.NewError()
	e.SetError(err)

	switch err {
	default:
		return err
	}
}
`,
			wantErr: false,
		},
	}
}

func getModelLayerTests() testTables {
	path := fmt.Sprintf("%s/model/newmodel.gotpl", edhexTemplatesPath)
	tpl := template.Must(template.New("newmodel.gotpl").Funcs(stringparser.GetTemplateFunctions()).ParseFiles(path))

	return testTables{
		{
			name: "model one word package name",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "newmodel.gotpl",
				data: model.Layer{
					Model: "User",
					Table: "users",
					Fields: model.Fields{
						{
							Name: "id",
							Type: "uint",
						},
						{
							Name: "first_name",
							Type: "string",
						},
						{
							Name: "email",
							Type: "string",
						},
						{
							Name: "is_active",
							Type: "bool",
						},
						{
							Name: "created_at",
							Type: "time.Time",
						},
					},
				},
			},
			wantWr: fmt.Sprintf(`package model

import "time"

// User model of table users
type User struct {
	ID uint %[1]sjson:"id"%[1]s
	FirstName string %[1]sjson:"first_name"%[1]s
	Email string %[1]sjson:"email"%[1]s
	IsActive bool %[1]sjson:"is_active"%[1]s
	CreatedAt time.Time %[1]sjson:"created_at"%[1]s
	}

func (u User) HasID() bool { return u.ID > 0 }

func (u User) Validate() error {
	// implement validation of fields for creation and update
	return nil
}

// Users slice of User
type Users []User

func (u Users) IsEmpty() bool { return len(Users) == 0 }
`, "`"),
			wantErr: false,
		},
		{
			name: "model two words package name",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "newmodel.gotpl",
				data: model.Layer{
					Model: "UserRole",
					Table: "user_roles",
					Fields: model.Fields{
						{
							Name: "id",
							Type: "uint",
						},
						{
							Name: "role_id",
							Type: "uint16",
						},
						{
							Name: "user_id",
							Type: "int",
						},
						{
							Name: "is_active",
							Type: "bool",
						},
						{
							Name: "created_at",
							Type: "time.Time",
						},
					},
				},
			},
			wantWr: fmt.Sprintf(`package model

import "time"

// UserRole model of table user_roles
type UserRole struct {
	ID uint %[1]sjson:"id"%[1]s
	RoleID uint16 %[1]sjson:"role_id"%[1]s
	UserID int %[1]sjson:"user_id"%[1]s
	IsActive bool %[1]sjson:"is_active"%[1]s
	CreatedAt time.Time %[1]sjson:"created_at"%[1]s
	}

func (u UserRole) HasID() bool { return u.ID > 0 }

func (u UserRole) Validate() error {
	// implement validation of fields for creation and update
	return nil
}

// UserRoles slice of UserRole
type UserRoles []UserRole

func (u UserRoles) IsEmpty() bool { return len(UserRoles) == 0 }
`, "`"),
			wantErr: false,
		},
	}
}

func getSQLMigrationLayerTests() testTables {
	path := fmt.Sprintf("%s/sqlmigration/sqlmigration.gotpl", edhexTemplatesPath)
	tpl := template.Must(template.New("sqlmigration.gotpl").Funcs(stringparser.GetTemplateFunctions()).ParseFiles(path))

	return testTables{
		{
			name: "",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "sqlmigration.gotpl",
				data: model.Layer{
					Model:        "Course",
					Table:        "courses",
					TableComment: "Write your comment",
					Fields: model.Fields{
						{
							Name:   "id",
							Type:   "uint",
							IsNull: false,
						},
						{
							Name:   "title",
							Type:   "string",
							IsNull: false,
						},
						{
							Name:   "is_premium",
							Type:   "bool",
							IsNull: false,
						},
						{
							Name:   "infographics",
							Type:   "json.RawMessage",
							IsNull: true,
						},
						{
							Name:   "created_at",
							Type:   "time.Time",
							IsNull: false,
						},
						{
							Name:   "updated_at",
							Type:   "time.Time",
							IsNull: true,
						},
					},
				},
			},
			wantWr: `CREATE TABLE courses (
	id SERIAL NOT NULL,
	title VARCHAR(SIZE) NOT NULL,
	is_premium BOOLEAN NOT NULL,
	infographics JSON,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP,
	CONSTRAINT courses_id_pk PRIMARY KEY (id)
);

COMMENT ON TABLE courses IS 'Write your comment';

-- Register the permission module for the routes
INSERT INTO modules (name) VALUES ('COURSE');
`,
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "sqlmigration.gotpl",
				data: model.Layer{
					Model:        "CoursePrice",
					Table:        "course_prices",
					TableComment: "Write your comment",
					Fields: model.Fields{
						{
							Name:   "id",
							Type:   "uint",
							IsNull: false,
						},
						{
							Name:   "course_id",
							Type:   "uint64",
							IsNull: false,
						},
						{
							Name:   "price",
							Type:   "float32",
							IsNull: false,
						},
						{
							Name:   "base_price",
							Type:   "float64",
							IsNull: true,
						},
						{
							Name:   "is_active",
							Type:   "bool",
							IsNull: false,
						},
						{
							Name:   "begins_at",
							Type:   "time.Time",
							IsNull: false,
						},
						{
							Name:   "ends_at",
							Type:   "time.Time",
							IsNull: true,
						},
					},
				},
			},
			wantWr: `CREATE TABLE course_prices (
	id SERIAL NOT NULL,
	course_id INTEGER NOT NULL,
	price NUMERIC(SIZE) NOT NULL,
	base_price NUMERIC(SIZE),
	is_active BOOLEAN NOT NULL,
	begins_at TIMESTAMP NOT NULL,
	ends_at TIMESTAMP,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP,
	CONSTRAINT course_prices_id_pk PRIMARY KEY (id)
);

COMMENT ON TABLE course_prices IS 'Write your comment';

-- Register the permission module for the routes
INSERT INTO modules (name) VALUES ('COURSE_PRICE');
`,
			wantErr: false,
		},
	}
}

func getHandlerRouteLayerTests() testTables {
	path := fmt.Sprintf("%s/infrastructure/handler/package/route.gotpl", edhexTemplatesPath)
	tpl := template.Must(template.New("route.gotpl").Funcs(stringparser.GetTemplateFunctions()).ParseFiles(path))

	return testTables{
		{
			name: moduleName,
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "route.gotpl",
				data: model.Layer{
					Model:      "Invoice",
					Table:      "invoices",
					ModuleName: moduleName,
				},
			},
			wantWr: fmt.Sprintf(`package invoice

import (
	"database/sql"

	"%[1]s/model"
	"%[1]s/infrastructure/handler/response"
	"%[1]s/domain/%[2]s"
	%[2]sStorage "%[1]s/infrastructure/postgres/%[2]s"
)

// NewRouter returns a router to handle model.Invoice requests
func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	// build middlewares to validate permissions on the routes

	adminRoutes(specification, handler)
	privateRoutes(specification, handler)
	publicRoutes(specification, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	responser := response.New(specification.Logger)

	useCase := %[2]s.New(%[2]sStorage.New(specification.DB))
	return newHandler(useCase, responser)
}

// adminRoutes handle the routes that requires a token and permissions to certain users
func adminRoutes(api *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := api.Group("api/v1/admin/invoices", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}

// privateRoutes handle the routes that requires a token
func privateRoutes(api *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := api.Group("/api/v1/private/invoices", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(api *echo.Echo, h handler) {
	route := api.Group("/api/v1/invoices")

	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}
`, moduleName, "invoice"),
			wantErr: false,
		},
		{
			name: moduleName,
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "route.gotpl",
				data: model.Layer{
					Model:      "InvoiceItem",
					Table:      "invoice_items",
					ModuleName: moduleName,
				},
			},
			wantWr: fmt.Sprintf(`package invoiceitem

import (
	"database/sql"

	"%[1]s/model"
	"%[1]s/infrastructure/handler/response"
	"%[1]s/domain/%[2]s"
	%[2]sStorage "%[1]s/infrastructure/postgres/%[2]s"
)

// NewRouter returns a router to handle model.InvoiceItem requests
func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	// build middlewares to validate permissions on the routes

	adminRoutes(specification, handler)
	privateRoutes(specification, handler)
	publicRoutes(specification, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	responser := response.New(specification.Logger)

	useCase := %[2]s.New(%[2]sStorage.New(specification.DB))
	return newHandler(useCase, responser)
}

// adminRoutes handle the routes that requires a token and permissions to certain users
func adminRoutes(api *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := api.Group("api/v1/admin/invoice-items", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}

// privateRoutes handle the routes that requires a token
func privateRoutes(api *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := api.Group("/api/v1/private/invoice-items", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}

// publicRoutes handle the routes that not requires a validation of any kind to be use
func publicRoutes(api *echo.Echo, h handler) {
	route := api.Group("/api/v1/invoice-items")

	route.GET("", h.GetAllWhere)
	route.GET("/:id", h.GetWhere)
}
`, moduleName, "invoiceitem"),
			wantErr: false,
		},
	}
}

func getHandlerLayerTests() testTables {
	path := fmt.Sprintf("%s/infrastructure/handler/package/handler.gotpl", edhexTemplatesPath)
	tpl := template.Must(template.New("handler.gotpl").Funcs(stringparser.GetTemplateFunctions()).ParseFiles(path))

	return testTables{
		{
			name: moduleName,
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "handler.gotpl",
				data: model.Layer{
					Model:      "Order",
					Table:      "orders",
					ModuleName: moduleName,
				},
			},
			wantWr: fmt.Sprintf(`package %[2]s

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strings"

	"%[1]s/%[2]s"
	"%[1]s/infrastructure/handler/request"
	"%[1]s/infrastructure/handler/response"
	"%[1]s/model"
)

type handler struct {
	useCase invoice.useCase
	response response.Responser
}

func newHandler(useCase %[2]s.useCase) handler {
	return handler{useCase: useCase}
}

// Create handles the creation of a model.%[3]s
func (h handler) Create(c echo.Context) error {
	m := model.%[3]s{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(c, err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}

// Update handles the update of a model.%[3]s
func (h handler) Update(c echo.Context) error {
	m := model.%[3]s{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(c, err)
	}

	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		return h.response.BindFailed(c, err)
	}
	m.ID = uint(ID)

	if err := h.useCase.Update(&m); err != nil {
		return h.response.Error(c, "useCase.Update()", err)
	}

	return c.JSON(h.response.Updated(m))
}

// Delete handles the deleting of a model.%[3]s
func (h handler) Delete(c echo.Context) error {
	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		return h.response.BindFailed(c, err)
	}

	err = h.useCase.Delete(uint(ID))
	if err != nil {
		return h.response.Error(c, "useCase.Delete()", err)
	}

	return c.JSON(h.response.Deleted(nil))
}

// GetWhere handles the search of a model.%[3]s
func (h handler) GetWhere(c echo.Context) error {
	userID := request.GetUserID(c)

	filtersSpecification, err := request.GetFiltersSpecification(c)
	if err != nil {
		return err
	}

	orderData, err := h.useCase.GetAllWhere(filtersSpecification)
	if err != nil {
		return h.response.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.response.OK(orderData))
}

// GetAllWhere handles the search of all model.%[3]s
func (h handler) GetAllWhere(c echo.Context) error {
	userID := request.GetUserID(c)

	filtersSpecification, err := request.GetFiltersSpecification(c)
	if err != nil {
		return err
	}

	orders, err := h.useCase.GetAllWhere(filtersSpecification)
	if err != nil {
		return h.response.Error(c, "useCase.GetAllWhere()", err)
	}

	return c.JSON(h.response.OK(orders))
}
`, moduleName, "order", "Order"),
			wantErr: false,
		},
		{
			name: moduleName,
			fields: fields{
				tpl: tpl,
			},
			args: args{
				templateName: "handler.gotpl",
				data: model.Layer{
					Model:      "OrderItem",
					Table:      "OrderItems",
					ModuleName: moduleName,
				},
			},
			wantWr: fmt.Sprintf(`package %[2]s

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strings"

	"%[1]s/%[2]s"
	"%[1]s/infrastructure/handler/request"
	"%[1]s/infrastructure/handler/response"
	"%[1]s/model"
)

type handler struct {
	useCase invoice.useCase
	response response.Responser
}

func newHandler(useCase %[2]s.useCase) handler {
	return handler{useCase: useCase}
}

// Create handles the creation of a model.%[3]s
func (h handler) Create(c echo.Context) error {
	m := model.%[3]s{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(c, err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}

// Update handles the update of a model.%[3]s
func (h handler) Update(c echo.Context) error {
	m := model.%[3]s{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(c, err)
	}

	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		return h.response.BindFailed(c, err)
	}
	m.ID = uint(ID)

	if err := h.useCase.Update(&m); err != nil {
		return h.response.Error(c, "useCase.Update()", err)
	}

	return c.JSON(h.response.Updated(m))
}

// Delete handles the deleting of a model.%[3]s
func (h handler) Delete(c echo.Context) error {
	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		return h.response.BindFailed(c, err)
	}

	err = h.useCase.Delete(uint(ID))
	if err != nil {
		return h.response.Error(c, "useCase.Delete()", err)
	}

	return c.JSON(h.response.Deleted(nil))
}

// GetWhere handles the search of a model.%[3]s
func (h handler) GetWhere(c echo.Context) error {
	userID := request.GetUserID(c)

	filtersSpecification, err := request.GetFiltersSpecification(c)
	if err != nil {
		return err
	}

	orderItemData, err := h.useCase.GetAllWhere(filtersSpecification)
	if err != nil {
		return h.response.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.response.OK(orderItemData))
}

// GetAllWhere handles the search of all model.%[3]s
func (h handler) GetAllWhere(c echo.Context) error {
	userID := request.GetUserID(c)

	filtersSpecification, err := request.GetFiltersSpecification(c)
	if err != nil {
		return err
	}

	orderItems, err := h.useCase.GetAllWhere(filtersSpecification)
	if err != nil {
		return h.response.Error(c, "useCase.GetAllWhere()", err)
	}

	return c.JSON(h.response.OK(orderItems))
}
`, moduleName, "orderitem", "OrderItem"),
			wantErr: false,
		},
	}

}
