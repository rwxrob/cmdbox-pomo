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

	x := cmdbox.Add("pomo", "show", "start", "stop",
		"duration", "emoji", "help", "version", "file")
	x.Default = "pomo show"
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

	x = cmdbox.Add("pomo blink")
	x.Usage = `[NEW]`
	x.Summary = `show or set the current blinking emoji`
	x.Method = func(args ...string) error {
		if len(args) == 0 {
			fmt.Println(config.Get("pomo.emoji.blink"))
		} else {
			config.Set("pomo.emoji.blink", args[0])
		}
		return nil
	}

	x = cmdbox.Add("pomo emoji")
	x.Usage = `[NEW]`
	x.Summary = `show or set the current emoji`
	x.Method = func(args ...string) error {
		if len(args) == 0 {
			fmt.Println(config.Get("pomo.emoji"))
		} else {
			config.Set("pomo.emoji", args[0])
		}
		return nil
	}

	x = cmdbox.Add("pomo stop")
	x.Summary = `stop the pomo timer without resetting`
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
		if emoji == "" {
			emoji = "🍅"
		}
		blinkEmoji := config.Get("pomo.emoji.blink")
		timeLeft := endt.Sub(time.Now()).Round(time.Second)
		if timeLeft < time.Second*30 && timeLeft%(time.Second*2) == 0 {
			fmt.Printf("%v %v\n", blinkEmoji, timeLeft)
			return nil
		}
		fmt.Printf("%v %v\n", emoji, timeLeft)
		return nil
	}

	x = cmdbox.Add("pomo start")
	x.Summary = `start the current pomo timer (without showing)`
	x.Method = func(args ...string) error {
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
		config.Set("pomo.up", up)
		return nil
	}

	x = cmdbox.Add("pomo duration")
	x.Usage = "[NEW]"
	x.Summary = `show duration or set new and start`
	x.Method = func(args ...string) error {
		if len(args) > 0 {
			config.Set("pomo.duration", args[0])
			return x.Call("pomo start")
		}
		fmt.Println(config.Get("pomo.duration"))
		return nil
	}

}
