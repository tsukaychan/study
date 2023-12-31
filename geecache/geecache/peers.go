package geecache

import (
	pb "geecache/geecache/geecachepb"
)

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(req *pb.Request, resp *pb.Response) error
}
