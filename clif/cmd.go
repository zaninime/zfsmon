package clif

import (
	"bytes"
	"os/exec"
)

type ZpoolCommand struct {
	path string
}

func NewZpoolCommand(path string) *ZpoolCommand {
	return &ZpoolCommand{
		path: path,
	}
}

func (cmd *ZpoolCommand) exec(args []string) ([]string, error) {
	exe := exec.Command(cmd.path, args...)
	var out bytes.Buffer
	exe.Stdout = &out

	err := exe.Run()

	if err != nil {
		return nil, err
	}

	return ParseListOutput(&out), nil
}

func NewDefaultZpoolCommand() *ZpoolCommand {
	return NewZpoolCommand("zpool")
}

func (cmd *ZpoolCommand) listArgs(option string) []string {
	return []string{"list", "-H", "-o", option}
}

func (cmd *ZpoolCommand) listByPoolNameArgs(poolName string, option string) []string {
	unfilteredArgs := cmd.listArgs(option)

	return append(unfilteredArgs, poolName)
}

func (cmd *ZpoolCommand) List(option string) ([]string, error) {
	args := cmd.listArgs(option)

	return cmd.exec(args)
}

func (cmd *ZpoolCommand) ListByPoolName(poolName, option string) (*string, error) {
	args := cmd.listByPoolNameArgs(poolName, option)
	res, err := cmd.exec(args)

	if err != nil {
		return nil, err
	}

	return &res[0], nil
}

func (cmd *ZpoolCommand) ListAllByPoolName(option string) (map[string]string, error) {
	pools, err := cmd.List("name")

	if err != nil {
		return nil, err
	}

	poolMap := map[string]string{}

	for _, pool := range pools {
		optionValue, err := cmd.ListByPoolName(pool, option)

		if err != nil {
			return nil, err
		}

		poolMap[pool] = *optionValue
	}

	return poolMap, nil
}
