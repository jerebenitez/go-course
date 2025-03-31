package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(anal *Analytics, msg Message) {
    anal.MessagesTotal++

    switch msg.Success {
        case true:
            anal.MessagesSucceeded++
        case false:
            anal.MessagesFailed++
    }
}
