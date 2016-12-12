import os


def clean_gobin(name):
    gopath = os.path.expandvars("$GOPATH")
    binpath = os.path.join(gopath, "bin", name)
    os.remove(binpath)
