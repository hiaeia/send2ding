package send2ding

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	"github.com/hiaeia/send2ding"
)

var markdown = &cobra.Command{
	Use:   "markdown",
	Short: "send dingtalk markdown message",
	Run: func(cmd *cobra.Command, args []string) {
		at, err := cmd.PersistentFlags().GetStringSlice("at")
		if err != nil {
			fmt.Println(err.Error())
		}
		title, err := cmd.PersistentFlags().GetString("title")
		if err != nil {
			fmt.Println(err.Error())
		}

		var text string
		if len(args) == 0 {
			b, _ := ioutil.ReadAll(os.Stdin)
			text = string(b)
		} else {
			text = args[0]
		}

		msg := send2ding.MarkdownMessage(title, text)

		if len(at) != 0 {
			err = msg.At(false, at...)
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		err = dingTalk.Send(msg)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			// 从stdin中获取数据
			f, err := os.Stdin.Stat()
			if err != nil {
				return err
			}
			if f.Size() == 0 {
				return errors.New("you need some message to send")
			}
		}
		return nil
	},
}

func init() {
	markdown.PersistentFlags().String("title", "", "markdown title")
	markdown.PersistentFlags().StringSlice("at", []string{}, "at some one")
}
