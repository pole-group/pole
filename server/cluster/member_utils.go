// Copyright (c) 2020, pole-group. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cluster

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/pole-group/pole/server/sys"
	"github.com/pole-group/pole/utils"
)

func OnSuccess(m *Member) {
	m.accessFailCnt = 0
	m.Status = Health
}

func OnFail(m *Member, err error) {
	m.Status = Impeach
	if strings.ContainsAny(err.Error(), "Connection refused") {
		m.Status = Down
	} else {
		m.accessFailCnt++
		if m.accessFailCnt >= sys.GetEnvHolder().ClusterCfg.LookupCfg.MaxProbeFailCnt {
			m.Status = Down
		}
	}
}

func MultiParse(arr ...string) []*Member {
	mList := make([]*Member, len(arr))
	for _, s := range arr {
		mList = append(mList, SingParse(s))
	}
	return mList
}

func SingParse(s string) *Member {
	ss := strings.Split(s, ":")
	ip := ss[0]
	port, err := strconv.ParseInt(ss[1], 10, 32)
	if err != nil {
		panic(err)
	}

	extensionPort := make(map[string]int32)

	for _, v := range strings.Split(strings.Split(s, "?")[1], "&") {
		item := strings.Split(v, "=")
		p, err := strconv.ParseInt(strings.TrimSpace(item[1]), 10, 32)

		if err != nil {
			panic(err)
		}

		extensionPort[strings.TrimSpace(item[0])] = int32(p)
	}

	return &Member{
		Ip:             ip,
		Port:           int32(port),
		ExtensionPorts: extensionPort,
		MetaData:       make(map[string]string),
	}
}

func KRandomMember(k int, members []*Member, filter func(m *Member) bool) []*Member {

	totalSize := len(members)
	set := utils.NewSet()

	for i := 0; i < 3*totalSize && set.Size() <= k; i++ {
		idx := rand.Intn(totalSize)

		m := members[idx]
		if filter(m) {
			set.Add(m)
		}
	}

	ms := make([]*Member, set.Size())

	index := 0

	set.Range(func(value interface{}) {
		ms[index] = value.(*Member)
		index++
	})

	return ms
}
