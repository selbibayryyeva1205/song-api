package song

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	songsFieldNames          = builder.RawFieldNames(&Songs{}, true)
	songsRows                = strings.Join(songsFieldNames, ",")
	songsRowsExpectAutoSet   = strings.Join(stringx.Remove(songsFieldNames, "id"), ",")
	songsRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(songsFieldNames, "id"))
)

type (
	songsModel interface {
		Insert(ctx context.Context, data *Songs) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Songs, error)
		Update(ctx context.Context, data *Songs) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSongsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Songs struct {
		//Id          int64          `db:"id"`
		GroupName   string         `db:"group_name"`
		SongName    string         `db:"song_name"`
		ReleaseDate time.Time      `db:"release_date"`
		Link        sql.NullString `db:"link"`
		Text        string         `db:"song_text"`
		// CreatedAt   time.Time      `db:"created_at"`
		// UpdatedAt   time.Time      `db:"updated_at"`
	}
)

func newSongsModel(conn sqlx.SqlConn) *defaultSongsModel {
	return &defaultSongsModel{
		conn:  conn,
		table: `"public"."songs"`,
	}
}

func (m *defaultSongsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSongsModel) FindOne(ctx context.Context, id int64) (*Songs, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", songsRows, m.table)
	var resp Songs
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

func (m *defaultSongsModel) Insert(ctx context.Context, data *Songs) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4,$5)", m.table, songsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.GroupName, data.SongName, data.ReleaseDate, data.Link, data.Text)
	return ret, err
}

func (m *defaultSongsModel) Update(ctx context.Context, data *Songs) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, songsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.GroupName, data.SongName, data.ReleaseDate, data.Link)
	return err
}

func (m *defaultSongsModel) tableName() string {
	return m.table
}
func (m *defaultSongsModel) FindAll(ctx context.Context, groupName, songName *string, page, pageSize int) ([]Songs, error) {
	baseQuery := `
		SELECT id, group_name, song_name, release_date, link, verses, created_at, updated_at
		FROM songs
		WHERE 1=1
	`
	var args []interface{}
	var conditions []string

	// Dynamically add conditions based on input
	if groupName != nil && *groupName != "" {
		conditions = append(conditions, "group_name ILIKE $"+fmt.Sprint(len(args)+1))
		args = append(args, "%"+*groupName+"%")
	}

	if songName != nil && *songName != "" {
		conditions = append(conditions, "song_name ILIKE $"+fmt.Sprint(len(args)+1))
		args = append(args, "%"+*songName+"%")
	}

	// Combine conditions into the query
	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Add ordering, pagination
	baseQuery += " ORDER BY created_at DESC LIMIT $" + fmt.Sprint(len(args)+1) + " OFFSET $" + fmt.Sprint(len(args)+2)
	args = append(args, pageSize, (page-1)*pageSize)

	// Execute query
	var songs []Songs
	err := m.conn.QueryRowsCtx(ctx, &songs, baseQuery, args...)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
