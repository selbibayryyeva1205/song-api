package verses

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ VersesModel = (*customVersesModel)(nil)

type (
	// VersesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVersesModel.
	VersesModel interface {
		versesModel
	}

	customVersesModel struct {
		*defaultVersesModel
	}
)

// NewVersesModel returns a model for the database table.
func NewVersesModel(conn sqlx.SqlConn) VersesModel {
	return &customVersesModel{
		defaultVersesModel: newVersesModel(conn),
	}
}
