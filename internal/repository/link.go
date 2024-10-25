package repository

import (
	"context"
	"errors"
	"fmt"
	"golang-url-shortener/internal/model"

	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

type ILinkRepo interface {
	Insert(ctx context.Context, data *model.LinkModel) error
	FindOneByShort(ctx context.Context, short string) (model.LinkModel, error)
	FindOneByURL(ctx context.Context, url string) (model.LinkModel, error)
}

type linkRepo struct {
	db *gorm.DB
}

func NewLinkRepo(
	db *gorm.DB,
) ILinkRepo {
	return &linkRepo{
		db: db,
	}
}

func (s *linkRepo) table(ctx context.Context) *gorm.DB {
	return s.db.Table(fmt.Sprintf("public.%s", model.LinkModel{}.TableName())).WithContext(ctx)
}

func (l *linkRepo) Insert(ctx context.Context, data *model.LinkModel) error {
	tx := l.table(ctx).Create(data)
	if tx.Error != nil {
		log.Error().Ctx(ctx).Err(tx.Error).Msg("insert link error")
	}
	return tx.Error
}

func (s *linkRepo) FindOneByShort(ctx context.Context, short string) (model.LinkModel, error) {
	var data model.LinkModel
	tx := s.table(ctx).Where("short = ?", short).First(&data)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		log.Error().Ctx(ctx).Stack().Err(tx.Error).Msg("find one by short link error")
		return model.LinkModel{}, tx.Error
	}
	return data, nil
}

func (s *linkRepo) FindOneByURL(ctx context.Context, url string) (model.LinkModel, error) {
	var data model.LinkModel
	tx := s.table(ctx).Where("original_url = ?", url).First(&data)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		log.Error().Ctx(ctx).Stack().Err(tx.Error).Msg("find one by url link error")
		return model.LinkModel{}, tx.Error
	}
	return data, nil
}
