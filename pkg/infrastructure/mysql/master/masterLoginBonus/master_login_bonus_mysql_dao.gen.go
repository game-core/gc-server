// Package masterLoginBonus ログインボーナス
package masterLoginBonus

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/masterLoginBonus"
)

type masterLoginBonusMysqlDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterLoginBonusMysqlDao(conn *database.MysqlHandler) masterLoginBonus.MasterLoginBonusMysqlRepository {
	return &masterLoginBonusMysqlDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterLoginBonusMysqlDao) Find(ctx context.Context, masterLoginBonusId int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "Find", fmt.Sprintf("%d_", masterLoginBonusId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.MasterLoginBonusId, t.MasterEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "Find", fmt.Sprintf("%d_", masterLoginBonusId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusMysqlDao) FindOrNil(ctx context.Context, masterLoginBonusId int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindOrNil", fmt.Sprintf("%d_", masterLoginBonusId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.MasterLoginBonusId, t.MasterEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindOrNil", fmt.Sprintf("%d_", masterLoginBonusId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusMysqlDao) FindByMasterEventId(ctx context.Context, masterEventId int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindByMasterEventId", fmt.Sprintf("%d_", masterEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_event_id = ?", masterEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.MasterLoginBonusId, t.MasterEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindByMasterEventId", fmt.Sprintf("%d_", masterEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusMysqlDao) FindOrNilByMasterEventId(ctx context.Context, masterEventId int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindOrNilByMasterEventId", fmt.Sprintf("%d_", masterEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_event_id = ?", masterEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.MasterLoginBonusId, t.MasterEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindOrNilByMasterEventId", fmt.Sprintf("%d_", masterEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusMysqlDao) FindList(ctx context.Context) (masterLoginBonus.MasterLoginBonuses, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonus.MasterLoginBonuses); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonuses()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonus.NewMasterLoginBonuses()
	for _, t := range ts {
		ms = append(ms, masterLoginBonus.SetMasterLoginBonus(t.MasterLoginBonusId, t.MasterEventId, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusMysqlDao) FindListByMasterEventId(ctx context.Context, masterEventId int64) (masterLoginBonus.MasterLoginBonuses, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindListByMasterEventId", fmt.Sprintf("%d_", masterEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonus.MasterLoginBonuses); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonuses()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_event_id = ?", masterEventId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonus.NewMasterLoginBonuses()
	for _, t := range ts {
		ms = append(ms, masterLoginBonus.SetMasterLoginBonus(t.MasterLoginBonusId, t.MasterEventId, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindListByMasterEventId", fmt.Sprintf("%d_", masterEventId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *masterLoginBonus.MasterLoginBonus) (*masterLoginBonus.MasterLoginBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterLoginBonus{
		MasterLoginBonusId: m.MasterLoginBonusId,
		MasterEventId:      m.MasterEventId,
		Name:               m.Name,
	}
	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonus.SetMasterLoginBonus(t.MasterLoginBonusId, t.MasterEventId, t.Name), nil
}

func (s *masterLoginBonusMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterLoginBonus.MasterLoginBonuses) (masterLoginBonus.MasterLoginBonuses, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterLoginBonuses()
	for _, m := range ms {
		t := &MasterLoginBonus{
			MasterLoginBonusId: m.MasterLoginBonusId,
			MasterEventId:      m.MasterEventId,
			Name:               m.Name,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterLoginBonusMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *masterLoginBonus.MasterLoginBonus) (*masterLoginBonus.MasterLoginBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterLoginBonus{
		MasterLoginBonusId: m.MasterLoginBonusId,
		MasterEventId:      m.MasterEventId,
		Name:               m.Name,
	}
	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Select("master_login_bonus_id", "master_event_id", "name").Where("master_login_bonus_id = ?", m.MasterLoginBonusId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonus.SetMasterLoginBonus(t.MasterLoginBonusId, t.MasterEventId, t.Name), nil
}

func (s *masterLoginBonusMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterLoginBonus.MasterLoginBonuses) (masterLoginBonus.MasterLoginBonuses, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterLoginBonuses()
	for _, m := range ms {
		t := &MasterLoginBonus{
			MasterLoginBonusId: m.MasterLoginBonusId,
			MasterEventId:      m.MasterEventId,
			Name:               m.Name,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterLoginBonus()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_login_bonus_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"master_event_id", "name"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterLoginBonusMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *masterLoginBonus.MasterLoginBonus) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Where("master_login_bonus_id = ?", m.MasterLoginBonusId).Delete(NewMasterLoginBonus())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterLoginBonusMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterLoginBonus.MasterLoginBonuses) error {
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
		ks = append(ks, []interface{}{m.MasterLoginBonusId})
	}

	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Where("(master_login_bonus_id) IN ?", ks).Delete(NewMasterLoginBonus())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
