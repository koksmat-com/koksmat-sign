package main

import (
	"runtime/debug"
	"strings"

	"github.com/365admin/koksmat-sign/magicapp"
	"github.com/365admin/koksmat-sign/utils"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	// split info.Main.Path by / and get the last element
	s1 := strings.Split(info.Main.Path, "/")
	name := s1[len(s1)-1]
	description := `---
title: koksmat-sign
description: Describe the main purpose of this kitchen
---

# koksmat-sign
`
	magicapp.Setup(".env")
	magicapp.RegisterServeCmd("koksmat-sign", description, "0.0.1", 8080)
	magicapp.RegisterCmds()
	magicapp.RootCmd.PersistentFlags().BoolVarP(&utils.Verbose, "verbose", "v", false, "verbose output")

	magicapp.Execute(name, "koksmat-sign", "")
}
