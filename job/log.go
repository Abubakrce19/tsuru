// Copyright 2022 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package job

import (
	"context"

	jobTypes "github.com/tsuru/tsuru/types/job"
)

type k8sJobLogService struct{}

func (k *k8sJobLogService) Watch(ctx context.Context, args jobTypes.ListLogArgs) error {
	// j, err := servicemanager.Job.GetByName(ctx, args.JobName)
	// if err != nil {
	// 	return nil, err
	// }
	return nil
}

func (k *k8sJobLogService) List(ctx context.Context, args jobTypes.ListLogArgs) error {
	return nil
}

// func (k *k8sJobLogService) Watch(ctx context.Context, args appTypes.ListLogArgs) (appTypes.LogWatcher, error) {
// 	a, err := servicemanager.App.GetByName(ctx, args.AppName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	tsuruWatcher, err := k.logService.Watch(ctx, args)
// 	if err != nil {
// 		return nil, err
// 	}

// 	logsProvisioner, err := k.provisionerGetter(ctx, a)
// 	if err == provision.ErrLogsUnavailable {
// 		return tsuruWatcher, nil
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	provisionerWatcher, err := logsProvisioner.WatchLogs(ctx, a, args)
// 	if err == provision.ErrLogsUnavailable {
// 		return tsuruWatcher, nil
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	return newMultiWatcher(provisionerWatcher, tsuruWatcher), nil
// }
