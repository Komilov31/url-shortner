package handler

import (
	"github.com/Komilov31/url-shortener/internal/dto"
	"github.com/Komilov31/url-shortener/internal/model"
)

type ShortnerServcie interface {
	GetAnalytics(string) ([]dto.RedirectInfo, error)
	GetUrlByShort(string, model.RedirectInfo) (*model.Url, error)
	CreateShortUrl(model.Url) (*model.Url, error)
	AggregateByUserAgent() ([]dto.UserAgentDTO, error)
	AggregateByDate() ([]dto.DateDTO, error)
	AggregateByMonth() ([]dto.MonthDTO, error)
}

type Handler struct {
	service ShortnerServcie
}

func New(service ShortnerServcie) *Handler {
	return &Handler{
		service: service,
	}
}
