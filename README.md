# Stay active
You can put a spoon on your laptop's trackpad... or use this script to keep you active.


## Usage
This script only push the caps lock key once per minute, so your laptop, Teams or Slack will not shut down or go afk.


### Download
Just download your OS executable in [releases](https://github.com/R-dVL/stay-active/releases).


## Security issues
This executable is not signed, your device may not trust the origin. You can compile your own binaries with [this repo Dockerfile](https://github.com/R-dVL/stay-active/blob/main/Dockerfile), there are [docker compose](https://github.com/R-dVL/stay-active/blob/main/docker-compose.yml) examples to compile Windows and Linux amd64 binaries.

### Compilation
Just git clone or download this repository and execute:

```bash
docker compose -f .\docker-compose.[YOUR OS].amd64.yml up
```

It will build the Docker image and compile binaries that will be located in the bin folder created.

