        ·············································································
        :  ____ _____  ____  ____    _       ____ _____ _____ ____ _   _    _ ____  :
        : / _  |____ |/ _  |/ _  |  / \     / _  |____ |_   _|___ | | | |  | |___ \ :
        :| (_| | |_  | | | | | | | / _ \   | (_| | |_  | | | / ___| | | |  | |   | |:
        : > _  |___| | |_| | |_| |/ ___ \   > _  |___| | | || (___| |_| ___| |___| |:
        :/_/ |_|_____|\____|\____/_/   \_\ /_/ |_|_____| |_| \____|\___|_____|____/ :
        ·············································································
Helper script for adding new cluster config to the .kube/config that already existed in that directory.
The script is written in go lang.
Currently, the config file of .kube/config has to have at least one cluster to use this script.
By running this script, you will be asked to do two things.
The first one is the username of your system, which is for finding the .kube directory.
The second one is the directory where your intended clusters' configs exist.
The binary file that exists in this repo is for Linux OS amd64 ARCH.
If you want to use another system, you can request it from me directly.
Or you can just build the binary with this " env GOOS=YOUROS GOARCH=YOURARCH go build -o YOURBINNAME/DIR " command.
