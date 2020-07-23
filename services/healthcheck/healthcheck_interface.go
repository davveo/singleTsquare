package healthcheck

type ServiceInterface interface {
	HealthCheck() bool
	Close()
}
