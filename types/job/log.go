// Copyright 2020 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package job

import (
	"time"

	"github.com/tsuru/tsuru/types/auth"
	"gopkg.in/mgo.v2/bson"
)

type ListLogArgs struct {
	JobName string
	JobTeam string
	Units   []string
	Limit   int
	Token   auth.Token
}

type LogWatcher interface {
	Chan() <-chan JobLog
	Close()
}

type JobLog struct {
	MongoID bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Date    time.Time
	Message string
	Source  string
	AppName string
	Unit    string
}
