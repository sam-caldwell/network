package namespace

import (
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestTypes(t *testing.T) {
	t.Run("size validation", func(t *testing.T) {
		const expectedSizeInBytes = 8
		if sz := int(unsafe.Sizeof(Types(0))); sz != expectedSizeInBytes {
			t.Fatal("unexpected size")
		}
	})
	t.Run("value check", func(t *testing.T) {
		type TestData struct {
			actual Types
			expect Types
		}
		testData := []TestData{
			{actual: JoinAnyNamespace, expect: 0},
			{actual: CloneNewcgroup, expect: unix.CLONE_NEWCGROUP},
			{actual: CloneNewipc, expect: unix.CLONE_NEWIPC},
			{actual: CloneNewnet, expect: unix.CLONE_NEWNET},
			{actual: CloneNewns, expect: unix.CLONE_NEWNS},
			{actual: CloneNewpid, expect: unix.CLONE_NEWPID},
			{actual: CloneNewtime, expect: unix.CLONE_NEWTIME},
			{actual: CloneNewuser, expect: unix.CLONE_NEWUSER},
			{actual: CloneNewuts, expect: unix.CLONE_NEWUTS},
		}
		for _, v := range testData {
			if v.actual != v.expect {
				t.Fatalf("unexpected type %v; got %v", v.actual, v.expect)
			}
		}
	})
}
