// Package masterHealth ヘルスチェック
package masterHealth

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/health/masterHealth"
)

type masterHealthMysqlDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterHealthMysqlDao(conn *database.MysqlHandler) masterHealth.MasterHealthMysqlRepository {
	return &masterHealthMysqlDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterHealthMysqlDao) Find(ctx context.Context, healthId int64) (*masterHealth.MasterHealth, error) {
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

	m := masterHealth.SetMasterHealth(t.HealthId, t.Name, t.MasterHealthEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_health", "Find", fmt.Sprintf("%d_", healthId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterHealthMysqlDao) FindOrNil(ctx context.Context, healthId int64) (*masterHealth.MasterHealth, error) {
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

	m := masterHealth.SetMasterHealth(t.HealthId, t.Name, t.MasterHealthEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_health", "FindOrNil", fmt.Sprintf("%d_", healthId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterHealthMysqlDao) FindList(ctx context.Context) (masterHealth.MasterHealths, error) {
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
		ms = append(ms, masterHealth.SetMasterHealth(t.HealthId, t.Name, t.MasterHealthEnum))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_health", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterHealthMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *masterHealth.MasterHealth) (*masterHealth.MasterHealth, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterHealth{
		HealthId:         m.HealthId,
		Name:             m.Name,
		MasterHealthEnum: m.MasterHealthEnum,
	}
	res := conn.Model(NewMasterHealth()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterHealth.SetMasterHealth(t.HealthId, t.Name, t.MasterHealthEnum), nil
}

func (s *masterHealthMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterHealth.MasterHealths) (masterHealth.MasterHealths, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

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
			MasterHealthEnum: m.MasterHealthEnum,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterHealth()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterHealthMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *masterHealth.MasterHealth) (*masterHealth.MasterHealth, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterHealth{
		HealthId:         m.HealthId,
		Name:             m.Name,
		MasterHealthEnum: m.MasterHealthEnum,
	}
	res := conn.Model(NewMasterHealth()).WithContext(ctx).Select("health_id", "name", "master_health_enum").Where("health_id = ?", m.HealthId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterHealth.SetMasterHealth(t.HealthId, t.Name, t.MasterHealthEnum), nil
}

func (s *masterHealthMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterHealth.MasterHealths) (masterHealth.MasterHealths, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

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
			MasterHealthEnum: m.MasterHealthEnum,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterHealth()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "health_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "master_health_enum"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterHealthMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *masterHealth.MasterHealth) error {
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

func (s *masterHealthMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterHealth.MasterHealths) error {
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
		ks = append(ks, []interface{}{m.HealthId})
	}

	res := conn.Model(NewMasterHealth()).WithContext(ctx).Where("(health_id) IN ?", ks).Delete(NewMasterHealth())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
