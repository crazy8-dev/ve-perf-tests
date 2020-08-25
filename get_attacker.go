package ve_perf_tests

import (
	"context"

	"github.com/insolar/loaderbot"

	"github.com/insolar/ve-perf-tests/util"
)

type GetContractTestAttack struct {
	*loaderbot.Runner
}

func (a *GetContractTestAttack) Setup(cfg loaderbot.RunnerConfig) error {
	return nil
}
func (a *GetContractTestAttack) Do(_ context.Context) loaderbot.DoResult {
	sw := a.TestData.(*loaderbot.SharedDataSlice).Get().(util.StickyWallet)
	url := sw.Url + util.WalletGetBalancePath
	ref := sw.Ref
	if _, err := util.GetWalletBalance(a.HTTPClient, url, ref); err != nil {
		return loaderbot.DoResult{
			Error:        err.Error(),
			RequestLabel: a.Name,
		}
	}
	return loaderbot.DoResult{
		RequestLabel: a.Name,
	}
}
func (a *GetContractTestAttack) Clone(r *loaderbot.Runner) loaderbot.Attack {
	return &GetContractTestAttack{Runner: r}
}

func (a *GetContractTestAttack) Teardown() error {
	return nil
}
