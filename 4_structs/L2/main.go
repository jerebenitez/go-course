package L2

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	return !(isEmpty(mToSend.recipient) || isEmpty(mToSend.sender))
}

func isEmpty(user user) bool {
    return user.name == "" || user.number == 0
}
