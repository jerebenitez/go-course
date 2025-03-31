package main

import "slices"

func findSuggestedFriends(username string, friendships map[string][]string) []string {
    suggestedFriendsSet := make(map[string]bool)

    friends, ok := friendships[username]

    if !ok || len(friends) == 0 {
        return nil
    }

    for _, user := range friends {
        friendship, ok := friendships[user]

        if !ok {
            continue
        }

        for _, friend := range friendship {
            if friend == username || slices.Contains(friends, friend) {
                continue
            }

            suggestedFriendsSet[friend] = true
        }
    }

    return setToSlice(suggestedFriendsSet)
}

func setToSlice(set map[string]bool) []string {
    slice := make([]string, len(set))

    i := 0
    for k := range set {
        slice[i] = k
        i++
    }

    return slice
}
