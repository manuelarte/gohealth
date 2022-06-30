package gohealth

type PingChecker struct{}

func (checker *PingChecker) CheckHealth() HealthCheckResult {
	result := HealthCheckResult{Service: "ping"}
	result.Status = UP
	return result
}
