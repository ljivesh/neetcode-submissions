import "slices"

func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
        return [][]string{}
    }

    var res [][]string

    m := make(map[string]int)

    for i := range strs {
        b := []byte(strs[i])

        slices.Sort(b)

        sorted := string(b)

        v, ok := m[sorted]
        if !ok {
            m[sorted] = len(m)

            res = append(res, []string{strs[i]})

            continue
        }


        res[v] = append(res[v], strs[i])
    }

    return res
}
