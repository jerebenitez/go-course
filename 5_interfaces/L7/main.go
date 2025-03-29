package main

import "fmt"

func (e email) cost() int {
    if !e.isSubscribed {
        return 5 * len(e.body)
    } else {
        return 2 * len(e.body)
    }
}

func (e email) format() string {
    status := "Not Subscribed"

    if e.isSubscribed {
        status = "Subscribed"
    }

    return fmt.Sprintf("'%s' | %s", e.body, status)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
