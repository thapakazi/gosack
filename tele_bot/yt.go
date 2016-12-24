package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/yanzay/tbot"
)

func main() {

	//get token first
	token := os.Getenv("BOT_TOKEN")
	//get a bot with token
	bot, err := tbot.NewServer(token)
	if err != nil {
		panic(err)
	}

	// whitelist who can download stuffs :D
	users_i_only_serve_to := os.Getenv("MY_MASTERS")
	whitelist := strings.Split(users_i_only_serve_to, ",")
	bot.AddMiddleware(tbot.NewAuth(whitelist))

	// yo yo
	bot.Handle("yo", "yo man, whatsup")

	// lets run shell command
	bot.HandleFunc("/cmd {cmd}", ExecuteShell, "execute some basic terminal stuffs")
	bot.HandleFunc("/yt {url}", Youtubedl, "downloading files with youtube-dl")

	// Start listening for messages
	err = bot.ListenAndServe()
	log.Fatal(err)
}

// simple handy function found on web, thanks i forgot your name
func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func ExecuteShell(message tbot.Message) {
	cmd_raw := message.Vars["cmd"]
	command_to_exec := strings.Split(cmd_raw, ",")

	// fmt.Println("executing: ", command_to_exec)
	output := execute_command(command_to_exec[0], command_to_exec[1:]...)
	message.Reply(string(output))
}

func Youtubedl(message tbot.Message) {
	url_to_download := message.Vars["url"]
	download_signature := []string{os.Getenv("YOUTUBE_DL_BINARY_PATH")}
	command_to_exec := append(download_signature, url_to_download)
	output := execute_command(command_to_exec[0], command_to_exec[1:]...)
	message.Reply(string(output))
}

// this function does most of the execution stuffs
// wip
func execute_command(binary string, args ...string) (output []byte) {
	//combine std{out,err}

	cmd := exec.Command(binary, args...)
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)
	return output
}
