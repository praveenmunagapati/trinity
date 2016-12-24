from .build_elements import *
from .build_cli import *

class Builder(object):
    def __init__(self, packages=[], command="", system_root=""):
        self.packages = sorted(packages, key=lambda x: x.precedence)
        self.system_root = system_root
        for package in self.packages:
            package.system_root = self.system_root
        if not command:
            command = "all"
        self.command = command

    def basic_build(self):
        for package in self.packages:
            if package.type not in complex_types:
                package.build()

    def build_all(self):
        for package in self.packages:
            package.build()

    def clean_all(self):
        for package in self.packages:
            if type(package) is Script:
                continue
            package.clean()

    def install_all(self):
        for package in self.packages:
            if type(package) is Script:
                package.build()
                continue
            package.install()
