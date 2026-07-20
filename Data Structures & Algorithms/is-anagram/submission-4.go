func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false;
    }

    var sMap [26]int;

    for _, char:= range s {
        sMap[char - 'a']++;
    };

    for _, char:= range t {
        if sMap[char - 'a'] == 0  {
            return false;
        }

        sMap[char - 'a']--;
    }

    return true;

}
