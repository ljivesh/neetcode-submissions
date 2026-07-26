func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)

    for _, s := range strs {
        var freq [26]int
        for i := 0; i < len(s); i++ {
            freq[s[i]-'a']++
        }
        groups[freq] = append(groups[freq], s)
    }

    ans := make([][]string, 0, len(groups))
    for _, g := range groups {
        ans = append(ans, g)
    }

    return ans
}