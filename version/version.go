package version

import (
	"fmt"
	"runtime"

	"strings"

	"github.com/gin-gonic/gin"
)

var (
	copyRightYear string
	hash          string
	short         string
	date          string
	count         string
	build         string
	version       string
	serviceName   string
	companyName   string
)

// SetupVersion for setup version string.
func SetupVersion(ver string, srvName string, cmpName string, year string) {

	if ver != "" {
		version = ver
	}

	serviceName = srvName
	companyName = cmpName
	copyRightYear = year
}

// GetVersion for get current version.
func GetVersion() string {
	return version
}

func drawLine(w int) {
	for i := 0; i < w; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}

// PrintServiceVersion provide print server engine
func PrintServiceVersion() {
	drawLine(70)
	fmt.Printf(`Version %s, Compiler: %s %s, Copyright (C) %s %s, Inc.`,
		version,
		runtime.Compiler,
		runtime.Version(),
		copyRightYear,
		companyName)
	fmt.Println()
	fmt.Println("Commit Hash:", hash)
	fmt.Println("Commit ShortHash:", short)
	fmt.Println("Commit Count:", count)
	fmt.Println("Commit Date:", date)
	fmt.Println("Build Date:", build)
	drawLine(70)
}

// HeaderVersionMiddleware : add version on header.
func HeaderVersionMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		k := fmt.Sprintf("X-%s-VERSION", strings.ToUpper(serviceName))
		c.Header(k, version)
		c.Next()
	}
}
