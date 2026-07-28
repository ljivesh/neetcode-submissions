func productExceptSelf(nums []int) []int {

    output := make([]int, 0)

    zeroCount := 0

    completeProduct := 1

    for _, elem:= range nums {

        if elem == 0 {
            zeroCount++
            continue
        }

        completeProduct = completeProduct * elem
    }

    atleastTwoZeros := zeroCount > 1

    for _, elem:= range nums {

        if atleastTwoZeros {
            output = append(output, 0)
            continue
        }

        if elem == 0 {
            output = append(output, completeProduct)
            continue
        }

        if zeroCount > 0 {
            output = append(output, 0)
            continue
        }

        output = append(output, completeProduct/elem)
    }

    return output
}
