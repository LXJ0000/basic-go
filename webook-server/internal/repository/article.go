package repository

import (
	"context"
	"webook-server/internal/domain"
	"webook-server/internal/repository/cache"
	"webook-server/internal/repository/dao"
)

type ArticleRepo interface {
	Create(ctx context.Context, article domain.Article) error
	Update(ctx context.Context, article domain.Article) error
}

type ArticleRepository struct {
	dao   dao.ArticleDAO
	cache cache.ArticleCACHE
}

func NewArticleRepository(dao dao.ArticleDAO, cache cache.ArticleCACHE) ArticleRepo {
	return &ArticleRepository{dao: dao, cache: cache}
}

func (r *ArticleRepository) Create(ctx context.Context, article domain.Article) error {
	return r.dao.Create(ctx, dao.Article{
		AuthorId: article.AuthorId,
		Title:    article.Title,
		Content:  article.Content,
	})
}
func (r *ArticleRepository) Update(ctx context.Context, article domain.Article) error {
	return nil
}
