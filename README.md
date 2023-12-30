# Stay active
You can put a spoon on your laptop's trackpad... or use this script to keep you active.


## Contents
1. [Usage](#Usage)
2. [Download](#Download)
3. [Security issues](#Security%20issues)

## Usage
This script will keep the computer active to prevent it from sleeping or appearing as afk, you will also have an option to schedule a computer shutdown.


### Modes
There are different modes for different situations.


#### Work
The caps lock is pressed every minute so that the computer remains active without interrupting you, so that you can have a coffee.


#### Gaming
The W, S, 1 and 2 keys are pressed randomly so your character will do random things and the server will not kick you.


#### Mouse mover
Mouse will move once a minute.


## Download
Just download the executable for your operating system at [releases](https://github.com/R-dVL/stay-active/releases).


## Security issues
This executable is not signed, your device may not trust the source. You can compile your own binaries with [this Dockerfile](https://github.com/R-dVL/stay-active/blob/main/Dockerfile), there are [docker compose](https://github.com/R-dVL/stay-active/blob/main/docker-compose.yml) examples to compile Windows and Linux amd64 binaries.

### Compilation
Just git clone or download this repository and run:

```bash
docker compose -f .\docker-compose.[YOUR OS].amd64.yml up
```

It will build the Docker image and compile the binaries that will be placed in the created bin folder.

