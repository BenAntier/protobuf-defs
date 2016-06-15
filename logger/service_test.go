package logger

import (
	"fmt"
	"net"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type loggerService struct {
}

func (ls *loggerService) LogEvent(ctx context.Context, er *Event) (*Empty, error) {
	return &Empty{}, nil
}

func (ls *loggerService) LogError(ctx context.Context, er *Error) (*Empty, error) {
	return &Empty{}, nil
}

const serverAddr = "127.0.0.1"

var serverPort = 2730

func newServer(t *testing.T) *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverAddr, serverPort))
	if err != nil {
		serverPort++
		return newServer(t)
	}

	s := grpc.NewServer()
	RegisterLoggerServer(s, new(loggerService))
	go s.Serve(lis)
	return s
}

func TestEvent(t *testing.T) {
	server := newServer(t)
	defer server.Stop()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", serverAddr, serverPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to open a connection to the RPC server: %s", err)
	}
	defer conn.Close()
	client := NewLoggerClient(conn)

	er := &Event{
		Type: Event_PAGE_VIEW,
		Client: &Client{
			PID: "abc",
			SID: "def",
		},
		User: &User{
			SessID:    "session 123",
			Path:      "/hello",
			Adblocker: true,
			UID:       "uid 123",
			PuID:      "PUID 123",
			UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.75 Safari/537.36",
			Location: &User_GeoIP{
				IP:      "8.8.8.8",
				Country: "USA",
				City:    "Palo Alto",
			},
		},
		Custom: []*Pair{
			{
				Key:   "c1",
				Value: "abcd",
			},
		},
	}

	if _, err := client.LogEvent(context.Background(), er); err != nil {
		t.Fatalf("Event returned an unexpected error: %s", err)
	}
}

func TestError(t *testing.T) {
	server := newServer(t)
	defer server.Stop()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", serverAddr, serverPort), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to open a connection to the RPC server: %s", err)
	}
	defer conn.Close()
	client := NewLoggerClient(conn)

	er := &Error{
		Client: &Client{
			PID: "abc",
			SID: "def",
		},
		User: &User{
			SessID:    "session 123",
			Path:      "/hello",
			Adblocker: true,
			UID:       "uid 123",
			PuID:      "PUID 123",
			UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.75 Safari/537.36",
			Location: &User_GeoIP{
				IP:      "8.8.8.8",
				Country: "USA",
				City:    "Palo Alto",
			},
		},
		Custom: []*Pair{
			{
				Key:   "c1",
				Value: "abcd",
			},
		},
		ErrorCode: 1000,
	}

	if _, err := client.LogError(context.Background(), er); err != nil {
		t.Fatalf("Error returned an unexpected error: %s", err)
	}
}
