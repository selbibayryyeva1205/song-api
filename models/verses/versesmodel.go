package verses

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

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
	query:=`CREATE TABLE IF NOT EXISTS verses (
    id SERIAL PRIMARY KEY,
    song_id INT REFERENCES songs(id) ON DELETE CASCADE,
    verse_number INT,
    song_text TEXT NOT NULL

);
`
	_, err := conn.Exec(query)
	if err!=nil{
		fmt.Println("ERROR IN CREATE",err)
	}
	return &customVersesModel{
		defaultVersesModel: newVersesModel(conn),
	}
}
