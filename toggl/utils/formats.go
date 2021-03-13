package utils

func MillisecondsToHours(seconds int64) float64 {
	return float64(seconds) / 1000 / 60 / 60
}
