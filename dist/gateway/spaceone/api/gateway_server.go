package api

import (
	"context"
	v1Board "github.com/cloudforet-io/api/dist/gateway/spaceone/api/board/v1"
	v1Config "github.com/cloudforet-io/api/dist/gateway/spaceone/api/config/v1"
	v1Core "github.com/cloudforet-io/api/dist/gateway/spaceone/api/core/v1"
	v1CostAnalysis "github.com/cloudforet-io/api/dist/gateway/spaceone/api/cost_analysis/v1"
	v1Dashboard "github.com/cloudforet-io/api/dist/gateway/spaceone/api/dashboard/v1"
	v1FileManager "github.com/cloudforet-io/api/dist/gateway/spaceone/api/file_manager/v1"
	v1Identity "github.com/cloudforet-io/api/dist/gateway/spaceone/api/identity/v1"
	v1Inventory "github.com/cloudforet-io/api/dist/gateway/spaceone/api/inventory/v1"
	v1Monitoring "github.com/cloudforet-io/api/dist/gateway/spaceone/api/monitoring/v1"
	v1Notification "github.com/cloudforet-io/api/dist/gateway/spaceone/api/notification/v1"
	v1Plugin "github.com/cloudforet-io/api/dist/gateway/spaceone/api/plugin/v1"
	v1Repository "github.com/cloudforet-io/api/dist/gateway/spaceone/api/repository/v1"
	v1Sample "github.com/cloudforet-io/api/dist/gateway/spaceone/api/sample/v1"
	v1Secret "github.com/cloudforet-io/api/dist/gateway/spaceone/api/secret/v1"
	v1Statistics "github.com/cloudforet-io/api/dist/gateway/spaceone/api/statistics/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"reflect"
)

func RegisterServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption, service string) {
	serviceFuncMap := map[string]interface{}{
		"board":         registerBoard,
		"config":        registerConfig,
		"core":          registerCore,
		"cost_analysis": registerCostAnalysis,
		"dashboard":     registerDashboard,
		"file_manager":  registerFileManager,
		"identity":      registerIdentity,
		"inventory":     registerInventory,
		"monitoring":    registerMonitoring,
		"notification":  registerNotification,
		"plugin":        registerPlugin,
		"repository":    registerRepository,
		"sample":        registerSample,
		"secret":        registerSecret,
		"statistics":    registerStatistics,
	}

	// Init log format
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if serviceFuncMap[service] == nil {
		log.Panic("[ERROR] ", service, " service is not exist")
	}

	serviceFunc := serviceFuncMap[service]
	funcValue := reflect.ValueOf(serviceFunc)
	args := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(mux), reflect.ValueOf(grpcServerEndpoint), reflect.ValueOf(opts)}
	funcValue.Call(args)
}

func registerBoard(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Board.RegisterPostHandlerFromEndpoint,
		v1Board.RegisterBoardHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[boardRegister] Register board service failed because of ", err)
		}
	}

	log.Println("[INFO] [boardRegister] Register board service success")
}

func registerConfig(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Config.RegisterUserConfigHandlerFromEndpoint,
		v1Config.RegisterDomainConfigHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[configRegister] Register config service failed because of ", err)
		}
	}

	log.Println("[INFO] [configRegister] Register config service success")
}

func registerCore(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Core.RegisterServerInfoHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[coreRegister] Register core service failed because of ", err)
		}
	}

	log.Println("[INFO] [coreRegister] Register core service success")
}

func registerCostAnalysis(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1CostAnalysis.RegisterCostHandlerFromEndpoint,
		v1CostAnalysis.RegisterCostQuerySetHandlerFromEndpoint,
		v1CostAnalysis.RegisterExchangeRateHandlerFromEndpoint,
		v1CostAnalysis.RegisterDataSourceHandlerFromEndpoint,
		v1CostAnalysis.RegisterJobTaskHandlerFromEndpoint,
		v1CostAnalysis.RegisterJobHandlerFromEndpoint,
		v1CostAnalysis.RegisterBudgetUsageHandlerFromEndpoint,
		v1CostAnalysis.RegisterCustomWidgetHandlerFromEndpoint,
		v1CostAnalysis.RegisterScheduleHandlerFromEndpoint,
		v1CostAnalysis.RegisterPublicDashboardHandlerFromEndpoint,
		v1CostAnalysis.RegisterUserDashboardHandlerFromEndpoint,
		v1CostAnalysis.RegisterDataSourceRuleHandlerFromEndpoint,
		v1CostAnalysis.RegisterBudgetHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[cost_analysisRegister] Register cost_analysis service failed because of ", err)
		}
	}

	log.Println("[INFO] [cost_analysisRegister] Register cost_analysis service success")
}

func registerDashboard(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Dashboard.RegisterDomainDashboardHandlerFromEndpoint,
		v1Dashboard.RegisterCustomWidgetHandlerFromEndpoint,
		v1Dashboard.RegisterProjectDashboardHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[dashboardRegister] Register dashboard service failed because of ", err)
		}
	}

	log.Println("[INFO] [dashboardRegister] Register dashboard service success")
}

func registerFileManager(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1FileManager.RegisterFileHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[file_managerRegister] Register file_manager service failed because of ", err)
		}
	}

	log.Println("[INFO] [file_managerRegister] Register file_manager service success")
}

func registerIdentity(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Identity.RegisterTokenHandlerFromEndpoint,
		v1Identity.RegisterPolicyHandlerFromEndpoint,
		v1Identity.RegisterEndpointHandlerFromEndpoint,
		v1Identity.RegisterAuthorizationHandlerFromEndpoint,
		v1Identity.RegisterRoleHandlerFromEndpoint,
		v1Identity.RegisterProviderHandlerFromEndpoint,
		v1Identity.RegisterServiceAccountHandlerFromEndpoint,
		v1Identity.RegisterAPIKeyHandlerFromEndpoint,
		v1Identity.RegisterProjectHandlerFromEndpoint,
		v1Identity.RegisterDomainHandlerFromEndpoint,
		v1Identity.RegisterRoleBindingHandlerFromEndpoint,
		v1Identity.RegisterProjectGroupHandlerFromEndpoint,
		v1Identity.RegisterUserHandlerFromEndpoint,
		v1Identity.RegisterDomainOwnerHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[identityRegister] Register identity service failed because of ", err)
		}
	}

	log.Println("[INFO] [identityRegister] Register identity service success")
}

func registerInventory(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Inventory.RegisterNoteHandlerFromEndpoint,
		v1Inventory.RegisterCloudServiceTypeHandlerFromEndpoint,
		v1Inventory.RegisterResourceGroupHandlerFromEndpoint,
		v1Inventory.RegisterChangeHistoryHandlerFromEndpoint,
		v1Inventory.RegisterRegionHandlerFromEndpoint,
		v1Inventory.RegisterJobTaskHandlerFromEndpoint,
		v1Inventory.RegisterJobHandlerFromEndpoint,
		v1Inventory.RegisterCollectorHandlerFromEndpoint,
		v1Inventory.RegisterCloudServiceStatsHandlerFromEndpoint,
		v1Inventory.RegisterCloudServiceQuerySetHandlerFromEndpoint,
		v1Inventory.RegisterCollectorRuleHandlerFromEndpoint,
		v1Inventory.RegisterCloudServiceHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[inventoryRegister] Register inventory service failed because of ", err)
		}
	}

	log.Println("[INFO] [inventoryRegister] Register inventory service success")
}

func registerMonitoring(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Monitoring.RegisterNoteHandlerFromEndpoint,
		v1Monitoring.RegisterEscalationPolicyHandlerFromEndpoint,
		v1Monitoring.RegisterDataSourceHandlerFromEndpoint,
		v1Monitoring.RegisterAlertHandlerFromEndpoint,
		v1Monitoring.RegisterLogHandlerFromEndpoint,
		v1Monitoring.RegisterEventRuleHandlerFromEndpoint,
		v1Monitoring.RegisterMaintenanceWindowHandlerFromEndpoint,
		v1Monitoring.RegisterMetricHandlerFromEndpoint,
		v1Monitoring.RegisterEventHandlerFromEndpoint,
		v1Monitoring.RegisterProjectAlertConfigHandlerFromEndpoint,
		v1Monitoring.RegisterWebhookHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[monitoringRegister] Register monitoring service failed because of ", err)
		}
	}

	log.Println("[INFO] [monitoringRegister] Register monitoring service success")
}

func registerNotification(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Notification.RegisterProjectChannelHandlerFromEndpoint,
		v1Notification.RegisterNotificationUsageHandlerFromEndpoint,
		v1Notification.RegisterQuotaHandlerFromEndpoint,
		v1Notification.RegisterUserChannelHandlerFromEndpoint,
		v1Notification.RegisterNotificationHandlerFromEndpoint,
		v1Notification.RegisterProtocolHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[notificationRegister] Register notification service failed because of ", err)
		}
	}

	log.Println("[INFO] [notificationRegister] Register notification service success")
}

func registerPlugin(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Plugin.RegisterSupervisorHandlerFromEndpoint,
		v1Plugin.RegisterPluginHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[pluginRegister] Register plugin service failed because of ", err)
		}
	}

	log.Println("[INFO] [pluginRegister] Register plugin service success")
}

func registerRepository(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Repository.RegisterPolicyHandlerFromEndpoint,
		v1Repository.RegisterPluginHandlerFromEndpoint,
		v1Repository.RegisterRepositoryHandlerFromEndpoint,
		v1Repository.RegisterSchemaHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[repositoryRegister] Register repository service failed because of ", err)
		}
	}

	log.Println("[INFO] [repositoryRegister] Register repository service success")
}

func registerSample(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Sample.RegisterHelloWorldHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[sampleRegister] Register sample service failed because of ", err)
		}
	}

	log.Println("[INFO] [sampleRegister] Register sample service success")
}

func registerSecret(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Secret.RegisterTrustedSecretHandlerFromEndpoint,
		v1Secret.RegisterSecretGroupHandlerFromEndpoint,
		v1Secret.RegisterSecretHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[secretRegister] Register secret service failed because of ", err)
		}
	}

	log.Println("[INFO] [secretRegister] Register secret service success")
}

func registerStatistics(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint string, opts []grpc.DialOption) {
	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		v1Statistics.RegisterStorageHandlerFromEndpoint,
		v1Statistics.RegisterResourceHandlerFromEndpoint,
		v1Statistics.RegisterHistoryHandlerFromEndpoint,
		v1Statistics.RegisterScheduleHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcServerEndpoint, opts); err != nil {
			log.Print("[statisticsRegister] Register statistics service failed because of ", err)
		}
	}

	log.Println("[INFO] [statisticsRegister] Register statistics service success")
}
