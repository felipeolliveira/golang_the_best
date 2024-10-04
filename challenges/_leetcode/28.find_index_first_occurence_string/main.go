package main

func main() {
	r := strStr("mississippi", "issipi")
	println(r)
}

func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)

	if needleLen == 0 || haystackLen == 0 || needleLen > haystackLen {
		return -1
	}

	for i, char := range haystack {
		// Checks
		// - if first char in needle is the same char in iteration: "(hello, ll) => he<l>lo == <l>l => true"
		// - if iteration + needleLen is not out of range of haystackLen
		// - if needle is equal the slice of haystack[<iteration index> : <iteration index + needleLen>] =>  ll == he<ll>o => true
		if byte(char) == needle[0] && i+needleLen <= haystackLen && needle == string(haystack[i:i+needleLen]) {
			return i
		}
	}

	return -1
}
