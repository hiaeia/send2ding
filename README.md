# Dingbot
DingTalk robot golang library, and command line

# Usage
## As library
```golang
package main

import (
	"github.com/hiaeia/send2ding"
)

func main() {
	token := "you ding talk access token"
	secret := "you ding talk secret, skip if old robot"

	text := "hello world"
	msg := send2ding.TextMessage(text)
	client := send2ding.New(token, secret)

	err := client.Send(msg)
	if err != nil {
		panic(err)
	}
}
```

## Command line
### preview
```shell script
send dingtalk message

Usage:
  send2ding [command]

Available Commands:
  feedcard    send dingtalk feedcard message
  help        Help about any command
  init        init config
  link        send dingtalk link message
  markdown    send dingtalk markdown message
  text        send dingtalk text message
  version     Print the version number of Dingbot

Flags:
      --config string   config file (default is $HOME/send2ding.toml)
  -h, --help            help for send2ding
      --secret string   dingtalk robot secret
      --token string    dingtalk robot token (require)

Use "send2ding [command] --help" for more information about a command.
```

### use Docker
```shell script
$ docker pull hiaeia/send2ding
$ docker run --rm hiaeia/send2ding send2ding --help
```

### Docker command alias
```shell script
$ alias send2ding="docker run --rm hiaeia/send2ding send2ding --token 'you token' --secret 'you secret'"
# then
$ send2ding text "hello world"
```

### build
```shell script
$ git clone https://github.com/hiaeia/send2ding.git
$ cd send2ding
$ go mod vendor
$ go build -o send2ding cmd/main.go
$ ./send2ding version
```

### send text message
```shell script
$ ./send2ding --token "you token" --secret "you secret" text "hello world"
# or
$ echo "hello world" | ./send2ding --token "you token" --secret "you secret" text
```

### send markdown message
```shell script
$ ./send2ding markdown --title hello '## hello world'
# or
$ echo "hello world" | ./send2ding --token "you token" --secret "you secret" markdown --title hello
```

### send link message
```shell script
$ ./send2ding --token "you token" --secret "you secret" link --title hello --message-url 'https://6cm.co' 'hello world'
# or
$ echo "hello world" | ./send2ding --token "you token" --secret "you secret" link --title hello --message-url 'https://6cm.co'
```

### send feed card message
```shell script
$ ./send2ding --token "you token" --secret "you secret" feedcard --json-link '{"title":"时代的火车向前开","messageURL":"https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI","picURL":"https://gw.alicdn.com/tfs/TB1ayl9mpYqK1RjSZLeXXbXppXa-170-62.png"}' --json-link '{"title":"时代的火车向前开","messageURL":"https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI","picURL":"https://gw.alicdn.com/tfs/TB1ayl9mpYqK1RjSZLeXXbXppXa-170-62.png"}'
```

### command alias
```shell script
# replace path/to to you send2ding path
$ alias send2ding="path/to/send2ding --token 'you token' --secret 'you secret'"
# then
$ send2ding text "hello world"
```

### use config 
```shell script
$ ./send2ding --token "you token" --secret "you secret" init dintbot.toml
$ ./send2ding --config dintbot.toml text "hello world"
```