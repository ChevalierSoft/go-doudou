package registry

import (
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"
	"github.com/unionj-cloud/go-doudou/framework/configmgr"
	"github.com/unionj-cloud/go-doudou/framework/configmgr/mock"
	"github.com/unionj-cloud/go-doudou/framework/internal/config"
	"github.com/unionj-cloud/go-doudou/framework/memberlist"
	"github.com/unionj-cloud/go-doudou/toolkit/maputils"
	"github.com/wubin1989/nacos-sdk-go/clients/cache"
	"github.com/wubin1989/nacos-sdk-go/clients/config_client"
	"github.com/wubin1989/nacos-sdk-go/vo"
	"os"
	"reflect"
	"testing"
	"time"
)

func setup() {
	_ = config.GddMemSeed.Write("")
	_ = config.GddServiceName.Write("seed")
	_ = config.GddMemName.Write("seed")
	_ = config.GddMemPort.Write("56199")
	_ = config.GddMemWeight.Write("8")
	_ = config.GddMemDeadTimeout.Write("8s")
	_ = config.GddMemSyncInterval.Write("8s")
	_ = config.GddMemReclaimTimeout.Write("8s")
	_ = config.GddMemProbeInterval.Write("8s")
	_ = config.GddMemProbeTimeout.Write("8s")
	_ = config.GddMemSuspicionMult.Write("8")
	_ = config.GddMemGossipNodes.Write("8")
	_ = config.GddMemGossipInterval.Write("8s")
	_ = config.GddMemWeightInterval.Write("8s")
	_ = config.GddMemTCPTimeout.Write("8s")
	_ = config.GddMemHost.Write("seed.seed-svc-headless.default.svc.cluster.local")
	_ = config.GddMemIndirectChecks.Write("8")
	_ = config.GddLogLevel.Write("debug")
	_ = config.GddPort.Write("8088")
	_ = config.GddRouteRootPath.Write("/v1")
	_ = config.GddApolloAddr.Write("http://apollo-config-dev-svc:8080")
	_ = config.GddNacosServerAddr.Write("http://localhost:8848")
}

func Test_seeds(t *testing.T) {
	type args struct {
		seedstr string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				seedstr: "seed-01,seed-02,seed-03",
			},
			want: []string{"seed-01:7946", "seed-02:7946", "seed-03:7946"},
		},
		{
			name: "",
			args: args{
				seedstr: "",
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				seedstr: "seed-01:56199,seed-02,seed-03:03,seed-04:abc",
			},
			want: []string{"seed-01:56199", "seed-02:7946", "seed-03:3", "seed-04:7946"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := seeds(tt.args.seedstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("seeds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_join(t *testing.T) {
	setup()
	err := NewNode()
	if err != nil {
		panic(err)
	}
	defer Shutdown()
	_ = config.GddMemSeed.Write("not exist seed")
	_ = config.GddServiceName.Write("testsvc")
	require.Error(t, join())
}

func TestAllNodes(t *testing.T) {
	setup()
	err := NewNode()
	if err != nil {
		panic(err)
	}
	defer Shutdown()
	Convey("There should be only one node", t, func() {
		nodes, _ := AllNodes()
		So(len(nodes), ShouldEqual, 1)
	})
}

func TestInfo(t *testing.T) {
	setup()
	err := NewNode()
	if err != nil {
		panic(err)
	}
	defer Shutdown()
	Convey("Should not zero value", t, func() {
		info := Info(LocalNode())
		So(info, ShouldNotBeZeroValue)
	})
}

func TestMetaWeight(t *testing.T) {
	setup()
	err := NewNode()
	if err != nil {
		panic(err)
	}
	defer Shutdown()
	Convey("Should not zero value", t, func() {
		weight, _ := MetaWeight(LocalNode())
		So(weight, ShouldNotBeZeroValue)
	})
}

func TestSvcName(t *testing.T) {
	setup()
	err := NewNode()
	if err != nil {
		panic(err)
	}
	defer Shutdown()
	Convey("Should be equal to seed", t, func() {
		So(SvcName(LocalNode()), ShouldEqual, "seed")
	})
}

func TestRegisterServiceProvider(t *testing.T) {
	setup()
	err := NewNode()
	if err != nil {
		panic(err)
	}
	defer Shutdown()
	Convey("", t, func() {
		provider := newMockServiceProvider("TEST")
		RegisterServiceProvider(provider)
		So(len(events.ServiceProviders), ShouldEqual, 1)
		So(len(provider.servers), ShouldEqual, 1)
	})
}

func Test_memConfigListener_OnChange(t *testing.T) {
	Convey("Test OnChange callback", t, func() {
		c := &memConfigListener{
			memConf: &memberlist.Config{},
		}
		c.OnChange(&storage.ChangeEvent{
			Changes: map[string]*storage.ConfigChange{
				"gdd.mem.dead.timeout": {
					OldValue:   "60s",
					NewValue:   "30s",
					ChangeType: storage.MODIFIED,
				},
			},
		})
		Convey("Should equal to 8s", func() {
			So(os.Getenv("GDD_MEM_DEAD_TIMEOUT"), ShouldEqual, "8s")
		})

		c.OnChange(&storage.ChangeEvent{
			Changes: map[string]*storage.ConfigChange{
				"gdd.mem.dead.timeout": {
					OldValue:   "8s",
					NewValue:   "30s",
					ChangeType: storage.MODIFIED,
				},
			},
		})
		Convey("Should equal to 30s", func() {
			So(os.Getenv("GDD_MEM_DEAD_TIMEOUT"), ShouldEqual, "30s")
			So(c.memConf.GossipToTheDeadTime, ShouldEqual, 30*time.Second)
		})

		c.OnChange(&storage.ChangeEvent{
			Changes: map[string]*storage.ConfigChange{
				"gdd.mem.dead.timeout": {
					OldValue:   "30s",
					NewValue:   "",
					ChangeType: storage.DELETED,
				},
			},
		})
		Convey("Should equal to 60s", func() {
			So(os.Getenv("GDD_MEM_DEAD_TIMEOUT"), ShouldEqual, "")
			So(c.memConf.GossipToTheDeadTime, ShouldEqual, 60*time.Second)
		})

	})
}

func TestNewNode_NacosConfigType(t *testing.T) {
	Convey("Should not have error", t, func() {
		setup()
		_ = config.GddConfigRemoteType.Write("nacos")
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		dataId := ".env"
		configClient := mock.NewMockIConfigClient(ctrl)
		configClient.
			EXPECT().
			GetConfig(vo.ConfigParam{
				DataId: dataId,
				Group:  config.DefaultGddNacosConfigGroup,
			}).
			AnyTimes().
			Return("GDD_READ_TIMEOUT=60s\nGDD_WRITE_TIMEOUT=60s\nGDD_IDLE_TIMEOUT=120s", nil)

		configClient.
			EXPECT().
			ListenConfig(gomock.Any()).
			AnyTimes().
			Return(nil)

		configmgr.NewConfigClient = func(param vo.NacosClientParam) (iClient config_client.IConfigClient, err error) {
			return configClient, nil
		}

		if configmgr.NacosClient != nil {
			configmgr.NacosClient = configmgr.NewNacosConfigMgr([]string{dataId},
				config.DefaultGddNacosConfigGroup, configmgr.DotenvConfigFormat, config.DefaultGddNacosNamespaceId, configClient, cache.NewConcurrentMap())
		}

		err := configmgr.LoadFromNacos(config.GetNacosClientParam(), dataId, string(config.DefaultGddNacosConfigFormat), config.DefaultGddNacosConfigGroup)
		So(err, ShouldBeNil)

		err = NewNode()
		So(err, ShouldBeNil)
		defer Shutdown()
	})
}

func TestCallbackOnChange(t *testing.T) {
	Convey("Should equal to 30s", t, func() {
		listener := &memConfigListener{
			memConf: &memberlist.Config{},
		}
		CallbackOnChange(listener)(&configmgr.NacosChangeEvent{
			Namespace: config.DefaultGddNacosNamespaceId,
			Group:     config.DefaultGddNacosConfigGroup,
			DataId:    ".env",
			Changes: map[string]maputils.Change{
				"gdd.mem.dead.timeout": {
					OldValue:   "8s",
					NewValue:   "30s",
					ChangeType: maputils.MODIFIED,
				},
			},
		})
		CallbackOnChange(listener)(&configmgr.NacosChangeEvent{
			Namespace: config.DefaultGddNacosNamespaceId,
			Group:     config.DefaultGddNacosConfigGroup,
			DataId:    ".env",
			Changes: map[string]maputils.Change{
				"gdd.mem.dead.timeout": {
					OldValue:   "8s",
					NewValue:   "30s",
					ChangeType: maputils.MODIFIED,
				},
			},
		})
		So(os.Getenv("GDD_MEM_DEAD_TIMEOUT"), ShouldEqual, "30s")
		So(listener.memConf.GossipToTheDeadTime, ShouldEqual, 30*time.Second)
	})
}
