/**
 * @Author: Hardews
 * @Date: 2022/11/3 0:39
**/

package bot

import (
	"fmt"
	"os/exec"
)

func Start() {
	command := `./go-cqhttp`
	cmd := exec.Command(command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}
