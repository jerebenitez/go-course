package L10

func sum(nums ...int) int {
    sum := 0

    for i, q := 0, len(nums); i < q; i++ {
        sum += nums[i]
    }

    return sum
}
