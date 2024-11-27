package song

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SongsModel = (*customSongsModel)(nil)

type (
	// SongsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSongsModel.
	SongsModel interface {
		songsModel
	}

	customSongsModel struct {
		*defaultSongsModel
	}
)

// NewSongsModel returns a model for the database table.
func NewSongsModel(conn sqlx.SqlConn) SongsModel {
	return &customSongsModel{
		defaultSongsModel: newSongsModel(conn),
	}
}
