// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// VideoColumns holds the columns for the "video" table.
	VideoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString, Size: 255},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"processing", "ready"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// VideoTable holds the schema information for the "video" table.
	VideoTable = &schema.Table{
		Name:        "video",
		Columns:     VideoColumns,
		PrimaryKey:  []*schema.Column{VideoColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		VideoTable,
	}
)

func init() {
}
