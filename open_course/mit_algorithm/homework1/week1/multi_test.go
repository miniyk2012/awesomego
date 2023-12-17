package week1

import (
	"bytes"
	"math/big"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	a      string
	b      string
	result string
}

var testCases = []testCase{
	{"1", "0", "0"},
	{"0", "0", "0"},
	{"112", "12", "1344"},
	{"2", "12", "24"},
	{"1233998", "870", "1073578260"},
	{"1233998", "870", "1073578260"},
	{"3141592653589793238462643383279502884197169399375105820974944592",
		"2718281828459045235360287471352662497757247093699959574966967627",
		"8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184"},
	{"0", "2121233", "0"},
}

type multiplyer interface {
	Multiply(a, b []byte) []byte
}

func TestSchoolMethod_Multiply(t *testing.T) {
	var mul multiplyer = gradeSchool{}
	testMul(mul, "gradeSchool", t)
	mul = karatsuba{}
	testMul(mul, "karatsuba", t)
}

func testMul(mul multiplyer, desc string, t *testing.T) {
	for i, tc := range testCases {
		t.Run(desc+strconv.Itoa(i), func(t *testing.T) {
			result := mul.Multiply([]byte(tc.a), []byte(tc.b))
			a, _ := new(big.Int).SetString(tc.a, 10)
			b, _ := new(big.Int).SetString(tc.b, 10)
			a.Mul(a, b)
			assert.Equal(t, tc.result, a.String())
			assert.Equal(t, a.String(), string(result))
		})
	}
	t.Run(desc+"(i*j)", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			for j := 0; j < 1000; j++ {
				result := mul.Multiply([]byte(strconv.Itoa(i)), []byte(strconv.Itoa(j)))
				assert.Equal(t, strconv.Itoa(i*j), string(result))
			}
		}
	})
	mulx := generateRandomIntBytes(1024)
	muly := generateRandomIntBytes(512)
	t.Run(desc+"BigMul", func(t *testing.T) {
		result := mul.Multiply(mulx, muly)
		a, _ := new(big.Int).SetString(string(mulx), 10)
		b, _ := new(big.Int).SetString(string(muly), 10)
		assert.Equal(t, a.Mul(a, b).String(), string(result))
	})
}

func TestMulNum(t *testing.T) {
	aline := []uint8{1, 2, 3, 0, 6}
	c := mulNum(aline, 4)
	assert.Equal(t, []uint8{4, 9, 2, 2, 4}, c)
	c = mulNum(aline, 0)
	assert.Equal(t, []uint8{0}, c)
	c = mulNum(aline, 9)
	assert.Equal(t, []uint8{1, 1, 0, 7, 5, 4}, c)
}

func TestAddLine(t *testing.T) {
	total := []uint8{6, 9, 3, 4}
	newLine := []uint8{4, 3, 8}
	c := addLine(total, newLine, 1)
	assert.Equal(t, []uint8{1, 1, 3, 1, 4}, c)
	c = addLine(total, newLine, 0)
	assert.Equal(t, []uint8{7, 3, 7, 2}, c)
	c = addLine(total, newLine, 2)
	assert.Equal(t, []uint8{5, 0, 7, 3, 4}, c)
	c = addLine(nil, newLine, 1)
	assert.Equal(t, []uint8{4, 3, 8, 0}, c)
	c = addLine(total, nil, 1)
	assert.Equal(t, []uint8{6, 9, 3, 4}, c)

}

func TestBigInt(t *testing.T) {
	a := big.NewInt(120)
	b := a.Append([]byte{'a'}, 10)
	t.Logf("%s", b)
	t.Logf("total=%s", strconv.Itoa(1))
	t.Logf("x=%c", byte(49)) // 1
	t.Logf("x=%c", byte(65)) // A
	t.Logf("x=%c", 5+'0')    // 5
}

func TestPaddingZeros(t *testing.T) {
	a := []byte("120")
	b := []byte("120")
	c, d := paddingTwoPower(a, b)
	assert.Equal(t, "0120", string(c))
	assert.Equal(t, "0120", string(d))

	a = []byte("0")
	b = []byte("120")
	c, d = paddingTwoPower(a, b)
	assert.Equal(t, "0000", string(c))
	assert.Equal(t, "0120", string(d))

	a = []byte("12299988887")
	b = []byte("12")
	c, d = paddingTwoPower(a, b)
	assert.Equal(t, "0000012299988887", string(c))
	assert.Equal(t, "0000000000000012", string(d))
}

func TestCheckPowerOfTwo(t *testing.T) {
	assert.True(t, CheckPowerOfTwo(1))
	assert.True(t, CheckPowerOfTwo(4))
	assert.True(t, CheckPowerOfTwo(32))
	assert.True(t, CheckPowerOfTwo(1024))
	assert.False(t, CheckPowerOfTwo(0))
	for i := 513; i < 1024; i++ {
		assert.False(t, CheckPowerOfTwo(i))
	}
}

func TestKaratsubaAuxiliaries(t *testing.T) {
	a := shortMul([]byte("12"), []byte("12"))
	assert.Equal(t, "144", string(a))
	a = shortMul([]byte("12"), []byte("0"))
	assert.Equal(t, "0", string(a))

	b := add([]byte("19"), []byte("123"))
	assert.Equal(t, "142", string(b))
	b = add([]byte("0"), []byte("0"))
	assert.Equal(t, "0", string(b))
	b = add([]byte("120"), []byte("0"))
	assert.Equal(t, "120", string(b))

	c := sub([]byte("100"), []byte("99"))
	assert.Equal(t, "1", string(c))
	c = sub([]byte("120"), []byte("102"))
	assert.Equal(t, "18", string(c))
	c = sub([]byte("0"), []byte("0"))
	assert.Equal(t, "0", string(c))
}

const (
	numericChars = "0123456789" // 数字字符
)

func generateRandomIntBytes(length int) []byte {
	// 设置随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var bf bytes.Buffer

	// 生成首位字符，排除 0
	randomIndex := r.Intn(9) + 1 // 生成 1-9 的随机数
	randomChar := numericChars[randomIndex]
	bf.WriteByte(randomChar)

	// 生成剩余字符
	for i := 1; i < length; i++ {
		randomIndex = r.Intn(len(numericChars))
		randomChar = numericChars[randomIndex]
		bf.WriteByte(randomChar)
	}

	return bf.Bytes()
}

// go test -bench .
func BenchmarkGradeSchool_Multiply(b *testing.B) {
	for l := 1; l <= 1024*8; l *= 2 {
		b.Run(strconv.Itoa(l), func(b *testing.B) {
			mul := gradeSchool{}
			benchmarkMul(b, mul, l)
		})
	}
}

func benchmarkMul(b *testing.B, mul multiplyer, length int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		mulx := generateRandomIntBytes(length)
		muly := generateRandomIntBytes(length)
		b.StartTimer()
		mul.Multiply(mulx, muly)
	}
}

// go test -bench .
func BenchmarkKaratsuba_Multiply(b *testing.B) {
	for l := 1; l <= 1024*8; l *= 2 {
		b.Run(strconv.Itoa(l), func(b *testing.B) {
			mul := karatsuba{}
			benchmarkMul(b, mul, l)
		})
	}
}
