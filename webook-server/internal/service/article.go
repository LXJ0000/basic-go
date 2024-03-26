package service

import (
	"context"
	"webook-server/internal/domain"
	"webook-server/internal/repository"
)

type ArticleService interface {
	CreateOrUpdate(ctx context.Context, article domain.Article) error
}

type ArticleSvc struct {
	repo repository.ArticleRepo
}

func NewArticleService(repo repository.ArticleRepo) ArticleService {
	return &ArticleSvc{repo: repo}
}

func (svc *ArticleSvc) CreateOrUpdate(ctx context.Context, article domain.Article) error {
	return svc.repo.Create(ctx, article)
}
