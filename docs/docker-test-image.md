# Docker Test Image

The `testImage` attribute is the name of the [Docker](https://docs.docker.com/get-started) image used to run the container responsible for testing the student's code.

The container runs with the following settings (options of `docker run`) :

- `--read-only`
  - Mount the container's root filesystem as read only
- `--network none`
  - Connect a container to a network without Internet
- `--memory 500M`
  - Memory limit
- `--cpus 2.0`
  - Number of CPUs
- `--user 1000:1000`
  - Username or UID (format: <name|uid>[:<group|gid>])
- `--env EXERCISE=hello-world`
  - Exercise name
- `--env USERNAME=aeinstein`
  - Student's login
- `--env HOME=/jail`
  - Home directory of the container
- `--env TMPDIR=/jail`
  - Temporary directory of the container
- `--workdir /jail`
  - Working directory inside the container
- `--tmpfs /jail:size=100M,noatime,exec,nodev,nosuid,uid=1000,gid=1000,nr_inodes=5k,mode=1700`
  - Mount a tmpfs directory on `/jail`, 100 MB writable.
- `--volume volume_containing_student_repository:/jail/student:ro`
  - Bind mount a volume containing the student repository, read-only.

Example of a [Dockerfile](https://github.com/01-edu/public/blob/master/js/tests/Dockerfile) and its [entrypoint](https://github.com/01-edu/public/blob/master/js/tests/entrypoint.sh).
