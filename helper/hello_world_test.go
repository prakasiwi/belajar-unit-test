package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld(" Galih")

	if result != "Hello Galih"{
	//error
	t.Fail()
	}

	fmt.Println("TestHelloWorld Done")
}

func TestJenisKelamin(t *testing.T) {
	result := JenisKelamin(" Laki-laki")

	if result != "Jenis Kelamin = Laki-laki" {
	//error
	t.FailNow()
	}
	fmt.Println("TestJenisKelamin Done")
}

func TestUsia(t *testing.T) {
	result := 20

	if result != 20 {
	//error
	t.Error("Result must be 20")
	}
	fmt.Println("ini seperti fail")
}

func TestPanggilan(t *testing.T) {
	result := Panggilan(" Galih")

	if result != "Ini Galih" {
	//error
	t.Fatal("Result must be 'Ini Galih'")
	}
	fmt.Println("ini seperti failNow")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld(" Sobat")

	assert.Equal(t, "Hello Sobat", result, "Result must be 'Hello Sobat'")
	fmt.Println("TestHelloWord with assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld(" Broku")

	require.Equal(t, "Hello Broku", result, "Result must be 'Hello Broku'")
	fmt.Println("TestHelloWorld with require done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Bro")
	require.Equal(t, "Hello Broku", result, "Result must be 'Hello Broku'")
	fmt.Println("TestHelloWorld with require done")
}

//sub test
func TestSubTest(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		result := HelloWorld(" Dunia")

		require.Equal(t, "Hello Dunia", result, "Result must be hello test 1")
	}) 

	t.Run("Test 2", func(t *testing.T) {
		result := HelloWorld(" Teman")

		require.Equal(t, "Hello Teman", result, "Result must be hello test 2")
	})
}

//table test

func TestTableHello(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "test_satu",
			request: " salam",
			expected: "Hello salam",
		},
		{
			name: "test_dua",
			request: " satu",
			expected: "Hello satu",
		},
		{
			name: "test_tiga",
			request: " jiwa",
			expected: "Hello jiwa",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)

			require.Equal(t, test.expected, result)
		})
	}
}