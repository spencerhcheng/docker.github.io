package main

import (
	"os/exec"
	"strings"

	"github.com/go-check/check"
)

func (s *DockerSuite) TestKillContainer(c *check.C) {
	runCmd := exec.Command(dockerBinary, "run", "-d", "busybox", "top")
	out, _, err := runCommandWithOutput(runCmd)
	if err != nil {
		c.Fatal(out, err)
	}

	cleanedContainerID := strings.TrimSpace(out)
	c.Assert(waitRun(cleanedContainerID), check.IsNil)

	killCmd := exec.Command(dockerBinary, "kill", cleanedContainerID)
	if out, _, err = runCommandWithOutput(killCmd); err != nil {
		c.Fatalf("failed to kill container: %s, %v", out, err)
	}

	listRunningContainersCmd := exec.Command(dockerBinary, "ps", "-q")
	out, _, err = runCommandWithOutput(listRunningContainersCmd)
	if err != nil {
		c.Fatalf("failed to list running containers: %s, %v", out, err)
	}

	if strings.Contains(out, cleanedContainerID) {
		c.Fatal("killed container is still running")
	}
}

func (s *DockerSuite) TestKillofStoppedContainer(c *check.C) {
	runCmd := exec.Command(dockerBinary, "run", "-d", "busybox", "top")
	out, _, err := runCommandWithOutput(runCmd)
	c.Assert(err, check.IsNil)

	cleanedContainerID := strings.TrimSpace(out)

	stopCmd := exec.Command(dockerBinary, "stop", cleanedContainerID)
	out, _, err = runCommandWithOutput(stopCmd)
	c.Assert(err, check.IsNil, check.Commentf("failed to stop container: %s, %v", out, err))

	killCmd := exec.Command(dockerBinary, "kill", "-s", "30", cleanedContainerID)
	_, _, err = runCommandWithOutput(killCmd)
	c.Assert(err, check.Not(check.IsNil), check.Commentf("Kill succeeded on a stopped container"))
}

func (s *DockerSuite) TestKillDifferentUserContainer(c *check.C) {
	runCmd := exec.Command(dockerBinary, "run", "-u", "daemon", "-d", "busybox", "top")
	out, _, err := runCommandWithOutput(runCmd)
	if err != nil {
		c.Fatal(out, err)
	}

	cleanedContainerID := strings.TrimSpace(out)
	c.Assert(waitRun(cleanedContainerID), check.IsNil)

	killCmd := exec.Command(dockerBinary, "kill", cleanedContainerID)
	if out, _, err = runCommandWithOutput(killCmd); err != nil {
		c.Fatalf("failed to kill container: %s, %v", out, err)
	}

	listRunningContainersCmd := exec.Command(dockerBinary, "ps", "-q")
	out, _, err = runCommandWithOutput(listRunningContainersCmd)
	if err != nil {
		c.Fatalf("failed to list running containers: %s, %v", out, err)
	}

	if strings.Contains(out, cleanedContainerID) {
		c.Fatal("killed container is still running")
	}
}
