package cmd

import (
	"fmt"
	"time"

	"github.com/rwxrob/cmdtab"
	"github.com/rwxrob/conf-go"
)

func init() {
	x := cmdtab.New("pomo", "start", "stop", "duration", "emoji")
	x.Summary = `sets or prints a countdown timer (with tomato)`
	x.Usage = `[start|stop|duration|emoji]`

	x.Description = `
		The *pomo* command assists those with creating scripts and other
		tools to help them follow the simple Pomodoro method of time boxing.

		If no value is passed prints an emoji (default: tomato) followed by
		the duration remaining unless *pomo.start* is empty in which case it
		prints nothing allowing it to be called in a loop and included in
		other tools such as TMUX [set -g status-left "#(cmd pomo)"].

		If *start* is passed sets *pomo.start* to the current time and
		*pomo.up* to the time at which the current Pomodoro session expires.

		If *stop* is passed sets *pomo.start* to empty string and
		effectively disables printing anything.

		When *duration* is passed it will change *pomo.duration* and
		effective call *start* as well.  If no argument to duration is
		passed it will simply print it.

    When *emoji* is passed it will change *pomo.emoji* to the argument
    passed to *emoji*.

		When any subcommand or argument other than the above is passed the
		*duration* subcommand is called and passed the argument.

		If more than two arguments are ever passed prints usage error.`

	x.Method = func(args []string) (err error) {
		config := conf.New()
		err = config.Load()
		if err != nil {
			return err
		}

		if len(args) > 0 {
			switch args[0] {
			case "stop":
				config.SetSave("pomo.up", "")
			case "duration":
				config.SetSave("pomo.duration", args[1])
				fallthrough
			case "start":
				s := config.Get("pomo.duration")
				if s == "" {
					s = "25m"
					config.Set("pomo.duration", s)
				}
				dur, err := time.ParseDuration(s)
				if err != nil {
					return err
				}
				up := time.Now().Add(dur).Format(time.RFC3339)
				config.SetSave("pomo.up", up)
			case "emoji":
				config.SetSave("pomo.emoji", args[1])
			default:
				return x.UsageError()
			}
			return nil
		}
		up := config.Get("pomo.up")
		if up == "" {
			return nil
		}
		endt, err := time.Parse(time.RFC3339, up)
		if err != nil {
			return err
		}
		emoji := config.Get("pomo.emoji")
		if emoji == "" {
			emoji = "🍅"
		}
		fmt.Printf("%v %v\n", emoji, endt.Sub(time.Now()).Round(time.Second))
		return nil
	}
}
