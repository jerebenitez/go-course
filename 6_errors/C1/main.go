package C1

import "errors"

func validateStatus(status string) error {
    if len(status) < 1 {
        return errors.New("status cannot be empty")
    }

    if len(status) > 140 {
        return errors.New("status exceeds 140 characters")
    }

    return nil
}
