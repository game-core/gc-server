// Package masterHealth ヘルスチェック
package masterHealth

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/health/masterHealth"
)

type masterHealthDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterHealthDao(conn *database.MysqlHandler) masterHealth.MasterHealthMysqlRepository {
	return &masterHealthDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterHealthDao) Find(ctx context.Context, healthId int64) (*masterHealth.MasterHealth, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_health", "Find", fmt.Sprintf("%d_", healthId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterHealth.MasterHealth); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterHealth()
	res := s.ReadMysqlConn.WithContext(ctx).Where("health_id = ?", healthId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterHealth.SetMasterHealth(t.HealthId, t.Name, t.MasterHealthType)
	s.Cache.Set(cashes.CreateCacheKey("master_health", "Find", fmt.Sprintf("%d_", healthId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterHealthDao) FindOrNil(ctx context.Context, healthId int64) (*masterHealth.MasterHealth, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_health", "FindOrNil", fmt.Sprintf("%d_", healthId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterHealth.MasterHealth); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterHealth()
	res := s.ReadMysqlConn.WithContext(ctx).Where("health_id = ?", healthId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterHealth.SetMasterHealth(t.HealthId, t.Name, t.MasterHealthType)
	s.Cache.Set(cashes.CreateCacheKey("master_health", "FindOrNil", fmt.Sprintf("%d_", healthId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterHealthDao) FindList(ctx context.Context) (masterHealth.MasterHealths, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_health", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterHealth.MasterHealths); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterHealths()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterHealth.NewMasterHealths()
	for _, t := range ts {
		ms = append(ms, masterHealth.SetMasterHealth(t.HealthId, t.Name, t.MasterHealthType))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_health", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterHealthDao) Create(ctx context.Context, tx *gorm.DB, m *masterHealth.MasterHealth) (*masterHealth.MasterHealth, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterHealth{
		HealthId:         m.HealthId,
		Name:             m.Name,
		MasterHealthType: m.MasterHealthType,
	}
	res := conn.Model(NewMasterHealth()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterHealth.SetMasterHealth(t.HealthId, t.Name, t.MasterHealthType), nil
}

func (s *masterHealthDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterHealth.MasterHealths) (masterHealth.MasterHealths, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterHealths()
	for _, m := range ms {
		t := &MasterHealth{
			HealthId:         m.HealthId,
			Name:             m.Name,
			MasterHealthType: m.MasterHealthType,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterHealth()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterHealthDao) Update(ctx context.Context, tx *gorm.DB, m *masterHealth.MasterHealth) (*masterHealth.MasterHealth, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterHealth{
		HealthId:         m.HealthId,
		Name:             m.Name,
		MasterHealthType: m.MasterHealthType,
	}
	res := conn.Model(NewMasterHealth()).WithContext(ctx).Where("health_id = ?", m.HealthId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterHealth.SetMasterHealth(t.HealthId, t.Name, t.MasterHealthType), nil
}

func (s *masterHealthDao) Delete(ctx context.Context, tx *gorm.DB, m *masterHealth.MasterHealth) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterHealth()).WithContext(ctx).Where("health_id = ?", m.HealthId).Delete(NewMasterHealth())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
