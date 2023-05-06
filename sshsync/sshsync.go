package sshsync

import (
    "fmt"
    "golang.org/x/crypto/ssh"
    "strings"
)

type Options struct {
    SourceDir      string
    DestinationDir string
    Username       string
    Password       string
    Host           string
}

func (o *Options) Validate() error {
    if o.SourceDir == "" {
        return fmt.Errorf("source directory is required")
    }
    if o.DestinationDir == "" {
        return fmt.Errorf("destination directory is required")
    }
    if o.Username == "" {
        return fmt.Errorf("username is required")
    }
    if o.Password == "" {
        return fmt.Errorf("password is required")
    }
    if o.Host == "" {
        return fmt.Errorf("host is required")
    }
    return nil
}

func SyncFiles(opts *Options) error {
    if err := opts.Validate(); err != nil {
        return err
    }

    config := &ssh.ClientConfig{
        User: opts.Username,
        Auth: []ssh.AuthMethod{
            ssh.Password(opts.Password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    conn, err := ssh.Dial("tcp", opts.Host, config)
    if err != nil {
        return fmt.Errorf("failed to connect to host %s: %v", opts.Host, err)
    }
    defer conn.Close()

    session, err := conn.NewSession()
    if err != nil {
        return fmt.Errorf("failed to create session: %v", err)
    }
    defer session.Close()

    destDirParts := strings.Split(opts.DestinationDir, ":")
    var cmd string
    if len(destDirParts) > 1 {
        // The destination directory is on another host
        destHost := destDirParts[0]
        destPath := destDirParts[1]
        cmd = fmt.Sprintf("rsync -avz --delete %s %s:%s", opts.SourceDir, destHost, destPath)
		fmt.Println(cmd)

    } else {
        cmd = fmt.Sprintf("rsync -avz --delete %s %s", opts.SourceDir, opts.DestinationDir)
    }

    err = session.Run(cmd)
    if err != nil {
        return fmt.Errorf("failed to run command: %v", err)
    }

    return nil
}


