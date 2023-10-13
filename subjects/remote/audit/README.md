#### Functional

###### Is virtualization software installed?

###### Is the VM, 01_remote, added to the virtualization software?

#### General

##### The auditeee is supposed to set a port forwarding rule in the VM settings that maps the host port to a guest port.

##### Ask the auditee to show you that rule in the settings or to set it up if it is not done yet.

##### Ask the auditee to connect to the VM via SSH thru the host port.

###### Did the auditee manage to connect via SSH with either this command: `ssh -p22 root@localhost` or `ssh -pANOTHER_PORT root@localhost`?

##### If the auditee connected thru port 22, ask the auditee to change the port of the Guest VM.

##### Ask the auditee to connect to the VM via SSH thru the new chosen host port.

###### Did the auditee manage to connect via SSH with this command `ssh -pANOTHER_PORT root@localhost`?
