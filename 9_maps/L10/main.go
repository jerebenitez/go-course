package main

func getNameCounts(names []string) map[rune]map[string]int {
    
    namesMap := make(map[rune]map[string]int)

    for _, name := range names {
        runes := []rune(name)

        if _, ok := namesMap[runes[0]]; !ok {
            namesMap[runes[0]] = make(map[string]int)
        }
        
        if _, ok := namesMap[runes[0]][name]; !ok {
            namesMap[runes[0]][name] = 1
        } else {
            namesMap[runes[0]][name]++
        }
    }

    return namesMap
}

