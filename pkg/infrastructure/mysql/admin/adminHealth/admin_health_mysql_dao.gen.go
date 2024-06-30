// Package adminHealth ヘルスチェック
package adminHealth

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/health/adminHealth"
)

type adminHealthMysqlDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

func NewAdminHealthMysqlDao(conn *database.MysqlHandler) adminHealth.AdminHealthMysqlRepository {
	return &adminHealthMysqlDao{
		ReadMysqlConn:  conn.Admin.ReadMysqlConn,
		WriteMysqlConn: conn.Admin.WriteMysqlConn,
	}
}

func (s *adminHealthMysqlDao) Find(ctx context.Context, healthId int64) (*adminHealth.AdminHealth, error) {
	t := NewAdminHealth()
	res := s.ReadMysqlConn.WithContext(ctx).Where("health_id = ?", healthId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return adminHealth.SetAdminHealth(t.HealthId, t.Name, t.AdminHealthEnum), nil
}

func (s *adminHealthMysqlDao) FindOrNil(ctx context.Context, healthId int64) (*adminHealth.AdminHealth, error) {
	t := NewAdminHealth()
	res := s.ReadMysqlConn.WithContext(ctx).Where("health_id = ?", healthId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return adminHealth.SetAdminHealth(t.HealthId, t.Name, t.AdminHealthEnum), nil
}

func (s *adminHealthMysqlDao) FindList(ctx context.Context) (adminHealth.AdminHealths, error) {
	ts := NewAdminHealths()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := adminHealth.NewAdminHealths()
	for _, t := range ts {
		ms = append(ms, adminHealth.SetAdminHealth(t.HealthId, t.Name, t.AdminHealthEnum))
	}

	return ms, nil
}

func (s *adminHealthMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *adminHealth.AdminHealth) (*adminHealth.AdminHealth, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &AdminHealth{
		HealthId:        m.HealthId,
		Name:            m.Name,
		AdminHealthEnum: m.AdminHealthEnum,
	}
	res := conn.Model(NewAdminHealth()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return adminHealth.SetAdminHealth(t.HealthId, t.Name, t.AdminHealthEnum), nil
}

func (s *adminHealthMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms adminHealth.AdminHealths) (adminHealth.AdminHealths, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewAdminHealths()
	for _, m := range ms {
		t := &AdminHealth{
			HealthId:        m.HealthId,
			Name:            m.Name,
			AdminHealthEnum: m.AdminHealthEnum,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewAdminHealth()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *adminHealthMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *adminHealth.AdminHealth) (*adminHealth.AdminHealth, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &AdminHealth{
		HealthId:        m.HealthId,
		Name:            m.Name,
		AdminHealthEnum: m.AdminHealthEnum,
	}
	res := conn.Model(NewAdminHealth()).WithContext(ctx).Select("health_id", "name", "admin_health_enum").Where("health_id = ?", m.HealthId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return adminHealth.SetAdminHealth(t.HealthId, t.Name, t.AdminHealthEnum), nil
}

func (s *adminHealthMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms adminHealth.AdminHealths) (adminHealth.AdminHealths, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewAdminHealths()
	for _, m := range ms {
		t := &AdminHealth{
			HealthId:        m.HealthId,
			Name:            m.Name,
			AdminHealthEnum: m.AdminHealthEnum,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewAdminHealth()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "health_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "admin_health_enum"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *adminHealthMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *adminHealth.AdminHealth) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewAdminHealth()).WithContext(ctx).Where("health_id = ?", m.HealthId).Delete(NewAdminHealth())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *adminHealthMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms adminHealth.AdminHealths) error {
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

	res := conn.Model(NewAdminHealth()).WithContext(ctx).Where("(health_id) IN ?", ks).Delete(NewAdminHealth())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
