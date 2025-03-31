package C1

import (
	"strings"
)

func countDistinctWords(messages []string) int {
    words := make(map[string]bool)

    for _, msg := range messages {
        if msg == "" {
            continue
        }

        for _, word := range strings.Split(msg, " ") {
            words[strings.ToLower(word)] = true
        } 
    }

    return len(words)
}
