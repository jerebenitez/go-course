package L11

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
    dayCosts := []float64{}

    for _, cost := range costs {
        if cost.day == day {
            dayCosts = append(dayCosts, cost.value)
        }
    }

    return dayCosts
}
