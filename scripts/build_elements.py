import os
import subprocess as sp

import colors
from build_commands import *

class Script(object):
    def __init__(self, name="", script_path="", lang="",
                 root_func=None, args=[],precedence=0):
        self.name = name
        self.script_path = script_path
        self.script_name = os.path.basename(self.script_path)
        self.working_dir = os.path.dirname(self.script_path)
        self.lang = lang
        self.root_func = root_func
        self.args=args
        self.precedence = precedence
        self.type = "script"

    def build(self):
        if self.lang == "python":
            self.run_py_script()

    def run_py_script(self):
        print "{}  {}".format(colors.running_msg, self.name)
        self.root_func(*self.args)



class Package(object):
    def __init__(self, name="", root="", type=type, lang="",
                       depend_dir="", file_path="", output="",
                       system_root="", name_override=False,
                       clean_func=None, precedence=0):
        self.name = name
        self.root_path = os.path.abspath(root)
        self.system_root = system_root
        self.type = type
        self.language = lang
        self.depend_dir = os.path.abspath(depend_dir)
        self.file_path = os.path.abspath(file_path)
        self.output = os.path.abspath(output)
        self.precedence = precedence
        self.name_override = name_override
        self.clean_func = clean_func

    def build(self):

        print "{}  {}".format(colors.building_msg, self.name)
        if self.language == "go":
            if self.type == "lib":
                if self.name_override:
                    go_command(command="install", wd=self.root_path, name=self.name)
                else:
                    go_command(command="install", wd=self.root_path)
            elif self.type == "bin":
                if self.name_override:
                    go_command(command="build", wd=self.root_path, name=self.name)
                else:
                    go_command(command="build", wd=self.root_path)
            elif self.type == "plugin":
                go_command(command="build -buildmode=plugin", wd=self.root_path,
                            system_root=self.system_root, plugin=True)

        if self.language == "proto":
                protoc_go(self.depend_dir, self.file_path, self.root_path)
                protoc_py(self.depend_dir, self.file_path, self.output)
                protoc_cpp(self.depend_dir, self.file_path, self.output)


    def install(self):
        print "{}  {}".format(colors.install_msg, self.name)
        if self.language == "go":
            if self.type == "plugin":
                go_command(command="build -buildmode=plugin", wd=self.root_path,
                            system_root=self.system_root, plugin=True)
            else:
                if self.name_override:
                    go_command(command="install", wd=self.root_path, name=self.name)
                else:
                    go_command(command="install", wd=self.root_path)

    def clean(self):
        print "{}  {}".format(colors.cleaning_msg, self.name)
        if self.language == "go":
            if self.clean_func:
                self.clean_func(self.name)
            else:
                go_command("clean -i", self.root_path)
