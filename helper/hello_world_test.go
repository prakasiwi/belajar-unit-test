package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
	result := HelloWorld("Bro")

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