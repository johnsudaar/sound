package utils

func F64ToF32(buf1 []float64, buf2 []float32) {
	for i := 0; i < len(buf1); i++ {
		buf2[i] = float32(buf1[i])
	}
}
