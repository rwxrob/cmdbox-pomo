package cmd

import (
	"fmt"
	"time"

	"github.com/rwxrob/cmdtab"
	"github.com/rwxrob/conf-go"
)

func init() {
	x := cmdtab.New("pomo", "start", "stop", "duration")
	x.Summary = `sets or prints a countdown timer (with tomato)`
	x.Usage = `[start|stop|duration]`

	x.Description = `
		If a Go time duration is passed then sets the pomo.start config
		value. If no value is passed prints a tomato emoji followed by the
		duration recmdsing. If <clear> is passed sets pomo.start to empty
		string.`

	// TODO let people set the emojis if different from defaults
	// TODO add emoji for when over time by whatever amount
	// TODO add like 5 emojis for past the deadline
	// TODO consider adding blink support if detected in terminal

	x.Method = func(args []string) (err error) {
		config := conf.New()
		err = config.Load()
		if err != nil {
			return err
		}

		if len(args) > 0 {
			switch args[0] {
			case "stop":
				config.SetSave("pomo.end", "")
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
				end := time.Now().Add(dur).Format(time.RFC3339)
				config.SetSave("pomo.end", end)
			default:
				return x.UsageError()
			}
			return nil
		}
		end := config.Get("pomo.end")
		if end == "" {
			return nil
		}
		endt, err := time.Parse(time.RFC3339, end)
		if err != nil {
			return err
		}
		fmt.Printf("🍅 %v\n", endt.Sub(time.Now()).Round(time.Second))
		return nil
	}
}
