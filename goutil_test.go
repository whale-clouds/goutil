package goutil_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whale-clouds/goutil"
)

func TestFuncName(t *testing.T) {
	name := goutil.FuncName(goutil.PkgName)
	assert.Equal(t, "github.com/whale-clouds/goutil.PkgName", name)

	name = goutil.FuncName(goutil.PanicIfErr)
	assert.Equal(t, "github.com/whale-clouds/goutil.PanicIfErr", name)
}

func TestPkgName(t *testing.T) {
	name := goutil.PkgName(goutil.FuncName(goutil.PanicIfErr))
	assert.Equal(t, "github.com/whale-clouds/goutil", name)
}

func TestPanicIfErr(t *testing.T) {
	goutil.PanicIfErr(nil)
}

func TestPanicf(t *testing.T) {
	assert.Panics(t, func() {
		goutil.Panicf("hi %s", "inhere")
	})
}

func TestGetCallStacks(t *testing.T) {
	msg := goutil.GetCallStacks(false)
	fmt.Println(string(msg))

	fmt.Println("-------------full stacks-------------")
	msg = goutil.GetCallStacks(true)
	fmt.Println(string(msg))
}
