package sshsyncer

import (
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"golang.org/x/crypto/ssh"
)

// copyFiles copies files from a remote server to a local destination directory
func copyFiles(sshConfig *ssh.ClientConfig, remoteSrcDir string, localDestDir string) error {
	// Connect to remote SSH server
	sshClient, err := ssh.Dial("tcp", "remote-server:22", sshConfig)
	if err != nil {
		return err
	}
	defer sshClient.Close()

	// Get list of files in remote source directory
	sshSession, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer sshSession.Close()

	cmd := "ls " + remoteSrcDir
	output, err := sshSession.Output(cmd)
	if err != nil {
		return err
	}

	fileList := strings.Split(string(output), "\n")

	// Create a wait group to wait for all copy processes to finish
	var wg sync.WaitGroup

	// Copy each file in a separate goroutine
	for _, file := range fileList {
		if file != "" {
			remoteFilePath := remoteSrcDir + "/" + file
			localFilePath := localDestDir + "/" + file

			wg.Add(1)
			go func() {
				defer wg.Done()

				// Open SSH session and remote file
				session, err := sshClient.NewSession()
				if err != nil {
					log.Printf("Failed to create SSH session for file %s: %s\n", file, err)
					return
				}
				defer session.Close()

				srcFile, err := session.Open(remoteFilePath)
				if err != nil {
					log.Printf("Failed to open remote file %s: %s\n", remoteFilePath, err)
					return
				}
				defer srcFile.Close()

				// Create local file and copy contents
				destFile, err := os.Create(localFilePath)
				if err != nil {
					log.Printf("Failed to create local file %s: %s\n", localFilePath, err)
					return
				}
				defer destFile.Close()

				_, err = io.Copy(destFile, srcFile)
				if err != nil {
					log.Printf("Failed to copy file %s: %s\n", file, err)
					return
				}

				log.Printf("Successfully copied file %s\n", file)
			}()
		}
	}

	// Wait for all copy processes to finish
	wg.Wait()

	return nil
}
