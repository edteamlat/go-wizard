package texttemplate

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"

	"github.com/edteamlat/go-wizard/domain/stringparser"
	"github.com/edteamlat/go-wizard/model"
)

const edhexTemplatesPath = "../../templates/edhex"

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
	path := fmt.Sprintf("%s/domain/domain.gotpl", edhexTemplatesPath)
	tpl := template.Must(template.New("domain.gotpl").Funcs(stringparser.GetTemplateFunctions()).ParseFiles(path))

	return testTables{
		{
			name: "",
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
			name: "",
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
			name: "",
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
			name: "",
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
