package L6

func getMessageCosts(messages []string) []float64 {
    costs := make([]float64, len(messages))

    for i, qty := 0, len(messages); i < qty; i++ {
        costs[i] = float64(len(messages[i])) * 0.01
    }

    return costs
}
