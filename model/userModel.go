package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		TransInsert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *defaultUserModel) TransInsert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	testStudyUserIdKey := fmt.Sprintf("%s%v", cacheTestStudyUserIdPrefix, data.Id)
	testStudyUserMobileKey := fmt.Sprintf("%s%v", cacheTestStudyUserMobilePrefix, data.Mobile)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.Nickname, data.Mobile)
	}, testStudyUserIdKey, testStudyUserMobileKey)
	return ret, err
}

// TransCtx 暴露给Logic开启事务
func (m *defaultUserModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		return fn(ctx, s)
	})
}
