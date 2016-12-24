import os
import subprocess as sp

def version_root(out_path):
    head_commit = get_head_commit_hash()
    write_version_file_go(head_commit, out_path)


def write_version_file_go(commit_hash, out_path):
    with open(out_path, "w") as output:
        output.write("package system\n"\
                     "\nvar CommitHead = \"{}\"\n".format(commit_hash))

def get_head_commit_hash():
    out =  sp.check_output(["git", "rev-parse", "HEAD"])
    return out.strip()
