package tsuru

import (
	appTypes "github.com/tsuru/tsuru/types/app"
	jobTypes "github.com/tsuru/tsuru/types/job"
)

type TsuruObject interface {
	appTypes.App
	jobTypes.Job
}
