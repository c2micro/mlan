package cmd

import (
	"os"
	"runtime/pprof"

	"github.com/c2micro/mlan/pkg/engine"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type App struct {
	Globals Globals
}

type Globals struct {
	File       string
	CpuProfile bool
}

func (g *Globals) RegisterFlags(f *pflag.FlagSet) {
	f.StringVar(&g.File, "file", "", "")
	f.BoolVar(&g.CpuProfile, "cpuprofile", false, "")
}

func (a *App) Run(cmd *cobra.Command, _ []string) error {
	// профилировка
	if a.Globals.CpuProfile {
		f, err := os.Create("cpu.pprof")
		if err != nil {
			return err
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			return err
		}
		defer pprof.StopCPUProfile()
	}
	// запуск программы
	err := engine.Evaluate(a.Globals.File)
	if err != nil {
		return err
	}
	return nil
}