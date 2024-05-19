// Package masterEvent イベント
package masterEvent

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/event/masterEvent"
)

type masterEventDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterEventDao(conn *database.MysqlHandler) masterEvent.MasterEventMysqlRepository {
	return &masterEventDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterEventDao) Find(ctx context.Context, masterEventId int64) (*masterEvent.MasterEvent, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_event", "Find", fmt.Sprintf("%d_", masterEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterEvent.MasterEvent); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterEvent()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_event_id = ?", masterEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterEvent.SetMasterEvent(t.MasterEventId, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt)
	s.Cache.Set(cashes.CreateCacheKey("master_event", "Find", fmt.Sprintf("%d_", masterEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterEventDao) FindOrNil(ctx context.Context, masterEventId int64) (*masterEvent.MasterEvent, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_event", "FindOrNil", fmt.Sprintf("%d_", masterEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterEvent.MasterEvent); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterEvent()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_event_id = ?", masterEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterEvent.SetMasterEvent(t.MasterEventId, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt)
	s.Cache.Set(cashes.CreateCacheKey("master_event", "FindOrNil", fmt.Sprintf("%d_", masterEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterEventDao) FindList(ctx context.Context) (masterEvent.MasterEvents, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_event", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterEvent.MasterEvents); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterEvents()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterEvent.NewMasterEvents()
	for _, t := range ts {
		ms = append(ms, masterEvent.SetMasterEvent(t.MasterEventId, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_event", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterEventDao) Create(ctx context.Context, tx *gorm.DB, m *masterEvent.MasterEvent) (*masterEvent.MasterEvent, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterEvent{
		MasterEventId: m.MasterEventId,
		Name:          m.Name,
		ResetHour:     m.ResetHour,
		IntervalHour:  m.IntervalHour,
		RepeatSetting: m.RepeatSetting,
		StartAt:       m.StartAt,
		EndAt:         m.EndAt,
	}
	res := conn.Model(NewMasterEvent()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterEvent.SetMasterEvent(t.MasterEventId, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt), nil
}

func (s *masterEventDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterEvent.MasterEvents) (masterEvent.MasterEvents, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterEvents()
	for _, m := range ms {
		t := &MasterEvent{
			MasterEventId: m.MasterEventId,
			Name:          m.Name,
			ResetHour:     m.ResetHour,
			IntervalHour:  m.IntervalHour,
			RepeatSetting: m.RepeatSetting,
			StartAt:       m.StartAt,
			EndAt:         m.EndAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterEvent()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterEventDao) Update(ctx context.Context, tx *gorm.DB, m *masterEvent.MasterEvent) (*masterEvent.MasterEvent, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterEvent{
		MasterEventId: m.MasterEventId,
		Name:          m.Name,
		ResetHour:     m.ResetHour,
		IntervalHour:  m.IntervalHour,
		RepeatSetting: m.RepeatSetting,
		StartAt:       m.StartAt,
		EndAt:         m.EndAt,
	}
	res := conn.Model(NewMasterEvent()).WithContext(ctx).Where("master_event_id = ?", m.MasterEventId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterEvent.SetMasterEvent(t.MasterEventId, t.Name, t.ResetHour, t.IntervalHour, t.RepeatSetting, t.StartAt, t.EndAt), nil
}

func (s *masterEventDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterEvent.MasterEvents) (masterEvent.MasterEvents, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterEvents()
	for _, m := range ms {
		t := &MasterEvent{
			MasterEventId: m.MasterEventId,
			Name:          m.Name,
			ResetHour:     m.ResetHour,
			IntervalHour:  m.IntervalHour,
			RepeatSetting: m.RepeatSetting,
			StartAt:       m.StartAt,
			EndAt:         m.EndAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterEvent()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_event_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "reset_hour", "interval_hour", "repeat_setting", "start_at", "end_at"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterEventDao) Delete(ctx context.Context, tx *gorm.DB, m *masterEvent.MasterEvent) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterEvent()).WithContext(ctx).Where("master_event_id = ?", m.MasterEventId).Delete(NewMasterEvent())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterEventDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterEvent.MasterEvents) error {
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
		ks = append(ks, []interface{}{m.MasterEventId})
	}

	res := conn.Model(NewMasterEvent()).WithContext(ctx).Where("(master_event_id) IN ?", ks).Delete(NewMasterEvent())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
