package C1

func countConnections(groupSize int) int {
    connections := 0

    for i := 1; i < groupSize; i++ {
        connections += groupSize - i
    }

    return connections
}
