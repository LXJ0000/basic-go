package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type ArticleDAO interface {
	Create(ctx context.Context, article Article) error
	Update(ctx context.Context, article Article) error
}

type ArticleDao struct {
	db *gorm.DB
}

func NewArticleDao(db *gorm.DB) ArticleDAO {

	return &ArticleDao{db: db}
}

func (r *ArticleDao) Create(ctx context.Context, article Article) error {
	now := time.Now().UnixMilli()
	article.CreateAt = now
	article.UpdateAt = now
	return r.db.WithContext(ctx).Model(&Article{}).Create(&article).Error
}
func (r *ArticleDao) Update(ctx context.Context, article Article) error {
	return nil
}

type Article struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	AuthorId int64  `gorm:"index=idx_authorId_createAt"`
	Title    string `gorm:"varchar(1024)"`
	Content  string `gorm:"BLOB"`

	CreateAt int64 `gorm:"index=idx_authorId_createAt"`
	UpdateAt int64

	//	需求：select * from . where authorId = . order by creatAt
}
