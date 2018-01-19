package secret

const testVersion = 2

// Handshake returns the secret handshake
func Handshake(secret uint) []string {
	var retval = []string{}
	if secret&0x1 != 0 {
		retval = append(retval, "wink")
	}
	if secret&0x2 != 0 {
		retval = append(retval, "double blink")
	}
	if secret&0x4 != 0 {
		retval = append(retval, "close your eyes")
	}
	if secret&0x8 != 0 {
		retval = append(retval, "jump")
	}
	if secret&0x10 != 0 {
		retval = reverse(retval)
	}
	return retval
}

func reverse(s []string) []string {
	var retval []string
	for i := len(s) - 1; i >= 0; i-- {
		retval = append(retval, s[i])
	}
	return retval
}
