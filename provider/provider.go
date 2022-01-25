package provider

import (
	"voltanet/wspl"
)

// Render interface is to be implemented by JSON, XML, HTML, YAML and so on.
type Render interface {
	// Render writes data with custom ContentType.
	Render(w *wspl.Bbo) error
	// WriteContentType writes custom ContentType.gi
	WriteContentType(w *wspl.Bbo)
}
