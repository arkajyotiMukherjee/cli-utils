package chaos

import (
	"fmt"

	"github.com/mayadata-io/cli-utils/pkg/constants"
)

// GetMode gets mode of agent installation as input
func GetMode() string {
	var mode int
	fmt.Println("\n🔌 Installation Modes:\n1. Cluster\n2. Namespace")
	fmt.Print("\n👉 Select Mode [", constants.DefaultMode, "]: ")
	fmt.Scanln(&mode)
	if mode == 0 {
		return "namespace"
	}
	for mode < 1 || mode > 2 {
		fmt.Println("🚫 Invalid mode. Please enter the correct mode")
		fmt.Print("👉 Select Mode [", constants.DefaultMode, "]: ")
		fmt.Scanln(&mode)
	}
	if mode == 1 {
		return "cluster"
	}
	return constants.DefaultMode
}
