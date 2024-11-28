package song

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
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
		Insert(ctx context.Context, data *Songs) (int64, error)
		FindOne(ctx context.Context, id int64, verseId int) (*Songs, error)
		Update(ctx context.Context, data *SongsUpdate) error
		Delete(ctx context.Context, id int64) error
		FindAll(ctx context.Context, groupName, songName *string, page, pageSize int) (song *SongsResult, err error)
	}

	defaultSongsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Songs struct {
		GroupName   string         `db:"group_name"`
		SongName    string         `db:"song_name"`
		ReleaseDate time.Time      `db:"release_date"`
		Link        sql.NullString `db:"link"`
		Text        string         `db:"song_text"`
	}

	SongsUpdate struct {
		GroupName   string         `db:"group_name"`
		SongName    string         `db:"song_name"`
		ReleaseDate time.Time      `db:"release_date"`
		Link        sql.NullString `db:"link"`
		Text        string         `db:"song_text"`
	}

	SongsResult struct {
		Songs []Songs `json:"songs"`
		Total int     `json:"total"`
	}

	GetOneSongResult struct {
		GroupName   string         `db:"group_name"`
		SongName    string         `db:"song_name"`
		ReleaseDate time.Time      `db:"release_date"`
		Link        sql.NullString `db:"link"`
		Text        string         `db:"song_text"`
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
	logx.WithContext(ctx).Infof("Executing delete query: %s with id: %d", query, id)
	_, err := m.conn.ExecCtx(ctx, query, id)
	if err != nil {
		logx.WithContext(ctx).Errorf("Error executing delete query: %v", err)
	}
	return err
}

func (m *defaultSongsModel) Insert(ctx context.Context, data *Songs) (int64, error) {
	var id int64
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		m.table, songsRowsExpectAutoSet,
	)
	logx.WithContext(ctx).Infof("Executing insert query: %s", query)
	err := m.conn.QueryRowCtx(ctx, &id, query, data.GroupName, data.SongName, data.ReleaseDate, data.Link, data.Text)
	if err != nil {
		logx.WithContext(ctx).Errorf("Error executing insert query: %v", err)
		return id, err
	}
	logx.WithContext(ctx).Debugf("Inserted song with id: %d", id)
	return id, err
}

func (m *defaultSongsModel) Update(ctx context.Context, data *SongsUpdate) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, songsRowsWithPlaceHolder)
	logx.WithContext(ctx).Infof("Executing update query: %s with data: %+v", query, data)
	_, err := m.conn.ExecCtx(ctx, query, data.GroupName, data.SongName, data.ReleaseDate, data.Link)
	if err != nil {
		logx.WithContext(ctx).Errorf("Error executing update query: %v", err)
	}
	return err
}

func (m *defaultSongsModel) tableName() string {
	return m.table
}

func (m *defaultSongsModel) FindAll(ctx context.Context, groupName, songName *string, page, pageSize int) (*SongsResult, error) {
	baseQuery := `
		SELECT s.id,
		       s.song_name, 
		       s.link, 
		       s.release_date, 
		       s.song_text
		FROM songs s
	`
	countQuery := `
		SELECT COUNT(*)
		FROM songs s
	`

	var args []interface{}
	var conditions []string

	if groupName != nil && *groupName != "" {
		conditions = append(conditions, "s.group_name ILIKE $"+fmt.Sprint(len(args)+1))
		args = append(args, "%"+*groupName+"%")
	}

	if songName != nil && *songName != "" {
		conditions = append(conditions, "s.song_name ILIKE $"+fmt.Sprint(len(args)+1))
		args = append(args, "%"+*songName+"%")
	}

	if len(conditions) > 0 {
		conditionString := " WHERE " + strings.Join(conditions, " AND ")
		baseQuery += conditionString
		countQuery += conditionString
	}

	baseQuery += " ORDER BY s.created_at DESC LIMIT $" + fmt.Sprint(len(args)+1) + " OFFSET $" + fmt.Sprint(len(args)+2)
	args = append(args, pageSize, (page-1)*pageSize)

	var totalCount int
	logx.WithContext(ctx).Infof("Executing count query: %s", countQuery)
	err := m.conn.QueryRowCtx(ctx, &totalCount, countQuery, args[:len(args)-2]...)
	if err != nil {
		logx.WithContext(ctx).Errorf("Error executing count query: %v", err)
		return nil, err
	}

	var songs []Songs
	logx.WithContext(ctx).Infof("Executing main query: %s", baseQuery)
	err = m.conn.QueryRowsCtx(ctx, &songs, baseQuery, args...)
	if err != nil {
		logx.WithContext(ctx).Errorf("Error executing main query: %v", err)
		return nil, err
	}

	song := &SongsResult{
		Songs: songs,
		Total: totalCount,
	}
	logx.WithContext(ctx).Debugf("Fetched songs: %+v", song)
	return song, nil
}

func (m *defaultSongsModel) FindOne(ctx context.Context, id int64, verseId int) (*Songs, error) {
	query := `SELECT 
    s.song_name, 
    s.link, 
    s.release_date, 
    v."text"
FROM 
    songs s
LEFT JOIN 
    verses v 
ON 
    v.song_id = s.id where v.song_id= $1 and v.verse_number =$2`

	var resp Songs
	logx.WithContext(ctx).Infof("Executing FindOne query: %s with id: %d, verseId: %d", query, id, verseId)
	err := m.conn.QueryRowCtx(ctx, &resp, query, id, verseId)
	switch err {
	case nil:
		logx.WithContext(ctx).Debugf("Found song: %+v", resp)
		return &resp, nil
	case sqlc.ErrNotFound:
		logx.WithContext(ctx).Errorf("Song not found with id: %d and verseId: %d", id, verseId)
		return nil, ErrNotFound
	default:
		logx.WithContext(ctx).Errorf("Error executing FindOne query: %v", err)
		return nil, err
	}
}
