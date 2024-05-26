// Package masterActionStep アクションステップ
package masterActionStep

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionStep"
)

type masterActionStepDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterActionStepDao(conn *database.MysqlHandler) masterActionStep.MasterActionStepMysqlRepository {
	return &masterActionStepDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterActionStepDao) Find(ctx context.Context, masterActionStepId int64) (*masterActionStep.MasterActionStep, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "Find", fmt.Sprintf("%d_", masterActionStepId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionStep.MasterActionStep); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionStep()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_step_id = ?", masterActionStepId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionStep.SetMasterActionStep(t.MasterActionStepId, t.Name, t.MasterActionStepEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "Find", fmt.Sprintf("%d_", masterActionStepId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionStepDao) FindOrNil(ctx context.Context, masterActionStepId int64) (*masterActionStep.MasterActionStep, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "FindOrNil", fmt.Sprintf("%d_", masterActionStepId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionStep.MasterActionStep); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionStep()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_step_id = ?", masterActionStepId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionStep.SetMasterActionStep(t.MasterActionStepId, t.Name, t.MasterActionStepEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "FindOrNil", fmt.Sprintf("%d_", masterActionStepId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionStepDao) FindByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (*masterActionStep.MasterActionStep, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "FindByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionStep.MasterActionStep); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionStep()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_step_enum = ?", masterActionStepEnum).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionStep.SetMasterActionStep(t.MasterActionStepId, t.Name, t.MasterActionStepEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "FindByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionStepDao) FindOrNilByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (*masterActionStep.MasterActionStep, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "FindOrNilByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionStep.MasterActionStep); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionStep()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_step_enum = ?", masterActionStepEnum).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionStep.SetMasterActionStep(t.MasterActionStepId, t.Name, t.MasterActionStepEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "FindOrNilByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionStepDao) FindList(ctx context.Context) (masterActionStep.MasterActionSteps, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionStep.MasterActionSteps); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionSteps()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionStep.NewMasterActionSteps()
	for _, t := range ts {
		ms = append(ms, masterActionStep.SetMasterActionStep(t.MasterActionStepId, t.Name, t.MasterActionStepEnum))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionStepDao) FindListByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (masterActionStep.MasterActionSteps, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_step", "FindListByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionStep.MasterActionSteps); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionSteps()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_step_enum = ?", masterActionStepEnum).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionStep.NewMasterActionSteps()
	for _, t := range ts {
		ms = append(ms, masterActionStep.SetMasterActionStep(t.MasterActionStepId, t.Name, t.MasterActionStepEnum))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_step", "FindListByMasterActionStepEnum", fmt.Sprintf("%d_", masterActionStepEnum)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionStepDao) Create(ctx context.Context, tx *gorm.DB, m *masterActionStep.MasterActionStep) (*masterActionStep.MasterActionStep, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterActionStep{
		MasterActionStepId:   m.MasterActionStepId,
		Name:                 m.Name,
		MasterActionStepEnum: m.MasterActionStepEnum,
	}
	res := conn.Model(NewMasterActionStep()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionStep.SetMasterActionStep(t.MasterActionStepId, t.Name, t.MasterActionStepEnum), nil
}

func (s *masterActionStepDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterActionStep.MasterActionSteps) (masterActionStep.MasterActionSteps, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterActionSteps()
	for _, m := range ms {
		t := &MasterActionStep{
			MasterActionStepId:   m.MasterActionStepId,
			Name:                 m.Name,
			MasterActionStepEnum: m.MasterActionStepEnum,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterActionStep()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterActionStepDao) Update(ctx context.Context, tx *gorm.DB, m *masterActionStep.MasterActionStep) (*masterActionStep.MasterActionStep, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterActionStep{
		MasterActionStepId:   m.MasterActionStepId,
		Name:                 m.Name,
		MasterActionStepEnum: m.MasterActionStepEnum,
	}
	res := conn.Model(NewMasterActionStep()).WithContext(ctx).Select("master_action_step_id", "name", "master_action_step_enum").Where("master_action_step_id = ?", m.MasterActionStepId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionStep.SetMasterActionStep(t.MasterActionStepId, t.Name, t.MasterActionStepEnum), nil
}

func (s *masterActionStepDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterActionStep.MasterActionSteps) (masterActionStep.MasterActionSteps, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterActionSteps()
	for _, m := range ms {
		t := &MasterActionStep{
			MasterActionStepId:   m.MasterActionStepId,
			Name:                 m.Name,
			MasterActionStepEnum: m.MasterActionStepEnum,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterActionStep()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_action_step_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "master_action_step_enum"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterActionStepDao) Delete(ctx context.Context, tx *gorm.DB, m *masterActionStep.MasterActionStep) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterActionStep()).WithContext(ctx).Where("master_action_step_id = ?", m.MasterActionStepId).Delete(NewMasterActionStep())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterActionStepDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterActionStep.MasterActionSteps) error {
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
		ks = append(ks, []interface{}{m.MasterActionStepId})
	}

	res := conn.Model(NewMasterActionStep()).WithContext(ctx).Where("(master_action_step_id) IN ?", ks).Delete(NewMasterActionStep())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
