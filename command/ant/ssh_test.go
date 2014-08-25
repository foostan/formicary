package ant

import(
	"testing"

	. "."
)

func TestSsh(t *testing.T) {
	Ssh("vagrant", "vagrant", "192.168.33.10:22", "ip a")
}
