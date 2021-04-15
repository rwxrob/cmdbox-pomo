package cmd

import (
	"fmt"
	"time"

	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/conf-go"
)

func init() {
	x := cmdbox.New("pomo", "start", "stop", "pause", "resume", "duration", "emoji", "help", "version", "file")
	x.Summary = `sets or prints a countdown timer (with tomato)`
	x.Usage = `[start|stop|duration|emoji|emoji.blink]`
	x.Version = `v1.0.0`

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

    When *emoji.blink* is passed it will change *pomo.emoji.blink* to the
    argument passed to *emoji.blink*.

		When any subcommand or argument other than the above is passed the
		*duration* subcommand is called and passed the argument.

		If more than two arguments are ever passed prints usage error.`

	x.Method = func(args []string) (err error) {
		config, err := conf.New()
		if err != nil {
			return err
		}
		err = config.Load()
		if err != nil {
			return err
		}

		if len(args) > 0 {
			switch args[0] {
			case "stop":
				config.SetSave("pomo.up", "")
				config.SetSave("pomo.status", "")
			case "duration":
				config.SetSave("pomo.duration", args[1])
				fallthrough
			case "start":
				s := config.Get("pomo.duration")
				if s == "" {
					s = "25m"
					config.Set("pomo.duration", s)
				}
				if err := run(s, config); err != nil {
					return err
				}
			case "pause":
				up := config.Get("pomo.up")
				if up == "" {
					return nil
				}
				endt, err := time.Parse(time.RFC3339, up)
				if err != nil {
					return err
				}
				config.SetSave("pomo.duration", endt.Sub(time.Now()).Round(time.Second).String())
				config.SetSave("pomo.status", "paused")
			case "resume":
				status := config.Get("pomo.status")
				if status == "" {
					return nil
				}
				dur := config.Get("pomo.duration")
				if dur == "" {
					return nil
				}
				if err := run(dur, config); err != nil {
					return err
				}
			case "emoji":
				config.SetSave("pomo.emoji", args[1])
			case "emoji.blink":
				config.SetSave("pomo.emoji.blink", args[1])
			case "file":
				fmt.Println(config.Path())
				return nil
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
		blinkEmoji := config.Get("pomo.emoji.blink")
		var timeLeft time.Duration
		status := config.Get("pomo.status")
		if status == "paused" {
			timeLeft, err = time.ParseDuration(config.Get("pomo.duration"))
			if err != nil {
				return err
			}
		} else {
			timeLeft = endt.Sub(time.Now()).Round(time.Second)
		}
		if timeLeft < time.Second*30 && timeLeft%(time.Second*2) == 0 {
			fmt.Printf("%v %v%s\n", blinkEmoji, timeLeft, displayStatus(status))
			return nil
		}
		fmt.Printf("%v %v%s\n", emoji, timeLeft, displayStatus(status))
		return nil
	}
}

func run(duration string, config *conf.Config) error {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}
	up := time.Now().Add(dur).Format(time.RFC3339)
	config.SetSave("pomo.up", up)
	config.SetSave("pomo.status", "")
	return nil
}

func displayStatus(s string) string {
	if s == "paused" {
		return " (PAUSED)"
	}
	return ""
}
