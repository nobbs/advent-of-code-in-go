package util

func ChineseRemainderTheorem(num []int, rem []int) int {
	N := 1
	for _, n := range num {
		N *= n
	}

	x := 0
	for i := 0; i < len(num); i++ {
		p := N / num[i]
		x += rem[i] * InverseModulo(p, num[i]) * p
	}

	return (x%N + N) % N
}

func ExtendedGCD(a, b int) (int, int, int) {
	x, y, u, v := 0, 1, 1, 0

	for a != 0 {
		q, r := b/a, b%a
		m, n := x-u*q, y-v*q
		b, a, x, y, u, v = a, r, u, v, m, n
	}

	gcd := b
	return gcd, x, y
}

func InverseModulo(a, n int) int {
	_, x, _ := ExtendedGCD(a, n)
	return x
}
