package repository

import (
	"context"
	"webook-server/internal/model"
	"webook-server/internal/repository/dao"
)

type MenuRepository interface {
	GetAllMenuList(ctx context.Context) ([]model.Menu, error)
}

type MenuRepositoryByDao struct {
	dao dao.MenuDao
}

func NewMenuRepository(dao dao.MenuDao) MenuRepository {
	return &MenuRepositoryByDao{dao: dao}
}

func (r *MenuRepositoryByDao) GetAllMenuList(ctx context.Context) ([]model.Menu, error) {
	return r.dao.GetAllMenuList(ctx)
}
