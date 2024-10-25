package service

import (
	"context"
	"golang-url-shortener/internal/dto"
	"golang-url-shortener/internal/model"
	"golang-url-shortener/internal/repository"
	"golang-url-shortener/pkg/crypto"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

type IShortService interface {
	Short(ctx context.Context, data dto.ShortRequest) (dto.ShortResponse, error)
	Find(ctx context.Context, shorten string) (string, error)
}

type shortService struct {
	linkRepo repository.ILinkRepo
	domain string
}

func NewShortService(
	linkRepo repository.ILinkRepo,
) IShortService {
	return &shortService{
		linkRepo: linkRepo,
		domain: os.Getenv("APP_DOMAIN"),
	}
}

func (s *shortService) Short(ctx context.Context, data dto.ShortRequest) (dto.ShortResponse, error) {
	log.Info().Ctx(ctx).Interface("data", data).Msg("shortService Short")
	exist, _ := s.linkRepo.FindOneByShort(ctx, data.URL)
	if exist.ID != 0 {
		return dto.ShortResponse{
			Shorten: strings.Join([]string{s.domain, "at", exist.Short}, "/"),
		}, nil
	}
	
	ans := crypto.ShortenURL(data.URL)
	err := s.linkRepo.Insert(ctx, &model.LinkModel{
		Short:       ans,
		OriginalURL: data.URL,
	})
	if err != nil {
		return dto.ShortResponse{}, err
	}
	return dto.ShortResponse{
		Shorten: strings.Join([]string{s.domain, "at", ans}, "/"),
	}, nil
}

func (s *shortService) Find(ctx context.Context, shorten string) (string, error) {
	log.Info().Ctx(ctx).Interface("shorten", shorten).Msg("shortService Find")
	data, err := s.linkRepo.FindOneByShort(ctx, shorten)
	if err != nil {
		return "", err
	}
	return data.OriginalURL, nil
}
