// Package masterShard シャード
package masterShard

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/shard/masterShard"
)

type masterShardDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterShardDao(conn *database.MysqlHandler) masterShard.MasterShardMysqlRepository {
	return &masterShardDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterShardDao) Find(ctx context.Context, masterShardId int64) (*masterShard.MasterShard, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_shard", "Find", fmt.Sprintf("%d_", masterShardId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterShard.MasterShard); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterShard()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_shard_id = ?", masterShardId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterShard.SetMasterShard(t.MasterShardId, t.Name, t.ShardKey, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_shard", "Find", fmt.Sprintf("%d_", masterShardId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterShardDao) FindOrNil(ctx context.Context, masterShardId int64) (*masterShard.MasterShard, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_shard", "FindOrNil", fmt.Sprintf("%d_", masterShardId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterShard.MasterShard); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterShard()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_shard_id = ?", masterShardId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterShard.SetMasterShard(t.MasterShardId, t.Name, t.ShardKey, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_shard", "FindOrNil", fmt.Sprintf("%d_", masterShardId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterShardDao) FindByShardKey(ctx context.Context, shardKey string) (*masterShard.MasterShard, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_shard", "FindByShardKey", fmt.Sprintf("%s_", shardKey)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterShard.MasterShard); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterShard()
	res := s.ReadMysqlConn.WithContext(ctx).Where("shard_key = ?", shardKey).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterShard.SetMasterShard(t.MasterShardId, t.Name, t.ShardKey, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_shard", "FindByShardKey", fmt.Sprintf("%s_", shardKey)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterShardDao) FindOrNilByShardKey(ctx context.Context, shardKey string) (*masterShard.MasterShard, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_shard", "FindOrNilByShardKey", fmt.Sprintf("%s_", shardKey)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterShard.MasterShard); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterShard()
	res := s.ReadMysqlConn.WithContext(ctx).Where("shard_key = ?", shardKey).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterShard.SetMasterShard(t.MasterShardId, t.Name, t.ShardKey, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_shard", "FindOrNilByShardKey", fmt.Sprintf("%s_", shardKey)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterShardDao) FindList(ctx context.Context) (masterShard.MasterShards, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_shard", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterShard.MasterShards); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterShards()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterShard.NewMasterShards()
	for _, t := range ts {
		ms = append(ms, masterShard.SetMasterShard(t.MasterShardId, t.Name, t.ShardKey, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_shard", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterShardDao) FindListByShardKey(ctx context.Context, shardKey string) (masterShard.MasterShards, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_shard", "FindListByShardKey", fmt.Sprintf("%s_", shardKey)))
	if found {
		if cachedEntity, ok := cachedResult.(masterShard.MasterShards); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterShards()
	res := s.ReadMysqlConn.WithContext(ctx).Where("shard_key = ?", shardKey).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterShard.NewMasterShards()
	for _, t := range ts {
		ms = append(ms, masterShard.SetMasterShard(t.MasterShardId, t.Name, t.ShardKey, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_shard", "FindListByShardKey", fmt.Sprintf("%s_", shardKey)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterShardDao) Create(ctx context.Context, tx *gorm.DB, m *masterShard.MasterShard) (*masterShard.MasterShard, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterShard{
		MasterShardId: m.MasterShardId,
		Name:          m.Name,
		ShardKey:      m.ShardKey,
		Count:         m.Count,
	}
	res := conn.Model(NewMasterShard()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterShard.SetMasterShard(t.MasterShardId, t.Name, t.ShardKey, t.Count), nil
}

func (s *masterShardDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterShard.MasterShards) (masterShard.MasterShards, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterShards()
	for _, m := range ms {
		t := &MasterShard{
			MasterShardId: m.MasterShardId,
			Name:          m.Name,
			ShardKey:      m.ShardKey,
			Count:         m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterShard()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterShardDao) Update(ctx context.Context, tx *gorm.DB, m *masterShard.MasterShard) (*masterShard.MasterShard, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterShard{
		MasterShardId: m.MasterShardId,
		Name:          m.Name,
		ShardKey:      m.ShardKey,
		Count:         m.Count,
	}
	res := conn.Model(NewMasterShard()).WithContext(ctx).Where("master_shard_id = ?", m.MasterShardId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterShard.SetMasterShard(t.MasterShardId, t.Name, t.ShardKey, t.Count), nil
}

func (s *masterShardDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterShard.MasterShards) (masterShard.MasterShards, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterShards()
	for _, m := range ms {
		t := &MasterShard{
			MasterShardId: m.MasterShardId,
			Name:          m.Name,
			ShardKey:      m.ShardKey,
			Count:         m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterShard()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_shard_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "shard_key", "count"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterShardDao) Delete(ctx context.Context, tx *gorm.DB, m *masterShard.MasterShard) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterShard()).WithContext(ctx).Where("master_shard_id = ?", m.MasterShardId).Delete(NewMasterShard())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterShardDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterShard.MasterShards) error {
	if len(ms) <= 0 {
		return nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	var ks [][]interface{}
	for _, m := range ms {
		ks = append(ks, []interface{}{m.MasterShardId})
	}

	res := conn.Model(NewMasterShard()).WithContext(ctx).Where("(master_shard_id) IN ?", ks).Delete(NewMasterShard())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
