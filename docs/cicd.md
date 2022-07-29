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