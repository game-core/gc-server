// Package commonHealth ヘルスチェック
package commonHealth

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/health/commonHealth"
)

type commonHealthDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

func NewCommonHealthDao(conn *database.MysqlHandler) commonHealth.CommonHealthMysqlRepository {
	return &commonHealthDao{
		ReadMysqlConn:  conn.Common.ReadMysqlConn,
		WriteMysqlConn: conn.Common.WriteMysqlConn,
	}
}

func (s *commonHealthDao) Find(ctx context.Context, healthId int64) (*commonHealth.CommonHealth, error) {
	t := NewCommonHealth()
	res := s.ReadMysqlConn.WithContext(ctx).Where("health_id = ?", healthId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return commonHealth.SetCommonHealth(t.HealthId, t.Name, t.CommonHealthType), nil
}

func (s *commonHealthDao) FindOrNil(ctx context.Context, healthId int64) (*commonHealth.CommonHealth, error) {
	t := NewCommonHealth()
	res := s.ReadMysqlConn.WithContext(ctx).Where("health_id = ?", healthId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return commonHealth.SetCommonHealth(t.HealthId, t.Name, t.CommonHealthType), nil
}

func (s *commonHealthDao) FindList(ctx context.Context) (commonHealth.CommonHealths, error) {
	ts := NewCommonHealths()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := commonHealth.NewCommonHealths()
	for _, t := range ts {
		ms = append(ms, commonHealth.SetCommonHealth(t.HealthId, t.Name, t.CommonHealthType))
	}

	return ms, nil
}

func (s *commonHealthDao) Create(ctx context.Context, tx *gorm.DB, m *commonHealth.CommonHealth) (*commonHealth.CommonHealth, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonHealth{
		HealthId:         m.HealthId,
		Name:             m.Name,
		CommonHealthType: m.CommonHealthType,
	}
	res := conn.Model(NewCommonHealth()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonHealth.SetCommonHealth(t.HealthId, t.Name, t.CommonHealthType), nil
}

func (s *commonHealthDao) CreateList(ctx context.Context, tx *gorm.DB, ms commonHealth.CommonHealths) (commonHealth.CommonHealths, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewCommonHealths()
	for _, m := range ms {
		t := &CommonHealth{
			HealthId:         m.HealthId,
			Name:             m.Name,
			CommonHealthType: m.CommonHealthType,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewCommonHealth()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *commonHealthDao) Update(ctx context.Context, tx *gorm.DB, m *commonHealth.CommonHealth) (*commonHealth.CommonHealth, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &CommonHealth{
		HealthId:         m.HealthId,
		Name:             m.Name,
		CommonHealthType: m.CommonHealthType,
	}
	res := conn.Model(NewCommonHealth()).WithContext(ctx).Where("health_id = ?", m.HealthId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return commonHealth.SetCommonHealth(t.HealthId, t.Name, t.CommonHealthType), nil
}

func (s *commonHealthDao) Delete(ctx context.Context, tx *gorm.DB, m *commonHealth.CommonHealth) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewCommonHealth()).WithContext(ctx).Where("health_id = ?", m.HealthId).Delete(NewCommonHealth())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
