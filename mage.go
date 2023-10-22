//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// package main is the magefile for this project. To install reference https://github.com/magefile/mage
// Recommend using go install github.com/magefile/mage@latest to install mage if you have go install on your machine

func Docs() error {
	return sh.RunV("terraform-docs", "markdown", "table", ".", "--output-file", "README.md")
}

var Default = Docs
