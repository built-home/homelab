# Built@Home

An opensource and collaborative approach to building homelab's. The idea is to acknowledge that we can expose homelab implementation while increasing security and monitoring. The intention is to learn and cover more ground by working together.

> Nothing is done in a vacuum, we must all stand on our forefathers to better ourselves and the world around us

# Introduction

`Built@Home` is intended to be used by the community or welcomes contributors to build our own community of homelab developers. If you are interested in empowering your home with IOT and smarthome technology, networking, kubernetes, data storage, all layers of security and lab testing, privacy in data storage and streaming, and empowering your home with applications, wifi performance, and telemetry and monitoring of all layers of homelab tech. Then this project is right for you.

## Architecture

## Repositories

### HomeOps

Continuous deployment of kubernetes workloads. All kubernetes definitions are in this repository using Kustomize. Flux is responsible for for deploying kubernetes resources, images, and updating cluster resources. The provisioning for the cluster hardware will be available in [HomeWare](), but provisioning for setting up `Flux`, `Sops`, and other tools are available in this repository.

### HomeWare

[Ansible](), [Terraform](), and [Packer]() are used to provision different infrastructure for `Built@Home`. This is the first repository that will be used to automate the setup of the homelab.

**Ansible**
1. `K3s` - provisions a k3s cluster on raspberry pis. (Refer to hardware for specifics)
2. `Unifi` - provision dream machine pro, access points, with unifi tools.
3. `Proxmox` - generally, proxmox is mostly controlled by terraform. Although, terraform will call specific playbooks for provisioning after a `VM` or `LXC` container for controlling additional tasks.

**Packer**
Packer is used for creating actual images that can be used for creating virtual machines.

**Terraform**
Terraform uses the [Proxmox Provider]() with [CloudInit]() to launch images created by packer.

### HomeAssistant
### HomeController
### HomeOps Webhook
### Crossplane Proxmox Provider

