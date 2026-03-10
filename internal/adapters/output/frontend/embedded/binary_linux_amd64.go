//go:build linux && amd64

package embedded

import _ "embed"

//go:embed dist/aidoris-web-linux-x64
var embeddedWebBinary []byte
