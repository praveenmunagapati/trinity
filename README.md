# trinity

<p align="center">
<img src="http://i.imgur.com/3S2iUbl.png" width="300">
</p>

<img src="https://travis-ci.org/clownpriest/trinity.svg?branch=master">

## install

#### dependencies

A lot of the dependencies will be handled by the build.py script, but you need to get a few things yourself (either through ```apt-get``` or ```homebrew```, or ```go get``` or whatever):
- **go-ipfs**
- **zstd**
- **gogoprotobuf**  (only necessary if you're compiling the proto yourself)
 - ```go get github.com/gogo/protobuf/proto```
 - ```go get github.com/gogo/protobuf/protoc-gen-gogo```
 - ```go get github.com/gogo/protobuf/gogoproto```
- **requests** (python library)
 - `$: pip3 install -U requests`

##### compilers
- **Go** (at least 1.8)
- **Python** (2.7, for build.py and other scripts)
- **protoc** (the standard protobuf compiler)

After you have everything listed above, build/install with:

```
$: git clone https://github.com/clownpriest/trinity.git
$: cd trinity
$: ./build.py install
$: trinity init
```

## develop

If you want to work on trinity itself:


```bash
$: git clone https://github.com/clownpriest/trinity.git
$: cd trinity
$: ./build.py all
$: ./build.py install
```



### build.py

This is the build script that makes dealing with all the packages and binaries more convenient.
I don't feel like dealing with Makefiles, so this python script is going to be the main interface
for building/installing/deleting this system.

###### commands:

These are the commands you can pass to build.py:
```bash
$: ./build.py           # basic build (ignores some things, e.g. like compiling proto)
$: ./build.py all       # builds everything
$: ./build.py install   # basic build + install (does not compile proto)
$: ./build.py clean     # delete everything that was built/installed
```
