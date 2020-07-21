package services

import "github.com/jinzhu/gorm"

type HealthCheckService struct {
	db *gorm.DB
}

func (h *HealthCheckService) HealthCheck() bool {
	rows, err := h.db.Raw("SELECT 1=1").Rows()
	defer rows.Close()

	var healthy bool
	if err == nil {
		healthy = true
	} else {
		healthy = false
	}
	return healthy
}
