//go:build windows && amd64

package embeddedadapter

import _ "embed"

//go:embed dist/aidoris-web-windows-x64.exe
var embeddedWebBinary []byte
