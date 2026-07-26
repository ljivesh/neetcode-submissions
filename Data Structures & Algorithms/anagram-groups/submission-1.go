func areValidAnagrams(a string, b string) bool {

	if len(a) != len(b) {
		return false;
	}

	var charMap [26]int;

	for i:=0; i<len(a); i++ {
		charMap[a[i] - 'a']++;
		charMap[b[i] - 'a']--;
	}

	return charMap == [26]int {};

}

func groupAnagrams(strs []string) [][]string {

	var arrayGroups = make([][]string, 0);

	fmt.Println(arrayGroups);
	

	for i:=0; i<len(strs); i++ {

		var isAdded = false;

		for idx, array := range arrayGroups {

			if len(array) == 0 {
				break;
			}

			if areValidAnagrams(array[0], strs[i]) {
				array = append(array, strs[i]);

				arrayGroups[idx] = array;
				isAdded = true;
				break;
			}

		}

		if !isAdded {
			var newArrayGroup = make([]string, 0);
			newArrayGroup = append(newArrayGroup, strs[i])

			arrayGroups = append(arrayGroups, newArrayGroup);
		}


	} 

	return arrayGroups

}
