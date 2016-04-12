package main

import (
  "fmt"
  "os"
  "time"
)

import (
  "github.com/codegangsta/cli"
  "gopkg.in/cheggaaa/pb.v1"
)

func total(c *cli.Context) (t int64) {
	t += int64(c.Int("seconds"))
	t += int64(c.Int("minutes") * 60)
	t += int64(c.Int("hours") * 3600)
	t += int64(c.Int("days") * 86400)
	return
}

func run(c *cli.Context) {
	t := total(c)

	bar := pb.New64(t)
  bar.ShowPercent = false
	bar.SetRefreshRate(500)
	bar.ShowSpeed = false
	bar.Start()

	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			bar.Increment()
		}
	}()

	time.Sleep(time.Duration(t) * time.Second)
	ticker.Stop()

	bar.Finish()
	fmt.Println(c.String("message"))
}

func main() {
	app := cli.NewApp()
	app.Name = "Timer"
	app.Usage = "A simple timer"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "days",
			Usage: "Number of days",
			Value: 0,
		},
		cli.IntFlag{
			Name:  "hours",
			Usage: "Number of hours",
			Value: 0,
		},
		cli.IntFlag{
			Name:  "minutes",
			Usage: "Number of minutes",
			Value: 0,
		},
		cli.IntFlag{
			Name:  "seconds",
			Usage: "Number of seconds",
			Value: 1,
		},
		cli.StringFlag{
			Name:  "message",
			Usage: "Message to print when the timer's finished",
			Value: "Time's up!",
		},
	}
	app.Action = run
	app.Run(os.Args)
}