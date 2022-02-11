// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package provider

import (
	"github.com/golang/protobuf/proto"
	"github.com/voltaspace/volta-net/wspl"
)

// ProtoBuf contains the given interface object.
type ProtoBuf struct {
	Data interface{}
}

var protobufContentType = "application/x-protobuf"

// Render (ProtoBuf) marshals the given interface object and writes data with custom ContentType.
func (r ProtoBuf) Render(w *wspl.Bbo) error {

	bytes, err := proto.Marshal(r.Data.(proto.Message))
	if err != nil {
		return err
	}
	res := wspl.Response{bytes}
	err = res.Write(w)
	return err
}

// WriteContentType (ProtoBuf) writes ProtoBuf ContentType.
func (r ProtoBuf) WriteContentType(w *wspl.Bbo) {
	res := wspl.Response{map[string]string{"contentType": protobufContentType}}
	res.WriteHeader(w)
}
