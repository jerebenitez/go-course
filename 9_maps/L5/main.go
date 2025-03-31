package L5

func getCounts(messagedUsers []string, validUsers map[string]int) {
    for _, user := range messagedUsers {
        if _, ok := validUsers[user]; ok {
            validUsers[user]++
        }
    }
}

