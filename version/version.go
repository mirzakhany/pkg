package version

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

// Settings version settings
type Settings struct {
	CopyRightYear string
	LongHash      string
	ShortHash     string
	CommitDate    string
	CommitCount   string
	BuildDate     string
	Version       string
	ServiceName   string
	CompanyName   string
}

var versionSettings Settings

// SetupVersion for setup version string.
func SetupVersion(settings Settings) {
	versionSettings = settings
}

// GetVersion return version
func GetVersion() string {
	return versionSettings.Version
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
		versionSettings.Version,
		runtime.Compiler,
		runtime.Version(),
		versionSettings.CopyRightYear,
		versionSettings.CompanyName)
	fmt.Println()
	fmt.Println("Commit Hash:", versionSettings.LongHash)
	fmt.Println("Commit ShortHash:", versionSettings.ShortHash)
	fmt.Println("Commit Count:", versionSettings.CommitCount)
	fmt.Println("Commit Date:", versionSettings.CommitDate)
	fmt.Println("Build Date:", versionSettings.BuildDate)
	drawLine(70)
}

// HeaderVersionMiddleware : add version on header.
func HeaderVersionMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		k := fmt.Sprintf("X-%s-VERSION", strings.ToUpper(versionSettings.ServiceName))
		c.Header(k, versionSettings.Version)
		c.Next()
	}
}
