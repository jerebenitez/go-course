package C1

import "fmt"

type Formatter interface {
    Format() string
}

type PlainText struct {
    message string
}

func (p PlainText) Format() string {
    return p.message
}

type Bold struct {
    message string
}

func (p Bold) Format() string {
    return fmt.Sprintf("**%s**", p.message)
}

type Code struct {
    message string
}

func (p Code) Format() string {
    return fmt.Sprintf("`%s`", p.message)
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
