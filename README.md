Cavalier
========

Cavalier is a command line interface *generator* for Go applications.  Cavalier
is not a library like those other CLI packages.  It is an application that
generates your command line parsing code for you.

Given a list of exported functions, it will create a git-like CLI with
subcommands, using the function names as the subcommands, the function arguments
as the flags and parameters, and the comments as the command's help docs.  

For example:

// in github.com/natefinch/cavalier

// Parse parses the exported functions in target and generates a CLI.
func Parse(
	target string, 	  // a go file or a directory containing a go package
	plan9style bool,  // use plan 9 style flag [use unix style]
	mainFunc *string, // name of "main" function [default: Main]
) error {}

Creates a command line that looks like:

$ cavalier help

Commands:
	Parse

$ cavalier help parse
	Parses the exported functions in target and generates a CLI.
	Target is either a go file or a directory containing a go package.
usage:
  	cavalier parse <target> [options]

options:
	--plan9style		use plan 9 style flag (use unix style)
	--mainFunc=string 	name of "main" function


By default, Cavalier inspects all files in the main package of your application
and treats each exported function as a subcommand.  A function called Main is
assumed to be the default target for commands that do not have a subcommand.

This behavior can be modified in several ways at generation time:

- Instead of package main, you can specify another package (generally in a
  subdirectory) that contains the methods that should be used for your CLI.
  This is a good practice, because it means your application's logic can be
  reused by others as if it is a library.

- Instead of a whole package, you can specify a specific file that contains the
  functions for which you wish to generate a CLI.  In this way, you can use only
  part of a package as your CLI.

- The name of the "Main" function is configurable.

- The flag behavior can be toggled between the std lib's plan9 style and unix
  style (i.e. if -xvf is the same as -x -v -f or --xvf).

Large quantities of flags can be hidden away as fields in a struct parameter,
which can be freely combined with non-struct parameters.


 
