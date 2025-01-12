// Copyright 2023 The Outline Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httpconnect

import (
	"errors"
	"github.com/Jigsaw-Code/outline-sdk/transport"
	"io"
)

// PipeConn is a [transport.StreamConn] that overrides [Read], [Write] (and corresponding [Close]) functions with the given [reader] and [writer]
type PipeConn struct {
	reader io.ReadCloser
	writer io.WriteCloser
	transport.StreamConn
}

func (p *PipeConn) Read(b []byte) (n int, err error) {
	return p.reader.Read(b)
}

func (p *PipeConn) Write(b []byte) (n int, err error) {
	return p.writer.Write(b)
}

func (p *PipeConn) CloseRead() error {
	return errors.Join(p.reader.Close(), p.StreamConn.CloseRead())
}

func (p *PipeConn) CloseWrite() error {
	return errors.Join(p.writer.Close(), p.StreamConn.CloseWrite())
}

func (p *PipeConn) Close() error {
	return errors.Join(p.reader.Close(), p.writer.Close(), p.StreamConn.Close())
}
