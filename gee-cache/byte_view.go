package geecache

// ByteView 只读字节结构，表示缓存值
type ByteView struct {
	b []byte
}

// Len 返回字节数
func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String 将数据作为字符串返回，必要时进行复制
func (v ByteView) String() string {
	return string(v.b)
}

// 将数据复制到指定的字节数组
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
