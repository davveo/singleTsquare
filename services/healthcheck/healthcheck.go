package healthcheck

import "github.com/jinzhu/gorm"

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Close() {}

func (s *Service) HealthCheck() bool {
	rows, err := s.db.Raw("SELECT 1=1").Rows()
	defer rows.Close()

	var healthy bool
	if err == nil {
		healthy = true
	} else {
		healthy = false
	}
	return healthy
}
