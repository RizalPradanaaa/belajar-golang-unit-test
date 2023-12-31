package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		panic("Result not Hello Eko")
		// t.Error("Harus hello Eko")
	}
}


// Saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil Fail(), artinya
// eksekusi unit test akan tetap dilanjutkan
func TestHelloWorldAssertion(t *testing.T)  {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result)
	fmt.Println("Dieksekusi")
}

// Sedangkan jika kita menggunakan require, jika pengecekan gagal, maka require akan memanggil
// FailNow(), artinya eksekusi unit test tidak akan dilanjutka
func TestHelloWorldRequire(t *testing.T)  {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result)
	fmt.Println("Tidak Dieksekusi")
}

func TestSkip(t *testing.T)  {
	if runtime.GOOS == "windows" {
		t.Skip("Unit test tidak bisa jalan di windows")
	}
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result)
}


// Before dan After Test
func TestMain(m *testing.M)  {
	fmt.Println("Sebelum unit test")

	m.Run()		// Eksekusi Semua Unit Test

	fmt.Println("Sesudah unit test")
}


// Sub Test
// Go-Lang mendukung fitur pembuatan
// function unit test di dalam function unit test

func TestSubTest(t *testing.T)  {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		assert.Equal(t, "Hello Eko", result)
	})

	t.Run("Rizal", func(t *testing.T) {
		result := HelloWorld("Rizal")
		assert.Equal(t, "Hello Rizal", result)
	})
}


// Benchmark adalah mekanisme menghitung kecepatan performa kode aplikasi kita
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}


func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})

	b.Run("Rizal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rizal")
		}
	})
}
