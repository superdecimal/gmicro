package cmd

import (
	"context"
	"fmt"
	"strconv"
	"time"

	gmrpc "superdecimal/gmicro/pkg/proto"
	"superdecimal/gmicro/services/cli/config"

	"google.golang.org/grpc"
	ishell "gopkg.in/abiosoft/ishell.v2"

	"github.com/fatih/color"
)

func CalcCommands(config *config.Configuration) *ishell.Cmd {
	red := color.New(color.FgRed).SprintFunc()
	srvAddr := fmt.Sprintf("%s:%d", config.Address, config.Port)
	calc := &ishell.Cmd{
		Name: "calc",
		Help: "CalcAPI operations",
	}

	calc.AddCmd(
		&ishell.Cmd{
			Name: "add",
			Help: "Adds two numbers",
			Func: func(c *ishell.Context) {
				c.ShowPrompt(false)
				defer c.ShowPrompt(true)

				ctx := context.Background()
				conn, err := grpc.Dial(srvAddr, grpc.WithInsecure())
				if err != nil {
					c.Println(red(err))
					return
				}
				defer conn.Close()

				gmclient := gmrpc.NewCalculatorAPIClient(conn)

				c.Println("Number a: ")
				a := c.ReadLine()
				ai, err := strconv.ParseInt(a, 10, 32)
				if err != nil {
					c.Println("Not a number: " + red(a))
					return
				}

				c.Println("Number b: ")
				b := c.ReadLine()
				bi, err := strconv.ParseInt(b, 10, 32)
				if err != nil {
					c.Println("Not a number: " + red(b))
					return
				}

				startTime := time.Now()

				resp, err := gmclient.Add(ctx, &gmrpc.AddRequest{
					A: int32(ai),
					B: int32(bi),
				})
				if err != nil {
					c.Println(red(err))
				}

				endTime := time.Now()

				c.Println("Done, Result: ", resp.GetResult(), " time: ",
					endTime.Sub(startTime))
			},
		})

	calc.AddCmd(
		&ishell.Cmd{
			Name: "sum",
			Help: "Sums numbers until eof (type eof to stop the stream)",
			Func: func(c *ishell.Context) {
				c.ShowPrompt(false)
				defer c.ShowPrompt(true)

				ctx := context.Background()

				conn, err := grpc.Dial(srvAddr, grpc.WithInsecure())
				if err != nil {
					c.Println(red(err))
					return
				}
				defer conn.Close()

				gmclient := gmrpc.NewCalculatorAPIClient(conn)

				stream, err := gmclient.Sum(ctx)
				if err != nil {
					c.Println(red(err))
					return
				}

				for {
					c.Println("Number: ")
					a := c.ReadLine()
					if a == "eof" {
						resp, err := stream.CloseAndRecv()
						if err != nil {
							c.Println(red(err))
						}
						c.Println("Done, Result: ", resp.GetResult())
						break
					}

					ai, err := strconv.ParseInt(a, 10, 32)
					if err != nil {
						c.Println("Not a number: " + red(a))
						return
					}

					if err := stream.Send(
						&gmrpc.Integer{Num: int32(ai)},
					); err != nil {
						c.Println(red(err))
						return
					}
				}
			},
		})

	return calc
}
