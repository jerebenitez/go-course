package main

type Membership struct {
    Type string
    MessageCharLimit int
}

type User struct {
    Membership
	Name string
}

func newUser(name string, membershipType string) User {
    var charLimit int

    switch membershipType {
        case "premium":
            charLimit = 1000
        default:
            charLimit = 100
    }

    user := User{
        Name: name,
        Membership: Membership{
            Type: membershipType,
            MessageCharLimit: charLimit,
        },
    }

    return user
}
