import os
import shutil
from distutils.dir_util import *

static_dir = "static"

home_dir = os.path.expandvars("$HOME")
system_root = os.path.join(home_dir, ".trinity")


def install_static_root(original_static):
    final_static = os.path.join(system_root, static_dir)
    if os.path.isdir(final_static):
        shutil.rmtree(final_static)
    shutil.copytree(original_static, final_static)
