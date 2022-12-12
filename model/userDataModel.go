package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserDataModel = (*customUserDataModel)(nil)

type (
	// UserDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserDataModel.
	UserDataModel interface {
		userDataModel
		TransInsert(ctx context.Context, session sqlx.Session, data *UserData) (sql.Result, error)
	}

	customUserDataModel struct {
		*defaultUserDataModel
	}
)

// NewUserDataModel returns a model for the database table.
func NewUserDataModel(conn sqlx.SqlConn, c cache.CacheConf) UserDataModel {
	return &customUserDataModel{
		defaultUserDataModel: newUserDataModel(conn, c),
	}
}

func (m *defaultUserDataModel) TransInsert(ctx context.Context, session sqlx.Session, data *UserData) (sql.Result, error) {
	testStudyUserDataIdKey := fmt.Sprintf("%s%v", cacheTestStudyUserDataIdPrefix, data.Id)
	testStudyUserDataUserIdKey := fmt.Sprintf("%s%v", cacheTestStudyUserDataUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userDataRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.Data)
	}, testStudyUserDataIdKey, testStudyUserDataUserIdKey)
	return ret, err
}
