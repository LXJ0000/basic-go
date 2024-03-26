package cache

import "github.com/redis/go-redis/v9"

type ArticleCACHE interface {
}

type ArticleCache struct {
	cmd redis.Cmdable
}

func NewArticleCache(cmd redis.Cmdable) ArticleCACHE {
	return &ArticleCache{cmd: cmd}
}

//func (r *ArticleCache) Create(ctx context.Context, article domain.Article) error {
//	return nil
//}
//func (r *ArticleCache) Update(ctx context.Context, article domain.Article) error {
//	return nil
//}
