package pomo

import (
	"fmt"
	"time"

	"github.com/rwxrob/cmdtab"
	"github.com/rwxrob/conf-go"
)

func init() {
	x := cmdtab.New("pomo")
	x.Summary = `sets or prints a countdown timer (with tomato)`
	x.Usage = `[<duration>|clear]`

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

		// [<duration>|clear]
		if len(args) > 0 {
			switch args[0] {
			case "clear":
				config.Set("pomo.up", "")
			case "dur":
				// TODO valid the duration
				config.Set("pomo.dur", args[1])
			case "start":
				// TODO detect optional duration argument
				s := config.Get("pomo.dur")
				if s == "" {
					s = "25m"
					config.Set("pomo.dur", s)
				}
				dur, err := time.ParseDuration(s)
				if err != nil {
					return err
				}
				up := time.Now().Add(dur).Format(time.RFC3339)
				config.SetSave("pomo.up", up)
			default:
				return x.UsageError()
			}
			return nil
		}
		up, err := time.Parse(time.RFC3339, config.Get("pomo.up"))
		if err != nil {
			return err
		}
		fmt.Printf("🍅 %v\n", up.Sub(time.Now()).Round(time.Second))
		return nil
	}
}
