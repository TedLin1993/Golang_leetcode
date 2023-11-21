package leetcode

func maximumXorProduct(a int64, b int64, n int) int {
	mod := int64(1e9 + 7)
	for i := n - 1; i >= 0; i-- {
		mask := int64(1 << i)
		aBit := a >> i % 2
		bBit := b >> i % 2
		if (aBit == 0 && bBit == 0) || (aBit == 0 && a>>i < b>>i) || (bBit == 0 && a>>i > b>>i) {
			a ^= mask
			b ^= mask
		}
	}
	return int((a % mod) * (b % mod) % mod)
}
