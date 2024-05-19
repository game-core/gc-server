// Package masterAction アクション
package masterAction

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterAction"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionStep"
)

type masterActionDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterActionDao(conn *database.MysqlHandler) masterAction.MasterActionMysqlRepository {
	return &masterActionDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterActionDao) Find(ctx context.Context, masterActionId int64) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "Find", fmt.Sprintf("%d_", masterActionId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_id = ?", masterActionId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "Find", fmt.Sprintf("%d_", masterActionId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindOrNil(ctx context.Context, masterActionId int64) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindOrNil", fmt.Sprintf("%d_", masterActionId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_id = ?", masterActionId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindOrNil", fmt.Sprintf("%d_", masterActionId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_step_enum = ?", masterActionStepEnum).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindByTargetId(ctx context.Context, targetId *int64) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindByTargetId", fmt.Sprintf("%d_", targetId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadMysqlConn.WithContext(ctx).Where("target_id = ?", targetId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindByTargetId", fmt.Sprintf("%d_", targetId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindByMasterActionStepEnumAndTargetId(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindByMasterActionStepEnumAndTargetId", fmt.Sprintf("%d_%d_", masterActionStepEnum, targetId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_step_enum = ?", masterActionStepEnum).Where("target_id = ?", targetId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindByMasterActionStepEnumAndTargetId", fmt.Sprintf("%d_%d_", masterActionStepEnum, targetId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindOrNilByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindOrNilByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_step_enum = ?", masterActionStepEnum).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindOrNilByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindOrNilByTargetId(ctx context.Context, targetId *int64) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindOrNilByTargetId", fmt.Sprintf("%d_", targetId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadMysqlConn.WithContext(ctx).Where("target_id = ?", targetId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindOrNilByTargetId", fmt.Sprintf("%d_", targetId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindOrNilByMasterActionStepEnumAndTargetId(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) (*masterAction.MasterAction, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindOrNilByMasterActionStepEnumAndTargetId", fmt.Sprintf("%d_%d_", masterActionStepEnum, targetId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterAction.MasterAction); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterAction()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_step_enum = ?", masterActionStepEnum).Where("target_id = ?", targetId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration)
	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindOrNilByMasterActionStepEnumAndTargetId", fmt.Sprintf("%d_%d_", masterActionStepEnum, targetId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionDao) FindList(ctx context.Context) (masterAction.MasterActions, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterAction.MasterActions); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActions()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterAction.NewMasterActions()
	for _, t := range ts {
		ms = append(ms, masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionDao) FindListByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (masterAction.MasterActions, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindListByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(masterAction.MasterActions); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActions()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_step_enum = ?", masterActionStepEnum).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterAction.NewMasterActions()
	for _, t := range ts {
		ms = append(ms, masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindListByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionDao) FindListByTargetId(ctx context.Context, targetId *int64) (masterAction.MasterActions, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindListByTargetId", fmt.Sprintf("%d_", targetId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterAction.MasterActions); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActions()
	res := s.ReadMysqlConn.WithContext(ctx).Where("target_id = ?", targetId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterAction.NewMasterActions()
	for _, t := range ts {
		ms = append(ms, masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindListByTargetId", fmt.Sprintf("%d_", targetId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionDao) FindListByMasterActionStepEnumAndTargetId(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) (masterAction.MasterActions, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action", "FindListByMasterActionStepEnumAndTargetId", fmt.Sprintf("%d_%d_", masterActionStepEnum, targetId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterAction.MasterActions); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActions()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_step_enum = ?", masterActionStepEnum).Where("target_id = ?", targetId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterAction.NewMasterActions()
	for _, t := range ts {
		ms = append(ms, masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action", "FindListByMasterActionStepEnumAndTargetId", fmt.Sprintf("%d_%d_", masterActionStepEnum, targetId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionDao) Create(ctx context.Context, tx *gorm.DB, m *masterAction.MasterAction) (*masterAction.MasterAction, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterAction{
		MasterActionId:          m.MasterActionId,
		Name:                    m.Name,
		MasterActionStepEnum:    m.MasterActionStepEnum,
		MasterActionTriggerEnum: m.MasterActionTriggerEnum,
		TargetId:                m.TargetId,
		TriggerMasterActionId:   m.TriggerMasterActionId,
		Expiration:              m.Expiration,
	}
	res := conn.Model(NewMasterAction()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration), nil
}

func (s *masterActionDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterAction.MasterActions) (masterAction.MasterActions, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterActions()
	for _, m := range ms {
		t := &MasterAction{
			MasterActionId:          m.MasterActionId,
			Name:                    m.Name,
			MasterActionStepEnum:    m.MasterActionStepEnum,
			MasterActionTriggerEnum: m.MasterActionTriggerEnum,
			TargetId:                m.TargetId,
			TriggerMasterActionId:   m.TriggerMasterActionId,
			Expiration:              m.Expiration,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterAction()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterActionDao) Update(ctx context.Context, tx *gorm.DB, m *masterAction.MasterAction) (*masterAction.MasterAction, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterAction{
		MasterActionId:          m.MasterActionId,
		Name:                    m.Name,
		MasterActionStepEnum:    m.MasterActionStepEnum,
		MasterActionTriggerEnum: m.MasterActionTriggerEnum,
		TargetId:                m.TargetId,
		TriggerMasterActionId:   m.TriggerMasterActionId,
		Expiration:              m.Expiration,
	}
	res := conn.Model(NewMasterAction()).WithContext(ctx).Where("master_action_id = ?", m.MasterActionId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterAction.SetMasterAction(t.MasterActionId, t.Name, t.MasterActionStepEnum, t.MasterActionTriggerEnum, t.TargetId, t.TriggerMasterActionId, t.Expiration), nil
}

func (s *masterActionDao) Delete(ctx context.Context, tx *gorm.DB, m *masterAction.MasterAction) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterAction()).WithContext(ctx).Where("master_action_id = ?", m.MasterActionId).Delete(NewMasterAction())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
