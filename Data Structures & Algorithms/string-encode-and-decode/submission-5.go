type Solution struct{}

func (s *Solution) Encode(strs []string) string {

    //Start of Text - STX
    encodedString:= string(2);

    for idx, str:= range strs {
        
        // if idx == 0 {
        //     encodedString = str;
        //     continue;
        // }

        //if the string is empty we send a substitute character
        if str == "" {
            
            if idx == 0 {
                encodedString = encodedString + string(26)
            } else {
                encodedString = encodedString + string(31) + string(26)
            }
            
            continue;
        }
        // append next string delimeter by 'US' unit seperator character

        if idx == 0 {
            encodedString = encodedString +  str;
        } else {
            encodedString = encodedString + string(31) + str;
        }

    }

    //end the payload with ETX - End of text character 
    encodedString = encodedString + string(3);

    return encodedString;

}

func (s *Solution) Decode(encoded string) []string {

    // Empty String
    if len(encoded) == 0 {
        return []string {};
    }

    //If STX not present then invalid payload
    if encoded[0] != 2 {
        return []string {};
    } 

    //Removing STX
    encoded = encoded[1:];
     
    //Collect strings in an array 
    decodedStrings := make([]string, 0);

    //Initialize two pointers
    i:=0;
    j:=0;

    //Iterate the encoded data till we reach the ETX character 
    for encoded[i] != 3 {

        //rest the 2nd pointer to new start
        j = i;

        //iterate till we get a seperator
        for (encoded[j] != 3) && (encoded[j] != 31) {
            j++
        }
        
        //get slice (j is currently at ETX so excluding that)
        str := encoded[i:j];

        //Replacing SUB back to ""
        if str == string(26) {
            str = "";
        }

        //appending the slice to data
        decodedStrings = append(decodedStrings, str);


        //Reached ETX
        if encoded[j]==3 {
            break;
        }

        //moving to the next data slice
        i = j+1;

    }


    return decodedStrings;

}
