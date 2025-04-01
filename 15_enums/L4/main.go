package L4

type emailStatus int

const (
    emailBounced emailStatus = iota
    emailInvalid
    emailDelivered
    emailOpened
)
