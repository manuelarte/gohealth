package health

type PingChecker struct{}

func (checker *PingChecker) CheckHealth() (result HealthCheckResult) {
	result.Service = "ping"
	result.Status = UP
	return
}
