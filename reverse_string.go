package reverse_string

func ReverseString(input string) (output string) {
	byte_str := []rune(input)
   	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
    	byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
   	}
   	return string(byte_str)
}
