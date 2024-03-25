package service

import (
	"context"
	"webook-server/internal/domain"
)

type ArticleService interface {
	CreateOrUpdate(ctx context.Context, article domain.Article) (int64, error)
}

type ArticleSvc struct {
}

func NewArticleService() ArticleService {
	return &ArticleSvc{}
}

func (svc *ArticleSvc) CreateOrUpdate(ctx context.Context, article domain.Article) (int64, error) {
	return 1, nil
}
