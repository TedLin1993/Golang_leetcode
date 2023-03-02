package leetcode

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	slow, fast := 0, 1
	count := 1
	for fast < len(chars) {
		if chars[fast] != chars[slow] {
			if count > 1 {
				byteArr := int2ByteArray(count)
				for i := 0; i < len(byteArr); i++ {
					slow++
					chars[slow] = byteArr[i]
				}
			}
			slow++
			chars[slow] = chars[fast]
			count = 1
		} else {
			count++
		}
		fast++
	}
	if count > 1 {
		byteArr := int2ByteArray(count)
		for i := 0; i < len(byteArr); i++ {
			slow++
			chars[slow] = byteArr[i]
		}
	}
	return slow + 1
}

func int2ByteArray(a int) []byte {
	res := []byte{}
	for a > 0 {
		remainder := byte('0' + a%10)
		res = append([]byte{remainder}, res...)
		a /= 10
	}
	return res
}
