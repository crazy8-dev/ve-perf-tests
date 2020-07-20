package ve_perf_tests

import (
	"context"
	"errors"
	"net/http"

	"github.com/insolar/loaderbot"

	"github.com/insolar/ve-perf-tests/util"
)

type GetContractTestAttack struct {
	*loaderbot.Runner
	client *http.Client
}

func (a *GetContractTestAttack) Setup(cfg loaderbot.RunnerConfig) error {
	a.client = loaderbot.NewLoggingHTTPClient(cfg.DumpTransport, 60)
	return nil
}
func (a *GetContractTestAttack) Do(_ context.Context) loaderbot.DoResult {
	url := a.Cfg.TargetUrl + util.WalletGetBalancePath
	ref := util.GetNextWallet(a.Runner)
	balance, err := util.GetWalletBalance(a.client, url, ref)
	if err != nil {
		return loaderbot.DoResult{
			Error:        err,
			RequestLabel: a.Name,
		}
	}
	if balance != util.StartBalance {
		return loaderbot.DoResult{
			Error:        errors.New("balance is not equal to start balance"),
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
