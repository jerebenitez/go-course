package C2

import "time"

func processMessages(messages []string) []string {
    if messages == nil || len(messages) == 0 {
        return []string{}
    }

    ch := make(chan string)
    result := make([]string, len(messages))

    for i, msg := range messages {
        go func() {
            ch <- process(msg)
        }()

        result[i] = <- ch
    }

    return result
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
