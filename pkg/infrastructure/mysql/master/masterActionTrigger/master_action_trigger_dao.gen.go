// Package masterActionTrigger アクショントリガー
package masterActionTrigger

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionTrigger"
)

type masterActionTriggerDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterActionTriggerDao(conn *database.MysqlHandler) masterActionTrigger.MasterActionTriggerMysqlRepository {
	return &masterActionTriggerDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterActionTriggerDao) Find(ctx context.Context, masterActionTriggerId int64) (*masterActionTrigger.MasterActionTrigger, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "Find", fmt.Sprintf("%d_", masterActionTriggerId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionTrigger.MasterActionTrigger); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionTrigger()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_trigger_id = ?", masterActionTriggerId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionTrigger.SetMasterActionTrigger(t.MasterActionTriggerId, t.Name, t.MasterActionTriggerEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "Find", fmt.Sprintf("%d_", masterActionTriggerId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionTriggerDao) FindOrNil(ctx context.Context, masterActionTriggerId int64) (*masterActionTrigger.MasterActionTrigger, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "FindOrNil", fmt.Sprintf("%d_", masterActionTriggerId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionTrigger.MasterActionTrigger); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionTrigger()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_trigger_id = ?", masterActionTriggerId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionTrigger.SetMasterActionTrigger(t.MasterActionTriggerId, t.Name, t.MasterActionTriggerEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "FindOrNil", fmt.Sprintf("%d_", masterActionTriggerId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionTriggerDao) FindByMasterActionTriggerEnum(ctx context.Context, masterActionTriggerEnum masterActionTrigger.MasterActionTriggerEnum) (*masterActionTrigger.MasterActionTrigger, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "FindByMasterActionTriggerEnum", fmt.Sprintf("%d_", masterActionTriggerEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionTrigger.MasterActionTrigger); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionTrigger()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_trigger_enum = ?", masterActionTriggerEnum).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionTrigger.SetMasterActionTrigger(t.MasterActionTriggerId, t.Name, t.MasterActionTriggerEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "FindByMasterActionTriggerEnum", fmt.Sprintf("%d_", masterActionTriggerEnum)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionTriggerDao) FindOrNilByMasterActionTriggerEnum(ctx context.Context, masterActionTriggerEnum masterActionTrigger.MasterActionTriggerEnum) (*masterActionTrigger.MasterActionTrigger, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "FindOrNilByMasterActionTriggerEnum", fmt.Sprintf("%d_", masterActionTriggerEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionTrigger.MasterActionTrigger); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionTrigger()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_trigger_enum = ?", masterActionTriggerEnum).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionTrigger.SetMasterActionTrigger(t.MasterActionTriggerId, t.Name, t.MasterActionTriggerEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "FindOrNilByMasterActionTriggerEnum", fmt.Sprintf("%d_", masterActionTriggerEnum)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionTriggerDao) FindList(ctx context.Context) (masterActionTrigger.MasterActionTriggers, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionTrigger.MasterActionTriggers); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionTriggers()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionTrigger.NewMasterActionTriggers()
	for _, t := range ts {
		ms = append(ms, masterActionTrigger.SetMasterActionTrigger(t.MasterActionTriggerId, t.Name, t.MasterActionTriggerEnum))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionTriggerDao) FindListByMasterActionTriggerEnum(ctx context.Context, masterActionTriggerEnum masterActionTrigger.MasterActionTriggerEnum) (masterActionTrigger.MasterActionTriggers, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_trigger", "FindListByMasterActionTriggerEnum", fmt.Sprintf("%d_", masterActionTriggerEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionTrigger.MasterActionTriggers); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionTriggers()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_trigger_enum = ?", masterActionTriggerEnum).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionTrigger.NewMasterActionTriggers()
	for _, t := range ts {
		ms = append(ms, masterActionTrigger.SetMasterActionTrigger(t.MasterActionTriggerId, t.Name, t.MasterActionTriggerEnum))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_trigger", "FindListByMasterActionTriggerEnum", fmt.Sprintf("%d_", masterActionTriggerEnum)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionTriggerDao) Create(ctx context.Context, tx *gorm.DB, m *masterActionTrigger.MasterActionTrigger) (*masterActionTrigger.MasterActionTrigger, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterActionTrigger{
		MasterActionTriggerId:   m.MasterActionTriggerId,
		Name:                    m.Name,
		MasterActionTriggerEnum: m.MasterActionTriggerEnum,
	}
	res := conn.Model(NewMasterActionTrigger()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionTrigger.SetMasterActionTrigger(t.MasterActionTriggerId, t.Name, t.MasterActionTriggerEnum), nil
}

func (s *masterActionTriggerDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterActionTrigger.MasterActionTriggers) (masterActionTrigger.MasterActionTriggers, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterActionTriggers()
	for _, m := range ms {
		t := &MasterActionTrigger{
			MasterActionTriggerId:   m.MasterActionTriggerId,
			Name:                    m.Name,
			MasterActionTriggerEnum: m.MasterActionTriggerEnum,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterActionTrigger()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterActionTriggerDao) Update(ctx context.Context, tx *gorm.DB, m *masterActionTrigger.MasterActionTrigger) (*masterActionTrigger.MasterActionTrigger, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterActionTrigger{
		MasterActionTriggerId:   m.MasterActionTriggerId,
		Name:                    m.Name,
		MasterActionTriggerEnum: m.MasterActionTriggerEnum,
	}
	res := conn.Model(NewMasterActionTrigger()).WithContext(ctx).Where("master_action_trigger_id = ?", m.MasterActionTriggerId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionTrigger.SetMasterActionTrigger(t.MasterActionTriggerId, t.Name, t.MasterActionTriggerEnum), nil
}

func (s *masterActionTriggerDao) Delete(ctx context.Context, tx *gorm.DB, m *masterActionTrigger.MasterActionTrigger) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterActionTrigger()).WithContext(ctx).Where("master_action_trigger_id = ?", m.MasterActionTriggerId).Delete(NewMasterActionTrigger())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}