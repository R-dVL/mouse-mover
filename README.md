# Stay active
You can put a spoon on your laptop's trackpad... or use this script to keep you active.


## Contents
1. [Usage](#Usage)
2. [Download](#Download)
3. [Compilation](#Compilation)

## Usage
This script will keep the computer active to prevent it from sleeping or appearing as afk, you will also have an option to schedule a computer shutdown.


### Modes
There are different modes for different situations.


#### Teams
The caps lock is pressed every minute so that the computer remains active without interrupting you, so that you can have a coffee.


#### Gaming
The W, S, 1 and 2 keys are pressed randomly so your character will do random things and the server will not kick you.


#### Mouse mover
Mouse will move once a minute.


## Download
Just download the executable for your operating system at [releases](https://github.com/R-dVL/stay-active/releases).


## Compilation
You can compile your own binaries with [this Dockerfile](https://github.com/R-dVL/stay-active/blob/main/Dockerfile), there is a [docker compose](https://github.com/R-dVL/stay-active/blob/main/docker-compose.yml) example to compile Windows, Linux and Darwin amd64 binaries.

### Compilation
Just git clone or download this repository and run:

```bash
docker compose up
```

It will build Docker images and compile the binaries that will be placed in the created bin folder.

