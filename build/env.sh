#!/bin/sh


# taken from https://github.com/ethereum/go-ethereum/blob/master/build/env.sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
tridir="$workspace/src/github.com/clownpriest"
if [ ! -L "$tridir/trinity" ]; then
    mkdir -p "$tridir"
    cd "$tridir"
    ln -s ../../../../../. trinity
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$tridir/trinity"
PWD="$tridir/trinity"

TRINITY_HOME="$HOME/.trinity"
export TRINITY_HOME

# Launch the arguments with the configured environment.
exec "$@"
