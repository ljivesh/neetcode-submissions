type Solution struct{}

func (s *Solution) Encode(strs []string) string {

    encoded:= ""

    for _, str := range strs {

        size := len(str)

        //append data in {size}#{string} format
        encoded = encoded + strconv.Itoa(size) + "#" + str

    }

    return encoded

}

func (s *Solution) Decode(encoded string) []string {

   
    decoded := make([]string, 0)

    // Initalize two pointers
    i := 0
    j := 0

    for i < len(encoded) {

        j = i

        //iterate till number end signal
        for encoded[j] != byte('#') {
            j++
        }

        //get size of str element
        num := encoded[i:j]

        //convert num string to int
        convertedNum, err := strconv.Atoi(num)

        if err != nil {
            fmt.Println(err);
            return []string {}
        }

        //set pointers to str element
        i = j+1
        j = i+convertedNum

        str := encoded[i:j]

        decoded = append(decoded, str)

        i = j
        
    }

    return decoded
}
