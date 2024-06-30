// Package masterLoginBonusSchedule ログインボーナススケジュール
package masterLoginBonusSchedule

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/masterLoginBonusSchedule"
)

type masterLoginBonusScheduleMysqlDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterLoginBonusScheduleMysqlDao(conn *database.MysqlHandler) masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository {
	return &masterLoginBonusScheduleMysqlDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterLoginBonusScheduleMysqlDao) Find(ctx context.Context, masterLoginBonusScheduleId int64) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "Find", fmt.Sprintf("%d_", masterLoginBonusScheduleId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusSchedule.MasterLoginBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusSchedule()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_schedule_id = ?", masterLoginBonusScheduleId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "Find", fmt.Sprintf("%d_", masterLoginBonusScheduleId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusScheduleMysqlDao) FindOrNil(ctx context.Context, masterLoginBonusScheduleId int64) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "FindOrNil", fmt.Sprintf("%d_", masterLoginBonusScheduleId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusSchedule.MasterLoginBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusSchedule()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_schedule_id = ?", masterLoginBonusScheduleId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "FindOrNil", fmt.Sprintf("%d_", masterLoginBonusScheduleId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusScheduleMysqlDao) FindByMasterLoginBonusId(ctx context.Context, masterLoginBonusId int64) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "FindByMasterLoginBonusId", fmt.Sprintf("%d_", masterLoginBonusId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusSchedule.MasterLoginBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusSchedule()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "FindByMasterLoginBonusId", fmt.Sprintf("%d_", masterLoginBonusId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusScheduleMysqlDao) FindByStep(ctx context.Context, step int32) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "FindByStep", fmt.Sprintf("%d_", step)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusSchedule.MasterLoginBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusSchedule()
	res := s.ReadMysqlConn.WithContext(ctx).Where("step = ?", step).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "FindByStep", fmt.Sprintf("%d_", step)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusScheduleMysqlDao) FindByMasterLoginBonusIdAndStep(ctx context.Context, masterLoginBonusId int64, step int32) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "FindByMasterLoginBonusIdAndStep", fmt.Sprintf("%d_%d_", masterLoginBonusId, step)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusSchedule.MasterLoginBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusSchedule()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_id = ?", masterLoginBonusId).Where("step = ?", step).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "FindByMasterLoginBonusIdAndStep", fmt.Sprintf("%d_%d_", masterLoginBonusId, step)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusScheduleMysqlDao) FindOrNilByMasterLoginBonusId(ctx context.Context, masterLoginBonusId int64) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "FindOrNilByMasterLoginBonusId", fmt.Sprintf("%d_", masterLoginBonusId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusSchedule.MasterLoginBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusSchedule()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "FindOrNilByMasterLoginBonusId", fmt.Sprintf("%d_", masterLoginBonusId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusScheduleMysqlDao) FindOrNilByStep(ctx context.Context, step int32) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "FindOrNilByStep", fmt.Sprintf("%d_", step)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusSchedule.MasterLoginBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusSchedule()
	res := s.ReadMysqlConn.WithContext(ctx).Where("step = ?", step).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "FindOrNilByStep", fmt.Sprintf("%d_", step)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusScheduleMysqlDao) FindOrNilByMasterLoginBonusIdAndStep(ctx context.Context, masterLoginBonusId int64, step int32) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "FindOrNilByMasterLoginBonusIdAndStep", fmt.Sprintf("%d_%d_", masterLoginBonusId, step)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusSchedule.MasterLoginBonusSchedule); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusSchedule()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_id = ?", masterLoginBonusId).Where("step = ?", step).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "FindOrNilByMasterLoginBonusIdAndStep", fmt.Sprintf("%d_%d_", masterLoginBonusId, step)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusScheduleMysqlDao) FindList(ctx context.Context) (masterLoginBonusSchedule.MasterLoginBonusSchedules, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonusSchedule.MasterLoginBonusSchedules); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonusSchedules()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonusSchedule.NewMasterLoginBonusSchedules()
	for _, t := range ts {
		ms = append(ms, masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusScheduleMysqlDao) FindListByMasterLoginBonusId(ctx context.Context, masterLoginBonusId int64) (masterLoginBonusSchedule.MasterLoginBonusSchedules, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "FindListByMasterLoginBonusId", fmt.Sprintf("%d_", masterLoginBonusId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonusSchedule.MasterLoginBonusSchedules); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonusSchedules()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonusSchedule.NewMasterLoginBonusSchedules()
	for _, t := range ts {
		ms = append(ms, masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "FindListByMasterLoginBonusId", fmt.Sprintf("%d_", masterLoginBonusId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusScheduleMysqlDao) FindListByStep(ctx context.Context, step int32) (masterLoginBonusSchedule.MasterLoginBonusSchedules, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "FindListByStep", fmt.Sprintf("%d_", step)))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonusSchedule.MasterLoginBonusSchedules); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonusSchedules()
	res := s.ReadMysqlConn.WithContext(ctx).Where("step = ?", step).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonusSchedule.NewMasterLoginBonusSchedules()
	for _, t := range ts {
		ms = append(ms, masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "FindListByStep", fmt.Sprintf("%d_", step)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusScheduleMysqlDao) FindListByMasterLoginBonusIdAndStep(ctx context.Context, masterLoginBonusId int64, step int32) (masterLoginBonusSchedule.MasterLoginBonusSchedules, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_schedule", "FindListByMasterLoginBonusIdAndStep", fmt.Sprintf("%d_%d_", masterLoginBonusId, step)))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonusSchedule.MasterLoginBonusSchedules); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonusSchedules()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_id = ?", masterLoginBonusId).Where("step = ?", step).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonusSchedule.NewMasterLoginBonusSchedules()
	for _, t := range ts {
		ms = append(ms, masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_schedule", "FindListByMasterLoginBonusIdAndStep", fmt.Sprintf("%d_%d_", masterLoginBonusId, step)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusScheduleMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *masterLoginBonusSchedule.MasterLoginBonusSchedule) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterLoginBonusSchedule{
		MasterLoginBonusScheduleId: m.MasterLoginBonusScheduleId,
		MasterLoginBonusId:         m.MasterLoginBonusId,
		Step:                       m.Step,
		Name:                       m.Name,
	}
	res := conn.Model(NewMasterLoginBonusSchedule()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name), nil
}

func (s *masterLoginBonusScheduleMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterLoginBonusSchedule.MasterLoginBonusSchedules) (masterLoginBonusSchedule.MasterLoginBonusSchedules, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterLoginBonusSchedules()
	for _, m := range ms {
		t := &MasterLoginBonusSchedule{
			MasterLoginBonusScheduleId: m.MasterLoginBonusScheduleId,
			MasterLoginBonusId:         m.MasterLoginBonusId,
			Step:                       m.Step,
			Name:                       m.Name,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterLoginBonusSchedule()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterLoginBonusScheduleMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *masterLoginBonusSchedule.MasterLoginBonusSchedule) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterLoginBonusSchedule{
		MasterLoginBonusScheduleId: m.MasterLoginBonusScheduleId,
		MasterLoginBonusId:         m.MasterLoginBonusId,
		Step:                       m.Step,
		Name:                       m.Name,
	}
	res := conn.Model(NewMasterLoginBonusSchedule()).WithContext(ctx).Select("master_login_bonus_schedule_id", "master_login_bonus_id", "step", "name").Where("master_login_bonus_schedule_id = ?", m.MasterLoginBonusScheduleId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonusSchedule.SetMasterLoginBonusSchedule(t.MasterLoginBonusScheduleId, t.MasterLoginBonusId, t.Step, t.Name), nil
}

func (s *masterLoginBonusScheduleMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterLoginBonusSchedule.MasterLoginBonusSchedules) (masterLoginBonusSchedule.MasterLoginBonusSchedules, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterLoginBonusSchedules()
	for _, m := range ms {
		t := &MasterLoginBonusSchedule{
			MasterLoginBonusScheduleId: m.MasterLoginBonusScheduleId,
			MasterLoginBonusId:         m.MasterLoginBonusId,
			Step:                       m.Step,
			Name:                       m.Name,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterLoginBonusSchedule()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_login_bonus_schedule_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"master_login_bonus_id", "step", "name"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterLoginBonusScheduleMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *masterLoginBonusSchedule.MasterLoginBonusSchedule) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterLoginBonusSchedule()).WithContext(ctx).Where("master_login_bonus_schedule_id = ?", m.MasterLoginBonusScheduleId).Delete(NewMasterLoginBonusSchedule())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterLoginBonusScheduleMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterLoginBonusSchedule.MasterLoginBonusSchedules) error {
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
		ks = append(ks, []interface{}{m.MasterLoginBonusScheduleId})
	}

	res := conn.Model(NewMasterLoginBonusSchedule()).WithContext(ctx).Where("(master_login_bonus_schedule_id) IN ?", ks).Delete(NewMasterLoginBonusSchedule())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
