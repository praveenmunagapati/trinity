package commands

/*
RootHelp string is the main usage info string for
the executable
*/
const RootHelp = `
trinity - IPFS search engine

usage:

trinity <command> [--options]

commands:

  init        initialize the daemon
  role        set role for the node
  swarm       get list of peer nodes
  config      edit configuration settings


typing "help" after any command will give you more
details about that command's usage/options
`

/*
FIndexHelp is the help message displayed for the
forward index subsystem.
*/
const FIndexHelp = `
tfindex - trinity forward index subsystem

usage:

$: tfindex [path/to/config]

The path to a configuration file is optional. Providing
one will override the default config file from being loaded
`

/*
IIndexHelp is the help message displayed for the
inverted index subsystem.
*/
const IIndexHelp = `
tiindex - trinity inverted index subsystem

usage:

$: tiindex [path/to/config]

The path to a configuration file is optional. Providing
one will override the default config file from being loaded
`

/*
HelpText stores the info that will be displayed in
CLI help text (given the "help" command)
*/
type HelpText struct {
}

/*
Command is a core function that a user can invoke,
wherever the Core API is exposed (CLI, REST API, etc...)
*/
type Command struct {
	Options   []Option
	Arguments []Argument
	HelpText  HelpText
	PreRun    func(req Request) error
	Run       Func
	PostRun   Func
}
