package helper

import (
	"fmt"
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
