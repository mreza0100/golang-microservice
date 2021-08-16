package connections

import (
	pb "freeSociety/proto/generated/notification"

	"freeSociety/services/relation/configs"

	"github.com/mreza0100/golog"
	"google.golang.org/grpc"
)

func NotificationSrvConn(lgr *golog.Core) pb.NotificationServiceClient {
	conn, err := grpc.Dial(configs.Configs.Addr, getGRPCDefaultOptions(configs.Configs.ConnectionTimeout))
	if err != nil {
		lgr.Fatal("Cant connect to notification service", err)
	}

	lgr.SuccessLog("Connected to notification service")
	return pb.NewNotificationServiceClient(conn)
}
