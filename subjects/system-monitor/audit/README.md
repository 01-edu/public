#### Functional

##### Try and run the application, and search for the system monitor.

###### Is the operating system name correct? (linux)

##### Try and run the application and search for the user logged in. Then run the command `"who"`.

###### Are both the users the same?

##### Try and run the application and search for the hostname. Then run the command `"hostname"`.

###### Are both the names the same?

##### Try and run the application and search for the total number of tasks/processes. Then run the command `"top"` and search for the "Tasks".

###### Are both the total number of tasks the same?

##### Try and run the application and search for the CPU type. Then run the command `"cat /proc/cpuinfo"`.

###### Is the CPU type provided by the application the same as the "model name" present in the cpuinfo file?

##### Try and run the application and search for the system monitor.

###### Can you confirm that there is a tabbed section?

###### And if so, does it have the following tabs : "CPU", "Fan", "Thermal"?

##### Try and run the application and open the "CPU" tab in the system monitor.

###### Is there a performance graph with the current CPU percentage?

###### And if so, does it have a working slide bar for the fps and the "y" scale?

###### Does it also have a way to stop the graph animation?

##### Try and run the application and open the "Thermal" tab in the system monitor.

###### Is there a performance graph?

###### And if so, does it have a working slide bar for the fps and the "y" scale?

###### Does it also have a way to stop the graph animation?

##### Try and run the command `"cat /proc/acpi/ibm/thermal"`.

###### Is the temperature the same as in the application?

##### Try and increase the computers temperature (without breaking anything), and look at the thermal graph.

###### Is the temperature increasing accordingly?

##### Try and run the application and open the "Fan" tab in the system monitor.

###### Is there a performance graph?

###### And if so, does it have a working slide bar for the fps and the "y" scale?

###### Does it also have a way to stop the graph animation?

##### Try and run the application and search for the memory and process monitor.

###### Can you confirm that there is a visual display of the memory usage?

##### Try and run the application and search for the RAM, then run the command `"free -h"` and compare the values.

###### Can you confirm that both usage and total(RAM) are the same as in the application?

##### Try and run the application and search for the SWAP, them run the command `"free -h"` and compare the values.

###### Can you confirm that both usage and total(SWAP) are the same as in the application?

##### Try and run the application and search fot the Disk, then run the command `"df -h /"` and compare the values.

###### Can you confirm that the disk size and usage are same as in the application?

##### Try and run the application and search for the Process Table.

###### Does the table present the columns PID, Name, State, CPU usage and Memory usage?

##### Try and run the application, then use the filter to search for the process `"monitor"`.

###### Were you able to filter the table?

##### Try and run the command `"top"`, then search for the process `"monitor"`.

###### Are the values from each column the same as in the command `"top"`?

##### Try and run the application, then select from the table three processes.

###### Was it possible to select three processes?

##### Try and run the application and search for the ip address, then run the command `"ifconfig"`.

###### Are the list of networks the same as in the application?

###### And if so, is the ip correct for each network?

##### Try and run the application and go to the RX (network receiver) table, then run the command `"cat /proc/net/dev"` and search for the column Receive.

###### Are the values from the table the same as in the application?

##### Try and run the application and go to the TX (network transmitter) table, then run the command `"cat /proc/net/dev"` and search for the column Transmit.

###### Are the values from the table the same as in the application?

##### Try and run the application and search for the tabbed section that contains the visual usage of the network and select the Receiver (RX).

###### Is there a visual representation for it?

##### Try and run the application and search for the tabbed section that contains the visual usage of the network and select the Receiver (RX). Then test your network download speed.

###### Is the value form the right network increasing while testing the download speed?

##### Try and run the application and search for the tabbed section that contains the visual usage of the network and select the Transmitter (TX). Then test your network upload speed.

###### Is the value form the right network increasing while testing the upload speed?

##### Try and run the application and search for the tabbed section that contains the visual usage of the network and select the Transmitter (TX).

###### Is there a visual representation for it?

##### Try and run the application and search for the tabbed section that contains the visual usage of the network and select the Receiver (RX), then run the command `"ifconfig"` and compare each network with the visual representation.

###### Are the values from the command the same as in the application?

##### Try and run the application and search for the tabbed section that contains the visual usage of the network and select the Transmitter (TX), then run the command `"ifconfig"` and compare each network with the visual representation.

###### Are the values from the command the same as in the application?

###### Are the values well converted for each network? (from bytes to GB, MB or KB, obeying the rules from the subject)
