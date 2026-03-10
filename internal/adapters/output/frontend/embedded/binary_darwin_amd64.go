//go:build darwin && amd64

package embedded

import _ "embed"

//go:embed dist/aidoris-web-darwin-x64
var embeddedWebBinary []byte
