package verses

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	versesFieldNames          = builder.RawFieldNames(&Verses{}, true)
	versesRows                = strings.Join(versesFieldNames, ",")
	versesRowsExpectAutoSet   = strings.Join(stringx.Remove(versesFieldNames, "id"), ",")
	versesRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(versesFieldNames, "id"))
)

type (
	versesModel interface {
		Insert(ctx context.Context, data *Verses) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Verses, error)
		Update(ctx context.Context, data *Verses) error
		Delete(ctx context.Context, id int64) error
	}

	defaultVersesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Verses struct {
		Id          int64  `db:"id"`
		SongId      int    `db:"song_id"`
		VerseNumber int    `db:"verse_number"`
		Text        string `db:"song_text"`
	}
)

func newVersesModel(conn sqlx.SqlConn) *defaultVersesModel {
	return &defaultVersesModel{
		conn:  conn,
		table: `"public"."verses"`,
	}
}

func (m *defaultVersesModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultVersesModel) FindOne(ctx context.Context, id int64) (*Verses, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", versesRows, m.table)
	var resp Verses
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVersesModel) Insert(ctx context.Context, data *Verses) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3)", m.table, versesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.SongId, data.VerseNumber, data.Text)
	return ret, err
}

func (m *defaultVersesModel) Update(ctx context.Context, data *Verses) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, versesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Id, data.SongId, data.VerseNumber, data.Text)
	return err
}

func (m *defaultVersesModel) tableName() string {
	return m.table
}
