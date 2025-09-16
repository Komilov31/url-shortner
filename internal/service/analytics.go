package service

import "github.com/Komilov31/url-shortener/internal/dto"

func (s *Service) AggregateByUserAgent() ([]dto.UserAgentDTO, error) {
	return s.storage.AggregateByUserAgent()
}

func (s *Service) AggregateByDate() ([]dto.DateDTO, error) {
	return s.storage.AggregateByDate()
}

func (s *Service) AggregateByMonth() ([]dto.MonthDTO, error) {
	return s.storage.AggregateByMonth()
}
