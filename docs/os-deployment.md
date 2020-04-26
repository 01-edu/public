# OS Deployment

## Image creation steps

- [Installation of Ubuntu](ubuntu-installation.md)
  - minimal OS installation (downloads ~200 MB)
  - Software installation (downloads ~900 MB)
  - Optimization
    - improve speed
    - reduce image size
    - reduce power (CPU) & memory usage
    - reduce surface of attack
    - reduce bandwidth usage
  - Customization
    - machine-dependent (drivers, bug workarounds...)
    - time zone of the school
    - school scripts
  - Cleaning
    - logs
    - temporary files
    - histories
    - caches
    - auto-generated IDs
- Preparation of the disk image
  - zero unallocated space of filesystem (~7 GB of data remains)
  - (optional) create compressed image with [lz4](https://lz4.github.io/lz4) (the resulting image is ~3.2 GB)

## Network installation

- Boot through PXE [UDPcast](http://udpcast.linux.lu) which allows an efficient transfer of the disk image (using multicast or broadcast)
