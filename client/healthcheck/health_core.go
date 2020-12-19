// Copyright (c) 2020, Conf-Group. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package healthcheck

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Conf-Group/pole/common"
	"github.com/Conf-Group/pole/pojo"
	"github.com/Conf-Group/pole/transport"
	"github.com/Conf-Group/pole/utils"
)

type CustomerHealthChecker func() bool

type customerTimeF struct {
	owner  *CustomHealthCheckPlugin
	task   *CustomerHealthCheckTask
	cancel int32
}

func (ctf *customerTimeF) Run() {
	defer func() {
		if atomic.LoadInt32(&ctf.cancel) == 1 {
			return
		}
		ctf.owner.htw.AddTask(ctf, time.Duration(5)*time.Second)
	}()
	if !ctf.task.F() {
		// TODO 当前实例的健康状态为非健康
	}
}

// 用户自定义健康检查方式
type CustomHealthCheckPlugin struct {
	owner          *HealthCheckManager
	rwLock         sync.RWMutex
	htw            *utils.HashTimeWheel
	taskRepository map[string]*customerTimeF
}

func (h *CustomHealthCheckPlugin) Name() string {
	return HealthCheck_Customer
}

func (h *CustomHealthCheckPlugin) Init(ctx *common.ContextPole) {
	h.taskRepository = make(map[string]*customerTimeF)
	h.htw = utils.NewTimeWheel(time.Duration(5)*time.Second, 16)
}

func (h *CustomHealthCheckPlugin) Run() {
	h.htw.Start()
}

func (h *CustomHealthCheckPlugin) AddTask(task HealthCheckTask) (bool, error) {
	cTask := task.(*CustomerHealthCheckTask)
	defer h.rwLock.Unlock()
	h.rwLock.Lock()
	ctf := &customerTimeF{owner: h, task: cTask}
	h.taskRepository[fmt.Sprintf("%s:%d", cTask.Instance.Ip, cTask.Instance.Port)] = ctf
	h.htw.AddTask(ctf, time.Duration(5)*time.Second)
	return true, nil
}

func (h *CustomHealthCheckPlugin) RemoveTask(task HealthCheckTask) (bool, error) {
	cTask := task.(*CustomerHealthCheckTask)
	key := fmt.Sprintf("%s:%d", cTask.Instance.Ip, cTask.Instance.Port)
	defer h.rwLock.Unlock()
	h.rwLock.Lock()

	if t, exist := h.taskRepository[key]; exist {
		atomic.StoreInt32(&t.cancel, 1)
	}

	delete(h.taskRepository, key)
	return true, nil
}

func (h *CustomHealthCheckPlugin) Destroy() {
	h.htw.Stop()
}

// 基于链接心跳包的检查
type heartbeatTask struct {
	task   *HeartbeatCheckTask
	cancel int32
}

func (hbt *heartbeatTask) Run() {

}

type ConnectionBeatHealthCheckPlugin struct {
	owner          *HealthCheckManager
	client         *transport.HttpClient
	rwLock         sync.RWMutex
	htw            *utils.HashTimeWheel
	taskRepository map[string]*heartbeatTask
}

func (h *ConnectionBeatHealthCheckPlugin) Name() string {
	return HealthCheck_Heartbeat
}

func (h *ConnectionBeatHealthCheckPlugin) Init(ctx *common.ContextPole) {
	h.client = transport.NewHttpClient(false)
	h.taskRepository = make(map[string]*heartbeatTask)
	h.htw = utils.NewTimeWheel(time.Duration(5)*time.Second, 16)
}

func (h *ConnectionBeatHealthCheckPlugin) Run() {
	h.htw.Start()
}

func (h *ConnectionBeatHealthCheckPlugin) AddTask(task HealthCheckTask) (bool, error) {
	httpTask := task.(*HeartbeatCheckTask)
	defer h.rwLock.Unlock()
	h.rwLock.Lock()

	ht := &heartbeatTask{
		task:   httpTask,
		cancel: 0,
	}

	h.taskRepository[fmt.Sprintf("%s:%d", httpTask.Instance.Ip, httpTask.Instance.Port)] = ht
	h.htw.AddTask(ht, time.Duration(5)*time.Second)
	return true, nil
}

func (h *ConnectionBeatHealthCheckPlugin) RemoveTask(task HealthCheckTask) (bool, error) {
	httpTask := task.(*HeartbeatCheckTask)
	defer h.rwLock.Unlock()
	h.rwLock.Lock()

	serverAddr := fmt.Sprintf("%s:%d", httpTask.Instance.Ip, httpTask.Instance.Port)

	if t, exist := h.taskRepository[serverAddr]; exist {
		atomic.StoreInt32(&t.cancel, 1)
	}

	delete(h.taskRepository, serverAddr)
	return true, nil
}

func (h *ConnectionBeatHealthCheckPlugin) Destroy() {
	h.htw.Stop()
}

// HTTP 返回码检测
type httpCodeTask struct {
	task   *HttpCodeCheckTask
	cancel int32
}

func (hct *httpCodeTask) Run() {

}

type HttpCodeHealthCheckPlugin struct {
	owner          *HealthCheckManager
	client         *http.Client
	rwLock         sync.RWMutex
	htw            *utils.HashTimeWheel
	taskRepository map[string]*httpCodeTask
}

func (hch *HttpCodeHealthCheckPlugin) Init(cxt *common.ContextPole) {
	hch.client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:          1200,            // 连接池中最大连接数
			MaxIdleConnsPerHost:   300,             // 连接池中每个ip的最大连接数
			TLSHandshakeTimeout:   5 * time.Second, // 限制TLS握手的时间
			ResponseHeaderTimeout: 5 * time.Second, // 限制读取response header的超时时间
			IdleConnTimeout:       90 * time.Second,
		},
		Timeout: time.Duration(30) * time.Second,
	}
	hch.taskRepository = make(map[string]*httpCodeTask)
	hch.htw = utils.NewTimeWheel(time.Duration(5)*time.Second, 16)
}

func (hch *HttpCodeHealthCheckPlugin) Run() {
	hch.htw.Start()
}

func (hch *HttpCodeHealthCheckPlugin) Name() string {
	return HealthCheck_Http
}

func (hch *HttpCodeHealthCheckPlugin) AddTask(task HealthCheckTask) (bool, error) {
	htTask := task.(*HttpCodeCheckTask)
	servAddr := fmt.Sprintf("%s:%d", htTask.Instance.Ip, htTask.Instance.Port)
	defer hch.rwLock.Unlock()
	hch.rwLock.Lock()

	ht := &httpCodeTask{
		task:   htTask,
		cancel: 0,
	}

	hch.taskRepository[servAddr] = ht
	hch.htw.AddTask(ht, time.Duration(5)*time.Second)
	return true, nil
}

func (hch *HttpCodeHealthCheckPlugin) RemoveTask(task HealthCheckTask) (bool, error) {
	htTask := task.(*HttpCodeCheckTask)
	servAddr := fmt.Sprintf("%s:%d", htTask.Instance.Ip, htTask.Instance.Port)
	defer hch.rwLock.Unlock()
	hch.rwLock.Lock()

	if t, exist := hch.taskRepository[servAddr]; exist {
		atomic.StoreInt32(&t.cancel, 1)
	}

	delete(hch.taskRepository, servAddr)
	return true, nil
}

func (hch *HttpCodeHealthCheckPlugin) Destroy() {

}

// TCP 健康检查
type TcpHealthCheckPlugin struct {
	owner          *HealthCheckManager
	rwLock         sync.RWMutex
	htw            *utils.HashTimeWheel
	taskRepository map[string]*tcpClient
}

type tcpClient struct {
	owner      *TcpHealthCheckPlugin
	serverAddr string
	conn       *net.TCPConn
	instance   *pojo.Instance
	cancel     int32
}

func (tc *tcpClient) Run() {
	defer func() {
		if atomic.LoadInt32(&tc.cancel) == 1 {
			if tc.conn != nil {
				_ = tc.conn.Close()
			}
			return
		}
		tc.owner.htw.AddTask(tc, time.Duration(5)*time.Second)
	}()
	if tc.conn == nil {
		tcpAddr, err := net.ResolveTCPAddr("tcp", tc.serverAddr)
		if err != nil {
			// TODO 不健康
			return
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			// TODO 不健康
			return
		}
		tc.conn = conn
	} else {
		r := make([]byte, 1, 1)
		if _, err := tc.conn.Read(r); err != nil {
			// TODO 不健康
			return
		}
	}
}

func (h *TcpHealthCheckPlugin) Name() string {
	return HealthCheck_Tcp
}

func (h *TcpHealthCheckPlugin) Init(ctx *common.ContextPole) {
	h.taskRepository = make(map[string]*tcpClient)
	h.htw = utils.NewTimeWheel(time.Duration(5)*time.Second, 16)
}

func (h *TcpHealthCheckPlugin) Run() {
	h.htw.Start()
}

func (h *TcpHealthCheckPlugin) AddTask(task HealthCheckTask) (bool, error) {
	tcpTask := task.(*TcpHealthCheckTask)
	servAddr := fmt.Sprintf("%s:%d", tcpTask.Instance.Ip, tcpTask.Instance.Port)
	defer h.rwLock.Unlock()
	h.rwLock.Lock()

	tc := &tcpClient{
		owner:      h,
		serverAddr: servAddr,
		instance:   &tcpTask.Instance,
	}

	h.taskRepository[servAddr] = tc
	h.htw.AddTask(tc, time.Duration(5)*time.Second)
	return true, nil
}

func (h *TcpHealthCheckPlugin) RemoveTask(task HealthCheckTask) (bool, error) {
	tcpTask := task.(*TcpHealthCheckTask)
	servAddr := fmt.Sprintf("%s:%d", tcpTask.Instance.Ip, tcpTask.Instance.Port)
	defer h.rwLock.Unlock()
	h.rwLock.Lock()

	if t, exist := h.taskRepository[servAddr]; exist {
		atomic.StoreInt32(&t.cancel, 1)
	}

	delete(h.taskRepository, servAddr)
	return true, nil
}

func (h *TcpHealthCheckPlugin) Destroy() {
	h.htw.Stop()
}
