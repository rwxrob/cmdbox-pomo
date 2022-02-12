package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/conf-go"
)

func init() {

	config := conf.NewMap()
	if config == nil {
		log.Printf("failed to create configuration map")
		return
	}

	setDefaults := func() {
		v := config.Get("pomo.duration")
		if v == "" {
			v = "25m"
			config.Set("pomo.duration", v)
		}
		v = config.Get("pomo.warning")
		if v == "" {
			v = "1m"
			config.Set("pomo.warning", v)
		}
		v = config.Get("pomo.emoji")
		if v == "" {
			v = "🍅"
			config.Set("pomo.emoji", v)
		}
		v = config.Get("pomo.warning.emoji")
		if v == "" {
			v = "💢"
			config.Set("pomo.warning.emoji", v)
		}
	}

	x := cmdbox.Add("pomo", "show", "start", "stop",
		"d|dur|duration", "emoji", "we|warning.emoji", "w|warn|warning",
		"h|help", "v|version", "f|file", "e|vi|ed|edit", "up",
	)

	x.Default = "pomo show"
	x.Hidden = []string{"up"}
	x.Summary = `sets or prints a countdown timer (with tomato)`
	x.Version = `v2.0.0`
	x.Source = `https://github.com/rwxrob/cmdbox-pomo`
	x.Issues = `https://github.com/rwxrob/cmdbox-pomo/issues`
	x.Description = `
		The *pomo* command assists those with creating scripts and other
		tools to help them follow the simple Pomodoro method of time boxing.
		When called without arguments *pomo show* is assumed.`

	x = cmdbox.Add("pomo file")
	x.Summary = `show full path to configuration file`
	x.Method = func(args ...string) error {
		fmt.Println(config.Path())
		return nil
	}

	x = cmdbox.Add("pomo up")
	x.Summary = `show date/time started`
	x.Method = func(args ...string) error {
		fmt.Println(config.Get("pomo.up"))
		return nil
	}

	x = cmdbox.Add("pomo warning.emoji")
	x.Usage = `[NEW]`
	x.Summary = `show/set warning emoji`
	x.Method = func(args ...string) error {
		if len(args) == 0 {
			fmt.Println(config.Get("pomo.warning.emoji"))
		} else {
			config.Set("pomo warning.emoji", args[0])
		}
		return nil
	}

	x = cmdbox.Add("pomo emoji")
	x.Usage = `[NEW]`
	x.Summary = `show/set current emoji`
	x.Method = func(args ...string) error {
		if len(args) == 0 {
			fmt.Println(config.Get("pomo.emoji"))
		} else {
			config.Set("pomo.emoji", args[0])
		}
		return nil
	}

	x = cmdbox.Add("pomo stop")
	x.Summary = `stop pomo timer without resetting`
	x.Method = func(args ...string) error {
		config.Set("pomo.up", "")
		return nil
	}

	x = cmdbox.Add("pomo show")
	x.Summary = `show the current pomo timer (default)`
	x.Method = func(args ...string) error {
		up := config.Get("pomo.up")
		if up == "" {
			return nil
		}
		endt, err := time.Parse(time.RFC3339, up)
		if err != nil {
			return err
		}
		emoji := config.Get("pomo.emoji")
		warnEmoji := config.Get("pomo.warning.emoji")
		warning, err := time.ParseDuration(config.Get("pomo.warning"))
		if err != nil {
			return err
		}
		timeLeft := endt.Sub(time.Now()).Round(time.Second)
		if timeLeft < warning && timeLeft%(time.Second*2) == 0 {
			fmt.Printf("%v%v\n", warnEmoji, timeLeft)
			return nil
		}
		fmt.Printf("%v%v\n", emoji, timeLeft)
		return nil
	}

	x = cmdbox.Add("pomo start")
	x.Summary = `start the current pomo timer (without showing)`
	x.Method = func(args ...string) error {
		setDefaults()
		s := config.Get("pomo.duration")
		dur, err := time.ParseDuration(s)
		if err != nil {
			return err
		}
		up := time.Now().Add(dur).Format(time.RFC3339)
		config.Set("pomo.up", up)
		return nil
	}

	x = cmdbox.Add("pomo duration")
	x.Usage = "[NEW]"
	x.Summary = `show/set duration and start over`
	x.Method = func(args ...string) error {
		if len(args) > 0 {
			config.Set("pomo.duration", args[0])
			return x.Call("pomo start")
		}
		fmt.Println(config.Get("pomo.duration"))
		return nil
	}

	x = cmdbox.Add("pomo warning")
	x.Usage = "[NEW]"
	x.Summary = `show/set warning seconds remaining`
	x.Method = func(args ...string) error {
		if len(args) > 0 {
			config.Set("pomo.warning", args[0])
		}
		fmt.Println(config.Get("pomo.warning"))
		return nil
	}

	x = cmdbox.Add("pomo edit")
	x.Summary = `edit pomo configuration file`
	x.Method = func(args ...string) error {
		return config.Edit()
	}

}
