// Package masterActionRun 実行されるアクション
package masterActionRun

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionRun"
)

type masterActionRunMysqlDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterActionRunMysqlDao(conn *database.MysqlHandler) masterActionRun.MasterActionRunMysqlRepository {
	return &masterActionRunMysqlDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterActionRunMysqlDao) Find(ctx context.Context, masterActionRunId int64) (*masterActionRun.MasterActionRun, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "Find", fmt.Sprintf("%d_", masterActionRunId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionRun.MasterActionRun); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionRun()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_run_id = ?", masterActionRunId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionRun.SetMasterActionRun(t.MasterActionRunId, t.Name, t.MasterActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "Find", fmt.Sprintf("%d_", masterActionRunId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionRunMysqlDao) FindOrNil(ctx context.Context, masterActionRunId int64) (*masterActionRun.MasterActionRun, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindOrNil", fmt.Sprintf("%d_", masterActionRunId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionRun.MasterActionRun); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionRun()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_run_id = ?", masterActionRunId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionRun.SetMasterActionRun(t.MasterActionRunId, t.Name, t.MasterActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindOrNil", fmt.Sprintf("%d_", masterActionRunId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionRunMysqlDao) FindByMasterActionId(ctx context.Context, masterActionId int64) (*masterActionRun.MasterActionRun, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindByMasterActionId", fmt.Sprintf("%d_", masterActionId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionRun.MasterActionRun); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionRun()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_id = ?", masterActionId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterActionRun.SetMasterActionRun(t.MasterActionRunId, t.Name, t.MasterActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindByMasterActionId", fmt.Sprintf("%d_", masterActionId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionRunMysqlDao) FindOrNilByMasterActionId(ctx context.Context, masterActionId int64) (*masterActionRun.MasterActionRun, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindOrNilByMasterActionId", fmt.Sprintf("%d_", masterActionId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterActionRun.MasterActionRun); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterActionRun()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_id = ?", masterActionId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterActionRun.SetMasterActionRun(t.MasterActionRunId, t.Name, t.MasterActionId)
	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindOrNilByMasterActionId", fmt.Sprintf("%d_", masterActionId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterActionRunMysqlDao) FindList(ctx context.Context) (masterActionRun.MasterActionRuns, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionRun.MasterActionRuns); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionRuns()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionRun.NewMasterActionRuns()
	for _, t := range ts {
		ms = append(ms, masterActionRun.SetMasterActionRun(t.MasterActionRunId, t.Name, t.MasterActionId))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionRunMysqlDao) FindListByMasterActionId(ctx context.Context, masterActionId int64) (masterActionRun.MasterActionRuns, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_action_run", "FindListByMasterActionId", fmt.Sprintf("%d_", masterActionId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterActionRun.MasterActionRuns); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterActionRuns()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_action_id = ?", masterActionId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterActionRun.NewMasterActionRuns()
	for _, t := range ts {
		ms = append(ms, masterActionRun.SetMasterActionRun(t.MasterActionRunId, t.Name, t.MasterActionId))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_action_run", "FindListByMasterActionId", fmt.Sprintf("%d_", masterActionId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterActionRunMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *masterActionRun.MasterActionRun) (*masterActionRun.MasterActionRun, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterActionRun{
		MasterActionRunId: m.MasterActionRunId,
		Name:              m.Name,
		MasterActionId:    m.MasterActionId,
	}
	res := conn.Model(NewMasterActionRun()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionRun.SetMasterActionRun(t.MasterActionRunId, t.Name, t.MasterActionId), nil
}

func (s *masterActionRunMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterActionRun.MasterActionRuns) (masterActionRun.MasterActionRuns, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterActionRuns()
	for _, m := range ms {
		t := &MasterActionRun{
			MasterActionRunId: m.MasterActionRunId,
			Name:              m.Name,
			MasterActionId:    m.MasterActionId,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterActionRun()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterActionRunMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *masterActionRun.MasterActionRun) (*masterActionRun.MasterActionRun, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterActionRun{
		MasterActionRunId: m.MasterActionRunId,
		Name:              m.Name,
		MasterActionId:    m.MasterActionId,
	}
	res := conn.Model(NewMasterActionRun()).WithContext(ctx).Select("master_action_run_id", "name", "master_action_id").Where("master_action_run_id = ?", m.MasterActionRunId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterActionRun.SetMasterActionRun(t.MasterActionRunId, t.Name, t.MasterActionId), nil
}

func (s *masterActionRunMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterActionRun.MasterActionRuns) (masterActionRun.MasterActionRuns, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterActionRuns()
	for _, m := range ms {
		t := &MasterActionRun{
			MasterActionRunId: m.MasterActionRunId,
			Name:              m.Name,
			MasterActionId:    m.MasterActionId,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterActionRun()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_action_run_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "master_action_id"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterActionRunMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *masterActionRun.MasterActionRun) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterActionRun()).WithContext(ctx).Where("master_action_run_id = ?", m.MasterActionRunId).Delete(NewMasterActionRun())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterActionRunMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterActionRun.MasterActionRuns) error {
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
		ks = append(ks, []interface{}{m.MasterActionRunId})
	}

	res := conn.Model(NewMasterActionRun()).WithContext(ctx).Where("(master_action_run_id) IN ?", ks).Delete(NewMasterActionRun())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
