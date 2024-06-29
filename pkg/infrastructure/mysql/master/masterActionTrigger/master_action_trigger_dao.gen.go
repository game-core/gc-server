// Package masterActionTrigger アクショントリガー
package masterActionTrigger

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionTrigger"
)

type masterActionTriggerMysqlDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterActionTriggerMysqlDao(conn *database.MysqlHandler) masterActionTrigger.MasterActionTriggerMysqlRepository {
	return &masterActionTriggerMysqlDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterActionTriggerMysqlDao) Find(ctx context.Context, masterActionTriggerId int64) (*masterActionTrigger.MasterActionTrigger, error) {
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

func (s *masterActionTriggerMysqlDao) FindOrNil(ctx context.Context, masterActionTriggerId int64) (*masterActionTrigger.MasterActionTrigger, error) {
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

func (s *masterActionTriggerMysqlDao) FindByMasterActionTriggerEnum(ctx context.Context, masterActionTriggerEnum masterActionTrigger.MasterActionTriggerEnum) (*masterActionTrigger.MasterActionTrigger, error) {
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

func (s *masterActionTriggerMysqlDao) FindOrNilByMasterActionTriggerEnum(ctx context.Context, masterActionTriggerEnum masterActionTrigger.MasterActionTriggerEnum) (*masterActionTrigger.MasterActionTrigger, error) {
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

func (s *masterActionTriggerMysqlDao) FindList(ctx context.Context) (masterActionTrigger.MasterActionTriggers, error) {
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

func (s *masterActionTriggerMysqlDao) FindListByMasterActionTriggerEnum(ctx context.Context, masterActionTriggerEnum masterActionTrigger.MasterActionTriggerEnum) (masterActionTrigger.MasterActionTriggers, error) {
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

func (s *masterActionTriggerMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *masterActionTrigger.MasterActionTrigger) (*masterActionTrigger.MasterActionTrigger, error) {
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

func (s *masterActionTriggerMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterActionTrigger.MasterActionTriggers) (masterActionTrigger.MasterActionTriggers, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

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

func (s *masterActionTriggerMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *masterActionTrigger.MasterActionTrigger) (*masterActionTrigger.MasterActionTrigger, error) {
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
	res := conn.Model(NewMasterActionTrigger()).WithContext(ctx).Select("master_action_trigger_id", "name", "master_action_trigger_enum").Where("master_action_trigger_id = ?", m.MasterActionTriggerId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionTrigger.SetMasterActionTrigger(t.MasterActionTriggerId, t.Name, t.MasterActionTriggerEnum), nil
}

func (s *masterActionTriggerMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterActionTrigger.MasterActionTriggers) (masterActionTrigger.MasterActionTriggers, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

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

	res := conn.Model(NewMasterActionTrigger()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_action_trigger_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "master_action_trigger_enum"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterActionTriggerMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *masterActionTrigger.MasterActionTrigger) error {
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

func (s *masterActionTriggerMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterActionTrigger.MasterActionTriggers) error {
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
		ks = append(ks, []interface{}{m.MasterActionTriggerId})
	}

	res := conn.Model(NewMasterActionTrigger()).WithContext(ctx).Where("(master_action_trigger_id) IN ?", ks).Delete(NewMasterActionTrigger())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
