package roles

import (
	"github.com/sev-2/raiden"
)

type SuperAdmin struct {
	raiden.RoleBase
}

func (r *SuperAdmin) Name() string {
	return "super_admin"
}

