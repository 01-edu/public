# tests

## Docker image

The tests must be in a Docker image that will be executed with (among others) the following options :

- `--read-only` Mount the container's root filesystem as read only
- `--user 1000:1000` Avoid to give root rights
- `--tmpfs /jail:size=200M,noatime,exec,nodev,nosuid,uid=1000,gid=1000,nr_inodes=5k,mode=1700` Mount a 200MB tmpfs directory to create files and run tests
- `--memory 500M` Memory limit of 500 MB
- `--cpus 2.0` Number of CPUs (2 threads)
- `--mount readonly,type=volume,source=student/username,destination=/app/student` The student's code is available in /app/student
- `--env HOME=/jail` Set the only writable folder as home directory
- `--env TMPDIR=/jail` Set the only writable folder as temporary directory
- `--env USERNAME=` Username
- `--env EXERCISE=` Exercise name
- `--env DOMAIN=` Domain name (e.g. gp.ynov-bordeaux.com)
- `--env EXPECTED_FILES=` A space-separated list of required files

The username is the Gitea login.
No command or arguments are used, the entrypoint has to run the tests.
The exit status of the container will determine whether or not the test has passed.
Any output will be printed in the platform but not interpreted.
