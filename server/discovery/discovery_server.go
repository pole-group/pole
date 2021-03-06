// Copyright (c) 2020, pole-group. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"github.com/jjeffcaii/reactor-go/mono"
	polerpc "github.com/pole-group/pole-rpc"

	"github.com/pole-group/pole/common"
	"github.com/pole-group/pole/pojo"
	"github.com/pole-group/pole/server/cluster"
	"github.com/pole-group/pole/server/sys"
)

const (
	InstanceCommonPath     = sys.APIVersion + "/discovery/instance/"
	InstanceRegister       = InstanceCommonPath + "register"
	InstanceDeregister     = InstanceCommonPath + "deregister"
	InstanceMetadataUpdate = InstanceCommonPath + "metadataUpdate"
	InstanceHeartbeat      = InstanceCommonPath + "heartbeat"
	InstanceSelect         = InstanceCommonPath + "select"
	InstanceDisabled       = InstanceCommonPath + "disabled"

	ServiceCommonPath = sys.APIVersion + "/discovery/service/"
	ServiceSubscribe  = InstanceCommonPath + "subscribe"
)

type DiscoveryServer struct {
	console *DiscoveryConsole
	api     *DiscoverySdkAPI
}

func NewDiscoveryServer(ctx *common.ContextPole, mgn *cluster.ServerClusterManager, httpServer *gin.Engine) *DiscoveryServer {
	server := &DiscoveryServer{
		console: newDiscoveryConsole(mgn, httpServer),
		api:     newDiscoverySdkAPI(mgn),
	}

	server.Init(ctx)
	return server
}

func (d *DiscoveryServer) Init(ctx *common.ContextPole) {
	d.console.Init(ctx.NewSubCtx())
	d.api.Init(ctx.NewSubCtx())
	if ok, err := InitDiscoveryStorage(); !ok || err != nil {
		panic(fmt.Errorf("init discovery storage plugin failed! err : %#v", err))
	}
}

func (d *DiscoveryServer) Shutdown() {
	d.console.Shutdown()
	d.api.Shutdown()
}

// DiscoveryConsole 用于控制台
type DiscoveryConsole struct {
	ctx        *common.ContextPole
	httpServer *gin.Engine
	core       *DiscoveryCore
	mgn        *cluster.ServerClusterManager
}

func newDiscoveryConsole(mgn *cluster.ServerClusterManager, httpServer *gin.Engine) *DiscoveryConsole {
	return &DiscoveryConsole{
		httpServer: httpServer,
		mgn:        mgn,
	}
}

func (nc *DiscoveryConsole) Init(ctx *common.ContextPole) {
	nc.ctx = ctx
}

func (nc *DiscoveryConsole) Shutdown() {
	nc.ctx.Cancel()
}

type DiscoverySdkAPI struct {
	ctx    *common.ContextPole
	server polerpc.TransportServer
	mgn    *cluster.ServerClusterManager
	core   *DiscoveryCore
}

func newDiscoverySdkAPI(mgn *cluster.ServerClusterManager) *DiscoverySdkAPI {
	return &DiscoverySdkAPI{
		mgn: mgn,
	}
}

// 注册请求处理器
func (na *DiscoverySdkAPI) Init(ctx *common.ContextPole) {
	na.ctx = ctx
	na.initRpcServer(ctx)
	na.initHttpServer(ctx)
}

func (na *DiscoverySdkAPI) initRpcServer(ctx *common.ContextPole) {
	na.server = polerpc.NewRSocketServer(context.Background(), "POLE-DISCOVERY",
		na.mgn.GetSelf().GetExtensionPort(cluster.DiscoveryPort), sys.GetEnvHolder().OpenSSL)
	na.server.RegisterRequestHandler(InstanceRegister, na.instanceRegister)
	na.server.RegisterRequestHandler(InstanceDeregister, na.instanceDeregister)
	na.server.RegisterRequestHandler(InstanceMetadataUpdate, na.instanceMetadataUpdate)
	na.server.RegisterRequestHandler(InstanceDisabled, na.instanceDisabled)
	na.server.RegisterChannelRequestHandler(ServiceSubscribe, na.subscribeService)
}

func (na *DiscoverySdkAPI) initHttpServer(ctx *common.ContextPole) {
}

func (na *DiscoverySdkAPI) Shutdown() {
	na.ctx.Cancel()
}

func (na *DiscoverySdkAPI) instanceRegister(cxt context.Context, rpcCtx polerpc.RpcServerContext) {
	registerReq := &pojo.InstanceRegister{}
	err := ptypes.UnmarshalAny(rpcCtx.GetReq().Body, registerReq)
	if err != nil {
		rpcCtx.Send(&polerpc.ServerResponse{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	na.core.serviceMgn.addInstance(registerReq, rpcCtx)
}

func (na *DiscoverySdkAPI) instanceDeregister(cxt context.Context, rpcCtx polerpc.RpcServerContext) {
	registerReq := &pojo.InstanceDeregister{}
	err := ptypes.UnmarshalAny(rpcCtx.GetReq().Body, registerReq)
	if err != nil {
		rpcCtx.Send(&polerpc.ServerResponse{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	na.core.serviceMgn.removeInstance(registerReq, rpcCtx)
}

// just for Http API
func (na *DiscoverySdkAPI) instanceHeartbeat(cxt context.Context, sink mono.Sink) {
}

func (na *DiscoverySdkAPI) instanceDisabled(cxt context.Context, rpcCtx polerpc.RpcServerContext) {
	registerReq := &pojo.InstanceDisabled{}
	err := ptypes.UnmarshalAny(rpcCtx.GetReq().Body, registerReq)
	if err != nil {
		rpcCtx.Send(&polerpc.ServerResponse{
			Code: -1,
			Msg:  err.Error(),
		})
	}
}

func (na *DiscoverySdkAPI) instanceMetadataUpdate(cxt context.Context, rpcCtx polerpc.RpcServerContext) {
}

// subscribeService 服务订阅
func (na *DiscoverySdkAPI) subscribeService(ctx context.Context, rpcCtx polerpc.RpcServerContext) {

}
