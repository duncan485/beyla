package transform

import (
	"bytes"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/grafana/http-autoinstrument/pkg/ebpf/nethttp"

	"github.com/gavv/monotime"
)

// HTTPRequestSpan contains the information being submitted as
type HTTPRequestSpan struct {
	Method   string
	Path     string
	Route    string
	Peer     string
	Host     string
	HostPort int
	LocalIP  string
	Status   int
	Start    time.Time
	End      time.Time
}

var localIP = getLocalIPv4()

func getLocalIPv4() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return ip.IP.String()
			}
		}
	}
	return ""
}

func ConvertToSpan(in <-chan nethttp.HTTPRequestTrace, out chan<- HTTPRequestSpan) {
	cnv := newConverter()
	for trace := range in {
		out <- cnv.convert(&trace)
	}
}

func newConverter() converter {
	return converter{
		monoClock: monotime.Now,
		clock:     time.Now,
	}
}

type converter struct {
	clock     func() time.Time
	monoClock func() time.Duration
}

func extractHostPort(b []uint8) (string, int) {
	addrLen := bytes.IndexByte(b, 0)
	if addrLen < 0 {
		addrLen = len(b)
	}

	peer := ""
	peerPort := 0

	if addrLen > 0 {
		addr := string(b[:addrLen])
		ip, port, err := net.SplitHostPort(addr)
		if err != nil {
			peer = addr
		} else {
			peer = ip
			peerPort, _ = strconv.Atoi(port)
		}
	}

	return peer, peerPort
}

func (c *converter) convert(trace *nethttp.HTTPRequestTrace) HTTPRequestSpan {
	now := time.Now()
	monoNow := c.monoClock()
	startDelta := monoNow - time.Duration(trace.StartMonotimeNs)
	endDelta := monoNow - time.Duration(trace.EndMonotimeNs)

	// From C, assuming 0-ended strings
	methodLen := bytes.IndexByte(trace.Method[:], 0)
	if methodLen < 0 {
		methodLen = len(trace.Method)
	}
	pathLen := bytes.IndexByte(trace.Path[:], 0)
	if pathLen < 0 {
		pathLen = len(trace.Path)
	}

	peer, _ := extractHostPort(trace.RemoteAddr[:])
	_, hostPort := extractHostPort(trace.Host[:])

	// We ignore the hostname from the request, and use the hostname of the machine
	hostname, _ := os.Hostname()

	return HTTPRequestSpan{
		Method:   string(trace.Method[:methodLen]),
		Path:     string(trace.Path[:pathLen]),
		Peer:     peer,
		Host:     hostname,
		HostPort: hostPort,
		LocalIP:  localIP,
		Start:    now.Add(-startDelta),
		End:      now.Add(-endDelta),
		Status:   int(trace.Status),
	}
}
