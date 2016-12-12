import os
import shutil
import subprocess as sp

def go_command(command, wd, plugin=False,
               system_root="", name=""):
    command_split = command.split()
    the_command = ["go"] + command_split
    if name:
        if command == "install":
            gopath = os.path.expandvars("$GOPATH")
            binpath = os.path.join(gopath, "bin", name)
            the_command = ["go", "build", "-o", binpath]
        else:
            the_command += ["-o", name]
    p = sp.Popen(the_command, cwd=wd)
    p.communicate()
    if plugin:
        move_plugin_to_system_root(system_root, wd)


def move_plugin_to_system_root(system_root, plugin_out_dir):
    targets = []
    for root, dirs, files in os.walk(plugin_out_dir):
        for file in files:
            if file.endswith(".so"):
                targets.append(os.path.join(root, file))
    if len(targets) == 1:
        target = targets[0]
        plugin_name = os.path.basename(target).replace(".so", "")
        subsystems_path = os.path.join(system_root, "subsystems")
        if not os.path.isdir(subsystems_path):
            os.makedirs(subsystems_path)
        shutil.copy(target, subsystems_path)
        os.remove(target)

def protoc_go(proto_dir, proto_file, wd):
    # $: protoc -I../protos --gofast_out=plugins=grpc:. ../protos/trinity.proto
    command = ["protoc", "-I"+proto_dir, "--gofast_out=plugins=grpc:"+wd, proto_file]

    p = sp.Popen(command, cwd=wd)
    p.communicate()


def protoc_py(proto_dir, proto_file, wd):
    # $: python -m grpc.tools.protoc -I../../protos --python_out=. --grpc_python_out=. ../../protos/route_guide.proto
    command = ["python", "-m", "grpc.tools.protoc",
               "-I"+proto_dir, "--python_out="+wd,
               "--grpc_python_out="+wd, proto_file]

    p = sp.Popen(command, cwd=wd)
    p.communicate()


def protoc_cpp(proto_dir, proto_file, wd):
    grpc_cpp_plugin = sp.check_output(["which", "grpc_cpp_plugin"])
    grpc_cpp_plugin = grpc_cpp_plugin.strip()

    command = ["protoc", "-I", proto_dir, "--grpc_out="+wd,
              "--plugin=protoc-gen-grpc={}".format(grpc_cpp_plugin),  proto_file]
    p = sp.Popen(command, cwd=wd)
    p.communicate()

    command = ["protoc", "-I"+proto_dir, "--cpp_out="+wd, proto_file]
    p = sp.Popen(command, cwd=wd)
    p.communicate()
