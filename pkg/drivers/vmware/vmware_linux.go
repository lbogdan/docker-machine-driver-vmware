/*
Copyright 2017 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vmware

import (
	"os/exec"
	"syscall"
)

func DhcpConfigFiles() string {
	return "/etc/vmware/vmnet*/dhcpd/dhcpd.conf"
}

func DhcpLeaseFiles() string {
	return "/etc/vmware/vmnet*/dhcpd/dhcpd.leases"
}

func SetUmask() {
	_ = syscall.Umask(022)
}

// detect the vmrun and vmware-vdiskmanager cmds' path if needed
func setVmwareCmd(cmd string) string {
	if path, err := exec.LookPath(cmd); err == nil {
		return path
	}
	return cmd
}
