// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LoadGroupsColumns holds the columns for the "load_groups" table.
	LoadGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "name", Type: field.TypeString},
		{Name: "remark", Type: field.TypeString, Default: ""},
	}
	// LoadGroupsTable holds the schema information for the "load_groups" table.
	LoadGroupsTable = &schema.Table{
		Name:       "load_groups",
		Columns:    LoadGroupsColumns,
		PrimaryKey: []*schema.Column{LoadGroupsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "group_name",
				Unique:  false,
				Columns: []*schema.Column{LoadGroupsColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LoadGroupsTable,
	}
)

func init() {
	LoadGroupsTable.Annotation = &entsql.Annotation{
		Table: "load_groups",
	}
}
