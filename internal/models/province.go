package models

import (
	"time"

	"github.com/sev-2/raiden/pkg/db"
)

type Province struct {
	db.ModelBase
	Id        int64     `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	Name      string    `json:"name,omitempty" column:"name:name;type:text;nullable:false"`
	CreatedAt time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestamp with time zone;default:now();nullable:false"`
	UpdatedAt time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestamp with time zone;default:now();nullable:false"`
	DeletedAt time.Time `json:"deleted_at,omitempty" column:"name:deleted_at;type:timestamp with time zone;default:now();nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"province" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
