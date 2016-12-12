package system

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

/*
Subsystem is a meta info struct for subsystems
that are loaded into the core daemon
*/
type Subsystem struct {
	Role   SubsystemType
	Loaded bool
	Active bool
}

// NewSubsystem returns default initialized Subsystem
func NewSubsystem(role SubsystemType) Subsystem {
	return Subsystem{Role: role}
}

// Subsystems is a list of Subsystem
type Subsystems []*Subsystem

/*
Add subsystem to a list
*/
func (systems *Subsystems) Add(system *Subsystem) {
	*systems = append(*systems, system)
}

/*
SubsystemType describes the type of the subsystem.
*/
type SubsystemType string

const (
	Gateway = "gateway"

	// Crawler handles crawling the network and queing data for ingestion
	Crawler = "crawler"

	// ForwardIndex handles building the forward index
	ForwardIndex = "findex"

	// InvertedIndex handles building the inverted index
	InvertedIndex = "iindex"

	// HTMLParser handles parsing HTML files
	HTMLParser = "htmlparse"
)

func StartSubsystemServer(port int) {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	//	trinity.RegisterRouteGuideServer(grpcServer, &routeGuideServer{})
	grpcServer.Serve(lis)
}
