package main

import (
	"context"
	"os"

	"github.com/c2micro/mlan/cmd/mlan/internal/cmd"
	"github.com/c2micro/mlan/internal/zapcfg"
	"github.com/fatih/color"
	"github.com/go-faster/sdk/zctx"
	"github.com/spf13/cobra"
)

func main() {
	var err error

	lg, err := zapcfg.New().Build()
	if err != nil {
		panic(err)
	}

	flush := func() {
		_ = lg.Sync()
	}

	exit := func(code int) {
		flush()
		os.Exit(code)
	}

	a := cmd.App{}
	ctx := zctx.Base(context.Background(), lg)

	root := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Use:           "mlan",
		Short:         "mlan executor",
		Args:          cobra.NoArgs,
		RunE:          a.Run,
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			flush()
		},
	}

	// отключаем автокомплит
	root.CompletionOptions.DisableDefaultCmd = true
	// регаем глобальные флаги
	a.Globals.RegisterFlags(root.PersistentFlags())

	// стартуем
	if err = root.ExecuteContext(ctx); err != nil {
		color.Red("%v", err)
		exit(2)
	}
}
