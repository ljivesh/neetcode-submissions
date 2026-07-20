func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false;
    }

    sMap := make(map[rune]int);

    for _, char:= range s {
        sMap[char]++;
    };

    for _, char:= range t {
        if sMap[char] == 0  {
            return false;
        }

        sMap[char]--;
    }

    return true;

}
