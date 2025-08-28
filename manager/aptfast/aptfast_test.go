// apt/apt_test.go
package aptfast_test

import (
	"testing"

	"github.com/awalsh128/syspkg/manager/apt"
)

func TestAptPackageManager(t *testing.T) {
	// Implement test cases for AptPackageManager
	aptManager := &apt.PackageManager{}
	if aptManager.IsAvailable() {
		t.Log("AptPackageManager is available")
	} else {
		t.Fatal("AptPackageManager is not available")
	}
}
