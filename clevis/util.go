package clevis

func expandBuffer(oldBuffer []byte, newLen int) []byte {
	oldLen := len(oldBuffer)
	if oldLen < newLen {
		newBuffer := make([]byte, newLen-oldLen, newLen)
		newBuffer = append(newBuffer, oldBuffer...)
		return newBuffer
	} else {
		return oldBuffer
	}
}
