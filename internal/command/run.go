package command

import (
	"os/exec"

	"github.com/charmbracelet/log"
)

func Run(command ...string) (status int, err error) {
	status = 0
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = log.StandardLog(log.StandardLogOptions{
		ForceLevel: log.InfoLevel,
	}).Writer()
	cmd.Stderr = log.StandardLog(log.StandardLogOptions{
		ForceLevel: log.ErrorLevel,
	}).Writer()
	err = cmd.Run()
	if exitError, ok := err.(*exec.ExitError); ok {
		status = exitError.ExitCode()
	}
	return
}
