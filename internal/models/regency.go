package models

import (
	"time"

	"github.com/sev-2/raiden/pkg/db"
)

type Regency struct {
	db.ModelBase
	Id         int64      `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
	ProvinceId *int64     `json:"province_id,omitempty" column:"name:province_id;type:bigint;nullable"`
	Name       *string    `json:"name,omitempty" column:"name:name;type:varchar;nullable"`
	CreatedAt  time.Time  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable:false;default:now()"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" column:"name:deleted_at;type:timestampz;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"regency" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`
}
