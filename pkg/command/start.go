package command

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/weaveworks-liquidmetal/microvm-action-runner/pkg/config"
	"github.com/weaveworks-liquidmetal/microvm-action-runner/pkg/flags"
	"github.com/weaveworks-liquidmetal/microvm-action-runner/pkg/handler"
	"github.com/weaveworks-liquidmetal/microvm-action-runner/pkg/payload"
)

func startCommand() *cli.Command {
	cfg := &config.Config{}

	return &cli.Command{
		Name:    "start",
		Usage:   "start the service",
		Aliases: []string{"s"},
		Before:  flags.ParseFlags(cfg),
		Flags: flags.CLIFlags(
			flags.WithHostFlag(),
			flags.WithAPITokenFlag(),
			flags.WithWebhookSecretFlag(),
			flags.WithSSHPublicKeyFlag(),
		),
		Action: func(c *cli.Context) error {
			return StartFn(cfg)
		},
	}
}

func StartFn(cfg *config.Config) error {
	// TODO: configurable logging levels
	log := logrus.NewEntry(logrus.StandardLogger())

	p := handler.Params{
		Config:  cfg,
		L:       log,
		Payload: payload.New(cfg.WebhookSecret),
		Client:  handler.NewFlintClient,
	}

	h, err := handler.New(p)
	if err != nil {
		return err
	}
	http.HandleFunc("/webhook", h.HandleWebhookPost)

	log.Infof("starting service on %s", cfg.Host)

	// TODO configurable port
	return http.ListenAndServe(":3000", nil)
}
