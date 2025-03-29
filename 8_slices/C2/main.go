package main

import "unicode"

func isValidPassword(password string) bool {
    if len(password) < 5 || len(password) > 12 {
        return false
    }

    hasNumber := false
    hasUppercase := false

    for _, char := range password {
        if unicode.IsNumber(char) {
            hasNumber = true
        } else if unicode.IsUpper(char) {
            hasUppercase = true
        }

        if hasNumber && hasUppercase {
            break
        }
    }

    return hasNumber && hasUppercase
}

