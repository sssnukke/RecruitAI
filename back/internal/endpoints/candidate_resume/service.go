package candidate_resume

import (
	"back/internal/config"
	"io"
)

type Service struct {
	repo *Repository
	cfg  *config.Config
}

func NewService(repo *Repository, cfg *config.Config) *Service {
	return &Service{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *Service) CheckCandidate(
	resume io.Reader,
	filename string,
	vacancy string,
) (bool, error) {

	fileID, err := UploadFile(resume, filename, s.cfg.KeyGPT)
	if err != nil {
		return false, err
	}

	return askGPT(fileID, vacancy, s.cfg.KeyGPT)
}
