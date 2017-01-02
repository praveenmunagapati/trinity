#!/usr/bin/env python3

import os
import sys
import shutil
import subprocess as sp

from scripts.colors import *
from scripts.build_cli import *
from scripts.build_elements import *
from scripts.builder import *
from scripts.clean_commands import *
from scripts.install_static import *
from scripts.version import *
from scripts.download_static import *

system_root = os.path.join(os.path.expanduser("~"), ".trinity")


if __name__ == "__main__":

    command = handle_cli(sys.argv)

    version_info   = Script(name="version_generate",
                            lang="python", root_func=version_root,
                            args=["core/system/version.go"],
                            precedence=0)

    download_static = Script(name="download_static",
                            root_func=download_static_root,
                            lang="python", args=["static"],
                            precedence=0)

    install_static = Script(name="install_static",
                            root_func=install_static_root,
                            lang="python", args=["static"],
                            precedence=1)

    core_proto_api = Package(name="core_proto_api", root="api/trinity",
                             type="proto", lang="proto", depend_dir="protos",
                             file_path="protos/trinity.proto", output="protos/out",
                             precedence=2)


    core_config    = Package(name="core_config", root="core/config",
                             type="lib", lang="go",precedence=3)

    core_system    = Package(name="core_system", root="core/system",
                             type="lib", lang="go", precedence=3)

    core_store     = Package(name="core_store", root="core/store",
                             type="lib", lang="go", precedence=3)

    core_token     = Package(name="core_token", root="core/token",
                             type="lib", lang="go", precedence=3)


    gateway        = Package(name="gateway", root="subsystems/trinity_gateway",
                             type="bin", lang="go", precedence=4)


    tfindex        = Package(name="findex", root="subsystems/trinity_findex",
                             type="bin", lang="go", precedence=4)

    thtml          = Package(name="html", root="subsystems/trinity_html",
                             type="bin", lang="go", precedence=4)

    tiindex        = Package(name="iindex", root="subsystems/trinity_iindex",
                             type="bin", lang="go", precedence=4)

    # trinity_subs   = Package(name="trinity_subsystem", root="trinity",
    #                          type="bin", lang="go", precedence=4,
    #                          clean_func=clean_gobin, name_override=True)

    trinity_core   = Package(name="trinity", root="trinity",
                             type="bin", lang="go", precedence=5)

    if command == "travis":
        scripts = [version_info]
    else:
        scripts = [version_info, download_static, install_static]

    proto_api = [core_proto_api]
    core_libs = [core_config, core_system, core_token, core_store]
    subsystems = [gateway, tfindex, tiindex, thtml]


    packages = [trinity_core]  + core_libs + subsystems + proto_api + scripts

    builder = Builder(packages, system_root=system_root)

    if command == "all":
        builder.build_all()
    elif command == "build":
        builder.basic_build()
    elif command == "clean":
        builder.clean_all()
    elif command == "install":
        builder.install_all()
    elif command == "travis":
        builder.install_all()
