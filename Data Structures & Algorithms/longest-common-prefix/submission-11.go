func longestCommonPrefix(strs []string) string {

	if(len(strs) == 1) {
		return strs[0];
	}
    
	var mostCommonPrefix = make(map[string]int);

	for i:=0; i<len(strs); i++ {

		for j:=0; j<=len(strs[i]); j++ {
			var prefix = strs[i][0:j];

			mostCommonPrefix[prefix]++;
			
		}

	}


	fmt.Println(mostCommonPrefix);

	var mostCommonPrefixString = "";
	var mostCommonPrefixStringCount = 1;

	for elem, count := range mostCommonPrefix {
		if (count > mostCommonPrefixStringCount) {

			fmt.Println(count, mostCommonPrefixStringCount, elem, mostCommonPrefixString)	

			mostCommonPrefixStringCount = count;
			mostCommonPrefixString = elem;
		}

		if (count == mostCommonPrefixStringCount) && (len(elem) > len(mostCommonPrefixString)) {
			mostCommonPrefixString = elem;
		}

	}

	if(mostCommonPrefixStringCount == 1) {
		return "";
	}

	return mostCommonPrefixString; 
	
}	

