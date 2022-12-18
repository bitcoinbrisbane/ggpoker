package p2p

import (
	"context"
	"net"
	"sync"

	"github.com/anthdm/ggpoker/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Node struct {
	ServerConfig

	peerLock sync.RWMutex
	peers    map[proto.GossipClient]*proto.Version

	proto.UnimplementedGossipServer
}

func NewNode(cfg ServerConfig) *Node {
	return &Node{
		ServerConfig: cfg,
		peers:        make(map[proto.GossipClient]*proto.Version),
	}
}

func (n *Node) Handshake(ctx context.Context, version *proto.Version) (*proto.Version, error) {
	client, err := makeGrpcClientConn(version.ListenAddr)
	if err != nil {
		return nil, err
	}

	n.addPeer(client, version)

	return n.getVersion(), nil
}

func (n *Node) addPeer(c proto.GossipClient, v *proto.Version) {
	n.peerLock.Lock()
	defer n.peerLock.Unlock()
	n.peers[c] = v

	logrus.WithFields(logrus.Fields{
		"we":     n.ListenAddr,
		"remote": v.ListenAddr,
	}).Info("new player connected")
}

func (n *Node) getVersion() *proto.Version {
	return &proto.Version{
		Version:    "GGPOKER-0.0.1",
		ListenAddr: n.ListenAddr,
	}
}

func (n *Node) Connect(addr string) error {
	client, err := makeGrpcClientConn(addr)
	if err != nil {
		return err
	}
	hs, err := client.Handshake(context.TODO(), n.getVersion())
	if err != nil {
		return err
	}
	n.addPeer(client, hs)
	return nil
}

func (n *Node) Start() error {
	grpcServer := grpc.NewServer()
	proto.RegisterGossipServer(grpcServer, n)

	ln, err := net.Listen("tcp", n.ListenAddr)
	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"port":       n.ListenAddr,
		"variant":    n.GameVariant,
		"maxPlayers": n.MaxPlayers,
	}).Info("started new poker game server")

	return grpcServer.Serve(ln)
}

func makeGrpcClientConn(addr string) (proto.GossipClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := proto.NewGossipClient(conn)
	return client, nil
}