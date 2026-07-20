func isAnagram(s string, t string) bool {

    var countS = len(s);
    var countT = len(t);

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

    return countS == countT;

}
