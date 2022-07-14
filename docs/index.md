# Built@Home

An opensource and collaborative approach to building homelab's. The idea is to acknowledge that we can expose homelab implementation while increasing security and monitoring. The intention is to learn and cover more ground by working together.

> Nothing is done in a vacuum, we must all stand on our forefathers to better ourselves and the world around us

# Introduction

`Built@Home` is intended to be used by the community or welcomes contributors to build our own community of homelab developers. If you are interested in empowering your home with IOT and smarthome technology, networking, kubernetes, data storage, all layers of security and lab testing, privacy in data storage and streaming, and empowering your home with applications, wifi performance, and telemetry and monitoring of all layers of homelab tech. Then this project is right for you.


## Architecture

**Repositories**

1. HomeOps
2. HomeWare
3. HomeAssistant
4. HomeController
5. HomeOps-Webhook
6. Crossplane Proxmox Provider

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

### Deployment

`built@home` uses gitops principles to deploying changes in the homelab. Below, are the components that orchestrate the deployment of homelab resources.

1. **Kubernetes** - [Flux]() is used with [Kustomize]
() to perform continuous integration and deployment of kubernetes workloads. With kustomize, we can seperate workloads by environments by using a layering approach in our templates to define a base layer, and then additional layers that contain differences.

**Example**:
```sh
clusters/
  base/ <---similarities between dev and prod
  dev/ <--- additional changes on top of base
  prod/ <--- additional changes on top of base
```

When testing new resources, there is a cicd process to control deploying resources that are not using kustomize and a cicd process to clean up any resources that were deployed when migrating over to kustomize.

**Example**:

If we test deploying a new helm chart, once we configure the settings, we are ready to move it over to kustomize. There is an [action]() that will make sure to remove the helm chart and other related resources before flux deploys the new release.

2. **Ansible**

Primarily used to automate the setup of hardware. The list below contains the playbooks available.

  *  Setup of OS and settings on raspberryPi
  *  Setup / Teardown / Reset of K3s Cluster
  *  Setup of Flux, GPG, and SOPS on K3s Master

### Continious Integration

[Github Actions]() is used for all automation of resource management, testing, integrations, and some deployment processes. The kubernetes cluster has a [github actions controller]() which can optionally run all the CICD processes. Below is a list of workflows used and available in the homelab.


  * **HomeOps**
    - Upgrade Flux
    - [Renovate]()
    - Linting
  * **HomeWare**
    - Linting playbooks
  * HomeAssistant