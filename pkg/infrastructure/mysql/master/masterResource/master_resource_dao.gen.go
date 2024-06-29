// Package masterResource リソース
package masterResource

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/resource/masterResource"
)

type masterResourceMysqlDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterResourceMysqlDao(conn *database.MysqlHandler) masterResource.MasterResourceMysqlRepository {
	return &masterResourceMysqlDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterResourceMysqlDao) Find(ctx context.Context, masterResourceId int64) (*masterResource.MasterResource, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "Find", fmt.Sprintf("%d_", masterResourceId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterResource.MasterResource); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterResource()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_resource_id = ?", masterResourceId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterResource.SetMasterResource(t.MasterResourceId, t.Name, t.MasterResourceEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_resource", "Find", fmt.Sprintf("%d_", masterResourceId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterResourceMysqlDao) FindOrNil(ctx context.Context, masterResourceId int64) (*masterResource.MasterResource, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "FindOrNil", fmt.Sprintf("%d_", masterResourceId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterResource.MasterResource); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterResource()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_resource_id = ?", masterResourceId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterResource.SetMasterResource(t.MasterResourceId, t.Name, t.MasterResourceEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_resource", "FindOrNil", fmt.Sprintf("%d_", masterResourceId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterResourceMysqlDao) FindByMasterResourceEnum(ctx context.Context, masterResourceEnum masterResource.MasterResourceEnum) (*masterResource.MasterResource, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "FindByMasterResourceEnum", fmt.Sprintf("%d_", masterResourceEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterResource.MasterResource); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterResource()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_resource_enum = ?", masterResourceEnum).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterResource.SetMasterResource(t.MasterResourceId, t.Name, t.MasterResourceEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_resource", "FindByMasterResourceEnum", fmt.Sprintf("%d_", masterResourceEnum)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterResourceMysqlDao) FindOrNilByMasterResourceEnum(ctx context.Context, masterResourceEnum masterResource.MasterResourceEnum) (*masterResource.MasterResource, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "FindOrNilByMasterResourceEnum", fmt.Sprintf("%d_", masterResourceEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterResource.MasterResource); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterResource()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_resource_enum = ?", masterResourceEnum).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterResource.SetMasterResource(t.MasterResourceId, t.Name, t.MasterResourceEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_resource", "FindOrNilByMasterResourceEnum", fmt.Sprintf("%d_", masterResourceEnum)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterResourceMysqlDao) FindList(ctx context.Context) (masterResource.MasterResources, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterResource.MasterResources); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterResources()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterResource.NewMasterResources()
	for _, t := range ts {
		ms = append(ms, masterResource.SetMasterResource(t.MasterResourceId, t.Name, t.MasterResourceEnum))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_resource", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterResourceMysqlDao) FindListByMasterResourceEnum(ctx context.Context, masterResourceEnum masterResource.MasterResourceEnum) (masterResource.MasterResources, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_resource", "FindListByMasterResourceEnum", fmt.Sprintf("%d_", masterResourceEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(masterResource.MasterResources); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterResources()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_resource_enum = ?", masterResourceEnum).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterResource.NewMasterResources()
	for _, t := range ts {
		ms = append(ms, masterResource.SetMasterResource(t.MasterResourceId, t.Name, t.MasterResourceEnum))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_resource", "FindListByMasterResourceEnum", fmt.Sprintf("%d_", masterResourceEnum)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterResourceMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *masterResource.MasterResource) (*masterResource.MasterResource, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterResource{
		MasterResourceId:   m.MasterResourceId,
		Name:               m.Name,
		MasterResourceEnum: m.MasterResourceEnum,
	}
	res := conn.Model(NewMasterResource()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterResource.SetMasterResource(t.MasterResourceId, t.Name, t.MasterResourceEnum), nil
}

func (s *masterResourceMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterResource.MasterResources) (masterResource.MasterResources, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterResources()
	for _, m := range ms {
		t := &MasterResource{
			MasterResourceId:   m.MasterResourceId,
			Name:               m.Name,
			MasterResourceEnum: m.MasterResourceEnum,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterResource()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterResourceMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *masterResource.MasterResource) (*masterResource.MasterResource, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterResource{
		MasterResourceId:   m.MasterResourceId,
		Name:               m.Name,
		MasterResourceEnum: m.MasterResourceEnum,
	}
	res := conn.Model(NewMasterResource()).WithContext(ctx).Select("master_resource_id", "name", "master_resource_enum").Where("master_resource_id = ?", m.MasterResourceId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterResource.SetMasterResource(t.MasterResourceId, t.Name, t.MasterResourceEnum), nil
}

func (s *masterResourceMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterResource.MasterResources) (masterResource.MasterResources, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterResources()
	for _, m := range ms {
		t := &MasterResource{
			MasterResourceId:   m.MasterResourceId,
			Name:               m.Name,
			MasterResourceEnum: m.MasterResourceEnum,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterResource()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_resource_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "master_resource_enum"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterResourceMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *masterResource.MasterResource) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterResource()).WithContext(ctx).Where("master_resource_id = ?", m.MasterResourceId).Delete(NewMasterResource())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterResourceMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterResource.MasterResources) error {
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
		ks = append(ks, []interface{}{m.MasterResourceId})
	}

	res := conn.Model(NewMasterResource()).WithContext(ctx).Where("(master_resource_id) IN ?", ks).Delete(NewMasterResource())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
