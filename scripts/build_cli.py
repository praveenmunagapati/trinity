subcommands = [
"clean",
"install",
"all",
"build",
"test",
"format",
]

complex_types = [
"proto"
]

def handle_cli(args):
    if len(args) not in [2, 1]:
        print "[ {} ]  has incorrect number of arguments".format(" ".join(sys.argv))
        sys.exit(0)
    elif len(args) == 1:
        command = "build"
    else:
        command = args[1]
    if command not in subcommands:
        print "[ \"{}\" ]  is not a recognized command".format(sys.argv[1])
        sys.exit(0)
    return command
