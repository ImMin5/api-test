package dist

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strings"

	inventorygw "github.com/ImMin5/api-test/gateway/spaceone/api/inventory/v1" // Update
)

func CustomMatcher(key string) (string, bool) {

	switch key {
	case "Authorization":
		fmt.Print(key)
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func myFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	//t, ok := resp.(*externalpb.Tokenizer)
	//if ok {
	//	w.Header().Set("X-My-Tracking-Token", t.Token)
	//	t.Token = ""
	//}
	auth := ""

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if uID, ok := md["x-user-id"]; ok {
			auth = strings.Join(uID, ",")
		}
	}

	fmt.Println("auth: ", auth)

	return nil
}

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

func makeMetadata(ctx context.Context, req *http.Request) metadata.MD {
	print(req.Header.Get("Authorization"))
	token := req.Header.Get("Authorization")
	if token != "" {
		token = strings.Split(token, " ")[1]
		print(token)
		req.Header.Set("token", token)
	}

	return metadata.Pairs("token", token)
}

func Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // end of main run function cancel() will be executed
	print("run")
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible

	mux := runtime.NewServeMux(
		runtime.WithMetadata(makeMetadata),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		inventorygw.RegisterCloudServiceHandlerFromEndpoint,
		inventorygw.RegisterChangeHistoryHandlerFromEndpoint,
		inventorygw.RegisterRegionHandlerFromEndpoint,
		inventorygw.RegisterCloudServiceHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, *grpcServerEndpoint, opts); err != nil {
			return err
		}
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe("localhost:8084", mux)
}

//func main() {
//	flag.Parse()
//	defer glog.Flush()
//
//	if err := run(); err != nil {
//		glog.Fatal(err)
//	}
//}
