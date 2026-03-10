//go:build linux && arm64

package embedded

import _ "embed"

//go:embed dist/aidoris-web-linux-arm64
var embeddedWebBinary []byte
