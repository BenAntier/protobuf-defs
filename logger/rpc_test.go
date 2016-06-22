package logger

import (
	"fmt"
	"net"
	"testing"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type loggerService struct {
}

func (ls *loggerService) LogEvent(ctx context.Context, er *Event) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (ls *loggerService) LogError(ctx context.Context, er *Error) (*empty.Empty, error) {
	return &empty.Empty{}, nil
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
		Publisher: &Publisher{
			Uuid:     "abc",
			ScriptId: "def",
		},
		User: &User{
			Session: &Session{
				Id:        "abc",
				UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.75 Safari/537.36",
				Adblocker: true,
				IsBot:     false,
			},
			Page: &Page{
				Url:    "https://www.google.com/?q=hello",
				Scheme: Page_HTTPS,
				Domain: "www.google.com",
				Path:   "/",
				Params: map[string]string{
					"q": "hello",
				},
				Referrer: "https://www.facebook.com",
				LoadTime: &duration.Duration{
					Seconds: 10,
					Nanos:   99,
				},
			},
			Location: &GeoIp{
				Ip:         &GeoIp_V4{V4: "8.8.8.8"},
				Country:    "US",
				City:       "Mountain View",
				PostalCode: "94035",
				MetroCode:  807,
				Isp: &GeoIp_Isp{
					AutonomousSystemNumber:       123,
					AutonomousSystemOrganization: "Comcast",
					Name:         "Comcast ISP",
					Organization: "Comcast Organization",
				},
			},
		},
		Custom: map[string]string{
			"c1": "abcd",
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
		Code: Error_RPC_CALL,
		Publisher: &Publisher{
			Uuid:     "abc",
			ScriptId: "def",
		},
		User: &User{
			Session: &Session{
				Id:        "abc",
				UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.75 Safari/537.36",
				Adblocker: true,
				IsBot:     false,
			},
			Page: &Page{
				Url:    "https://www.google.com/?q=hello",
				Scheme: Page_HTTPS,
				Domain: "www.google.com",
				Path:   "/",
				Params: map[string]string{
					"q": "hello",
				},
				Referrer: "https://www.facebook.com",
				LoadTime: &duration.Duration{
					Seconds: 10,
					Nanos:   99,
				},
			},
			Location: &GeoIp{
				Ip:         &GeoIp_V4{V4: "8.8.8.8"},
				Country:    "US",
				City:       "Mountain View",
				PostalCode: "94035",
				MetroCode:  807,
				Isp: &GeoIp_Isp{
					AutonomousSystemNumber:       123,
					AutonomousSystemOrganization: "Comcast",
					Name:         "Comcast ISP",
					Organization: "Comcast Organization",
				},
			},
		},
		Custom: map[string]string{
			"c1": "abcd",
		},
	}

	if _, err := client.LogError(context.Background(), er); err != nil {
		t.Fatalf("Error returned an unexpected error: %s", err)
	}
}
