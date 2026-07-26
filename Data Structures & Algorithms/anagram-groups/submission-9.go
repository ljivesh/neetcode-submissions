func getAnagramSignature(a string) [26]int {

	var charMap [26]int;

	for i:=0; i<len(a); i++ {
		charMap[a[i] - 'a']++;
	}

	return charMap;

}

func groupAnagrams(strs []string) [][]string {

	var anagramMap = make(map[[26]int] []string);

	for i:=0; i<len(strs); i++ {
		
		var anagramSignature = getAnagramSignature(strs[i]);

		var current =  anagramMap[anagramSignature]
		anagramMap[anagramSignature] = append(current, strs[i]); 
		
	} 

	fmt.Println(anagramMap);

	var arrayMap = make([][]string, 0);

	for _, values := range anagramMap {
		arrayMap = append(arrayMap, values);
	}

	return arrayMap;

}
