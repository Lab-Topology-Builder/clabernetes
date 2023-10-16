package launcher

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"

	clabernetesconstants "github.com/srl-labs/clabernetes/constants"
	claberneteserrors "github.com/srl-labs/clabernetes/errors"
)

func (c *clabernetes) runClab() error {
	clabLogFile, err := os.Create("clab.log")
	if err != nil {
		return err
	}

	clabOutWriter := io.MultiWriter(c.logger, clabLogFile)

	args := []string{
		"deploy",
		"-t",
		"topo.clab.yaml",
	}

	if os.Getenv(clabernetesconstants.LauncherContainerlabDebug) == clabernetesconstants.True {
		args = append(args, "--debug")
	}

	cmd := exec.Command("containerlab", args...)

	cmd.Stdout = clabOutWriter
	cmd.Stderr = clabOutWriter

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (c *clabernetes) runClabVxlanTools(
	localNodeName, cntLink, vxlanRemote string,
	vxlanID int,
) error {
	resolvedVxlanRemotes, err := net.LookupIP(vxlanRemote)
	if err != nil {
		return err
	}

	if len(resolvedVxlanRemotes) != 1 {
		return fmt.Errorf(
			"%w: did not get exactly one ip resolved for remote vxlan endpoint",
			claberneteserrors.ErrConnectivity,
		)
	}

	resolvedVxlanRemote := resolvedVxlanRemotes[0].String()

	c.logger.Debugf("resolved remote vxlan tunnel service address as '%s'", resolvedVxlanRemote)

	cmd := exec.Command( //nolint:gosec
		"containerlab",
		"tools",
		"vxlan",
		"create",
		"--remote",
		resolvedVxlanRemote,
		"--id",
		fmt.Sprint(vxlanID),
		"--link",
		fmt.Sprintf("%s-%s", localNodeName, cntLink),
		"--port",
		fmt.Sprint(clabernetesconstants.VXLANServicePort),
	)

	_, err = cmd.Output()
	if err != nil {
		return err
	}

	return nil
}