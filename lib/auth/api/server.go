/*
Copyright 2016 SPIFFE Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"bytes"
	"io"

	"github.com/gravitational/teleport/lib/auth/api/protogen"
	"github.com/gravitational/teleport/lib/session"

	"github.com/gravitational/trace/trail"
)

type ChunkPoster interface {
	PostSessionChunk(namespace string, sid session.ID, reader io.Reader) error
}

type NopPoster struct{}

func (*NopPoster) PostSessionChunk(namespace string, sid session.ID, reader io.Reader) error {
	return nil
}

func NewServer(poster ChunkPoster) (*Server, error) {
	if poster == nil {
		poster = &NopPoster{}
	}
	return &Server{poster: poster}, nil
}

type Server struct {
	poster ChunkPoster
}

// SessionChunks
func (s *Server) SessionChunks(srv protogen.Audit_SessionChunksServer) error {
	ctx := srv.Context()
	for {
		chunk, err := srv.Recv()
		if err != nil {
			return trail.Send(ctx, err)
		}
		err = s.poster.PostSessionChunk(chunk.Namespace, session.ID(chunk.SessionID), bytes.NewReader(chunk.Chunk))
		if err != nil {
			return trail.Send(ctx, err)
		}
	}
}
