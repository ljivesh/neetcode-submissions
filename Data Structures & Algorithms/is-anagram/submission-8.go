func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false;
    }

    var sMap [26]int;

    for i := 0; i<len(s); i++ {
        sMap[s[i]-'a']++;
        sMap[t[i]-'a']--;
    }

    return sMap == [26]int {};

}
