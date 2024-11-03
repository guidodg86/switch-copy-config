package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

// need to replace session by shell
func main() {
	const SOURCE_SWITCH = "192.168.201.20:22"

	// Create connection options for source switch
	source_switch_config := &ssh.ClientConfig{
		User: "cisco",
		Auth: []ssh.AuthMethod{
			ssh.Password("cisco"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// Connect to source switch
	client_source, err := ssh.Dial("tcp", SOURCE_SWITCH, source_switch_config)
	if err != nil {
		log.Fatal("Failed to connect to source switch: ", err)
	}
	defer client_source.Close()

	// Start  new session with source switch
	session_source, err := client_source.NewSession()
	if err != nil {
		log.Fatal("Failed to create new session in source switch: ", err)
	}
	defer session_source.Close()

	// Get running config
	var running_config bytes.Buffer
	session_source.Stdout = &running_config
	if err := session_source.Run("show running-config"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	session_source.Close()

	session_source2, err := client_source.NewSession()
	if err != nil {
		log.Fatal("Failed to create new session in source switch: ", err)
	}
	defer session_source2.Close()

	// Get vlan database
	var vlan_database bytes.Buffer
	session_source2.Stdout = &running_config
	if err := session_source2.Run("show vlan brief"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}

	fmt.Println(vlan_database.String())
	var _ = 2
}
