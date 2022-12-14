// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userDataFieldNames          = builder.RawFieldNames(&UserData{})
	userDataRows                = strings.Join(userDataFieldNames, ",")
	userDataRowsExpectAutoSet   = strings.Join(stringx.Remove(userDataFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userDataRowsWithPlaceHolder = strings.Join(stringx.Remove(userDataFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheTestStudyUserDataIdPrefix     = "cache:testStudy:userData:id:"
	cacheTestStudyUserDataUserIdPrefix = "cache:testStudy:userData:userId:"
)

type (
	userDataModel interface {
		Insert(ctx context.Context, data *UserData) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserData, error)
		FindOneByUserId(ctx context.Context, userId int64) (*UserData, error)
		Update(ctx context.Context, data *UserData) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserDataModel struct {
		sqlc.CachedConn
		table string
	}

	UserData struct {
		Id         int64     `db:"id"`
		UserId     int64     `db:"user_id"` // user.id
		Data       string    `db:"data"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newUserDataModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserDataModel {
	return &defaultUserDataModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_data`",
	}
}

func (m *defaultUserDataModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	testStudyUserDataIdKey := fmt.Sprintf("%s%v", cacheTestStudyUserDataIdPrefix, id)
	testStudyUserDataUserIdKey := fmt.Sprintf("%s%v", cacheTestStudyUserDataUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, testStudyUserDataIdKey, testStudyUserDataUserIdKey)
	return err
}

func (m *defaultUserDataModel) FindOne(ctx context.Context, id int64) (*UserData, error) {
	testStudyUserDataIdKey := fmt.Sprintf("%s%v", cacheTestStudyUserDataIdPrefix, id)
	var resp UserData
	err := m.QueryRowCtx(ctx, &resp, testStudyUserDataIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userDataRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserDataModel) FindOneByUserId(ctx context.Context, userId int64) (*UserData, error) {
	testStudyUserDataUserIdKey := fmt.Sprintf("%s%v", cacheTestStudyUserDataUserIdPrefix, userId)
	var resp UserData
	err := m.QueryRowIndexCtx(ctx, &resp, testStudyUserDataUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userDataRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserDataModel) Insert(ctx context.Context, data *UserData) (sql.Result, error) {
	testStudyUserDataIdKey := fmt.Sprintf("%s%v", cacheTestStudyUserDataIdPrefix, data.Id)
	testStudyUserDataUserIdKey := fmt.Sprintf("%s%v", cacheTestStudyUserDataUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userDataRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Data)
	}, testStudyUserDataIdKey, testStudyUserDataUserIdKey)
	return ret, err
}

func (m *defaultUserDataModel) Update(ctx context.Context, newData *UserData) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	testStudyUserDataIdKey := fmt.Sprintf("%s%v", cacheTestStudyUserDataIdPrefix, data.Id)
	testStudyUserDataUserIdKey := fmt.Sprintf("%s%v", cacheTestStudyUserDataUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userDataRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.Data, newData.Id)
	}, testStudyUserDataIdKey, testStudyUserDataUserIdKey)
	return err
}

func (m *defaultUserDataModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTestStudyUserDataIdPrefix, primary)
}

func (m *defaultUserDataModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userDataRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserDataModel) tableName() string {
	return m.table
}
