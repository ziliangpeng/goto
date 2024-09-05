func scoreOfString(s string) int {
    first := s[:len(s)-1]
    second := s[1:]
    total := 0
    for i := range len(first) {
        c1 := int(first[i])
        c2 := int(second[i])
        if c1 > c2 {
            total += c1 - c2
        } else {
            total += c2 - c1
        }
    }
    return total
}