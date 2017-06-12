GIT_HASH=`git rev-parse HEAD`

echo "package system\n\nvar CommitHead = \"$GIT_HASH\"\n\n" > core/system/version.go
