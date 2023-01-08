package staff

type FixedPriceJob struct {
	description string
	fixedPrice  float64
}
type HourlyJob struct {
	description string
	hourlyRate  float64
	numberHours int
}
type JobInterface interface {
	Cost() float64
	GetDescription() string
}

func (job FixedPriceJob) Cost() float64 {
	return job.fixedPrice
}
func (job FixedPriceJob) GetDescription() string {
	return job.description
}

func (job HourlyJob) Cost() float64 {
	return job.hourlyRate * float64(job.numberHours)
}
func (job HourlyJob) GetDescription() string {
	return job.description
}

func TotalJobCost(jobs []JobInterface) float64 {
	result := 0.0
	for _, job := range jobs {
		result += job.Cost()
	}
	return result
}
