package main

func reverse(x int) int {
	reversedInt := 0
	for x != 0 {
		reversedInt = reversedInt*10 + x%10
		x /= 10
        if reversedInt != int(int32(reversedInt)){
            return 0
        }
	}

	return reversedInt
}
