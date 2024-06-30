// Package userAccount ユーザーアカウント
package userAccount

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/pkg/domain/model/account/userAccount"
)

type userAccountRedisDao struct {
	ReadRedisConn  *redis.Client
	WriteRedisConn *redis.Client
}

func NewUserAccountRedisDao(conn *database.RedisHandler) userAccount.UserAccountRedisRepository {
	return &userAccountRedisDao{
		ReadRedisConn:  conn.User.ReadRedisConn,
		WriteRedisConn: conn.User.WriteRedisConn,
	}
}

func (s *userAccountRedisDao) Find(ctx context.Context, userId string) (*userAccount.UserAccount, error) {
	t := NewUserAccount()
	data, err := s.ReadRedisConn.HGet(ctx, t.TableName(), fmt.Sprintf("%s:userId:%v", t.TableName(), userId)).Result()
	if err != nil {
		return nil, err
	}

	if err := t.JsonToTable(data); err != nil {
		return nil, err
	}

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountRedisDao) FindOrNil(ctx context.Context, userId string) (*userAccount.UserAccount, error) {
	t := NewUserAccount()
	data, err := s.ReadRedisConn.HGet(ctx, t.TableName(), fmt.Sprintf("%s:userId:%v", t.TableName(), userId)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	if err := t.JsonToTable(data); err != nil {
		return nil, err
	}

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountRedisDao) Set(ctx context.Context, tx redis.Pipeliner, m *userAccount.UserAccount) (*userAccount.UserAccount, error) {
	var conn redis.Pipeliner
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteRedisConn.TxPipeline()
	}

	t := &UserAccount{
		UserId:   m.UserId,
		Name:     m.Name,
		Password: m.Password,
		LoginAt:  m.LoginAt,
		LogoutAt: m.LogoutAt,
	}

	jt, err := t.TableToJson()
	if err != nil {
		return nil, err
	}

	if err := conn.HSet(ctx, t.TableName(), fmt.Sprintf("%s:userId:%v", t.TableName(), m.UserId), jt).Err(); err != nil {
		return nil, err
	}

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountRedisDao) Delete(ctx context.Context, tx redis.Pipeliner, m *userAccount.UserAccount) error {
	var conn redis.Pipeliner
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteRedisConn.TxPipeline()
	}

	t := NewUserAccount()
	if err := conn.HDel(ctx, t.TableName(), fmt.Sprintf("%s:userId:%v", t.TableName(), m.UserId)).Err(); err != nil {
		return err
	}

	return nil
}
