//go:build darwin && arm64

package embedded

import _ "embed"

//go:embed dist/aidoris-web-darwin-arm64
var embeddedWebBinary []byte
