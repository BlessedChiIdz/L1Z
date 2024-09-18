package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func powBig(a, n int) *big.Int {
	tmp := big.NewInt(int64(a))
	res := big.NewInt(1)
	for n > 0 {
		temp := new(big.Int)
		if n%2 == 1 {
			temp.Mul(res, tmp)
			res = temp
		}
		temp = new(big.Int)
		temp.Mul(tmp, tmp)
		tmp = temp
		n /= 2
	}
	return res
}

func euklidF(a, b int) []int {
	u := []int{a, 1, 0}
	v := []int{b, 0, 1}
	t := []int{0, 0, 0}
	for v[0] != 0 {
		q := u[0] / v[0]
		t[0] = u[0] % v[0]
		t[1] = u[1] - q*v[1]
		t[2] = u[2] - q*v[2]
		copy(u, v)
		copy(v, t)
	}
	return u
}

func powmod(a, b, p int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % p
		}
		a = (a * a) % p
		b >>= 1
	}
	return res
}

func generator(p int) int {
	fact := []int{}
	phi := p - 1
	n := phi

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fact = append(fact, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		fact = append(fact, n)
	}

	for res := 2; res < p; res++ {
		ok := true
		for _, f := range fact {
			if powmod(res, phi/f, p) == 1 {
				ok = false
				break
			}
		}
		if ok {
			return res
		}
	}
	return -1
}

func diffieHellmanGlobal() (p int, g int) {
	reader := rand.New(rand.NewSource(time.Now().UnixNano()))
	p = reader.Int()
	//g = (p - 1) / 2
	//p = 46
	//g = 23
	g = generator(p)
	return
}

func diffieHellmanLocal(p, g int) (int, int) {
	reader := rand.New(rand.NewSource(2))
	sKey := reader.Int() / 100000000000000000
	//sKey := 22
	pKey := int(math.Pow(float64(g), float64(sKey))) % p
	return pKey, sKey
}

func diffieHellmanServer(p, g int) (int, int) {
	reader := rand.New(rand.NewSource(1))
	sKey := reader.Int() / 100000000000000000
	//sKey := 55
	pKey := int(math.Pow(float64(g), float64(sKey))) % p
	return pKey, sKey
}

func diffieHellmanLocalSec(pKey, sKey, p int) int {
	key := int(math.Pow(float64(pKey), float64(sKey))) % p
	return key
}

func diffieHellmanServerSec(pKey, sKey, p int) int {
	key := int(math.Pow(float64(pKey), float64(sKey))) % p
	return key
}

// Функция для вычисления обратного элемента по модулю p
func modInverse(a, p int) int64 {
	m := int64(p)
	var m0, x0, x1 int64 = m, 0, 1
	a0 := int64(a)
	if m == 1 {
		return 0
	}
	for a0 > 1 {
		// q - частное
		q := int64(a0) / m
		m0, m = m, int64(a0)%m
		x0, x1 = x1-q*x0, x0
	}
	if x1 < 0 {
		x1 += m0
	}
	return x1
}

func main() {
	powRes := powBig(1234567886666, 2)
	euklidRes := euklidF(28, 19)
	fmt.Println(powRes)
	fmt.Println(euklidRes)
	p, g := diffieHellmanGlobal()
	//fmt.Printf("p=%d,g=%d\n", p, g)
	pLocal, sLocal := diffieHellmanLocal(p, g)
	fmt.Printf("pLocal = %d, sLocal = %d\n", pLocal, sLocal)
	pServer, sServer := diffieHellmanServer(p, g)
	fmt.Printf("pserv = %d, sServe = %d\n", pServer, sServer)
	keyL := diffieHellmanLocalSec(pLocal, sLocal, p)
	keyS := diffieHellmanServerSec(pServer, sServer, p)
	fmt.Println(keyL)
	fmt.Println(keyS)
}
