// Package userAccountToken ユーザーアカウントトークン
package userAccountToken

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/pkg/domain/model/account/userAccountToken"
)

type userAccountTokenRedisDao struct {
	ReadRedisConn  *redis.Client
	WriteRedisConn *redis.Client
}

func NewUserAccountTokenRedisDao(conn *database.RedisHandler) userAccountToken.UserAccountTokenRedisRepository {
	return &userAccountTokenRedisDao{
		ReadRedisConn:  conn.User.ReadRedisConn,
		WriteRedisConn: conn.User.WriteRedisConn,
	}
}

func (s *userAccountTokenRedisDao) Find(ctx context.Context, userId string) (*userAccountToken.UserAccountToken, error) {
	t := NewUserAccountToken()
	data, err := s.ReadRedisConn.HGet(ctx, t.TableName(), fmt.Sprintf("%s:userId:%v", t.TableName(), userId)).Result()
	if err != nil {
		return nil, err
	}

	if err := t.JsonToTable(data); err != nil {
		return nil, err
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenRedisDao) FindOrNil(ctx context.Context, userId string) (*userAccountToken.UserAccountToken, error) {
	t := NewUserAccountToken()
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

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenRedisDao) Set(ctx context.Context, tx redis.Pipeliner, m *userAccountToken.UserAccountToken) (*userAccountToken.UserAccountToken, error) {
	var conn redis.Pipeliner
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteRedisConn.TxPipeline()
	}

	t := &UserAccountToken{
		UserId: m.UserId,
		Token:  m.Token,
	}

	jt, err := t.TableToJson()
	if err != nil {
		return nil, err
	}

	if err := conn.HSet(ctx, t.TableName(), fmt.Sprintf("%s:userId:%v", t.TableName(), m.UserId), jt).Err(); err != nil {
		return nil, err
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenRedisDao) Delete(ctx context.Context, tx redis.Pipeliner, m *userAccountToken.UserAccountToken) error {
	var conn redis.Pipeliner
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteRedisConn.TxPipeline()
	}

	t := NewUserAccountToken()
	if err := conn.HDel(ctx, t.TableName(), fmt.Sprintf("%s:userId:%v", t.TableName(), m.UserId)).Err(); err != nil {
		return err
	}

	return nil
}
