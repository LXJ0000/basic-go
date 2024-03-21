package dao

import (
	"context"
	"gorm.io/gorm"
	"webook-server/internal/model"
)

type MenuDao interface {
	GetAllMenuList(ctx context.Context) ([]model.Menu, error)
}

type MenuDaoByGorm struct {
	db *gorm.DB
}

func NewMenuDao(db *gorm.DB) MenuDao {
	return &MenuDaoByGorm{db: db}
}

func (dao *MenuDaoByGorm) GetAllMenuList(ctx context.Context) ([]model.Menu, error) {
	var menus []model.Menu
	err := dao.db.WithContext(ctx).Model(&model.Menu{}).Find(&menus).Error
	return menus, err
}
