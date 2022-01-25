// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package provider

import (
	"voltanet/wspl"
	"encoding/json"
)

// JSON contains the given interface object.
type JSON struct {
	Data interface{}
}

var jsonContentType = "application/json"

// Render (JSON) writes data with custom ContentType.
func (r JSON) Render(w *wspl.Bbo) (err error) {
	bytes, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}
	res := wspl.Response{bytes}
	err = res.Write(w)
	return err
}

// WriteContentType (JSON) writes JSON ContentType.
func (r JSON) WriteContentType(w *wspl.Bbo) {
	res := wspl.Response{map[string]string{"contentType": jsonContentType}}
	res.WriteHeader(w)
}
