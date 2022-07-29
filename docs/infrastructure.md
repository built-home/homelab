
### Hardware

1. Unifi Dream Machine Pro
2. Two Unifi Wifi 6 Access points (1 upstairs | 1 downstairs)
3. Virtualization Server
  - Supermicro Motherboard M11....
  - 2 64GB Memory sticks
  - 1TB SSD nvme.2 disk drive
4. Raspberry Pi
  - 1 8gb raspberry pi 4
  - 1 4gb raspberry pi 4
  - 1 4gb raspberry pi 3
  - 3 PoE raspberryPi adapters
  - 1 PoE switch

### Infrastructure

1. **Kubernetes** - Running a 3 Node K3s cluster on raspberry pi with Ubuntu 20.04 OS.
2. **Proxmox** - Running virtual host machines and LXC Containers.
3. **Portainer** - Have not decided where to host portainer yet, but it is strickly a docker host running containers. This will probably run on an old 2011 4gb macbook pro. This will make sense when learning about the operations that the containers will perform.