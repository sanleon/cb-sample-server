# Cloud Build Sample Source

This is sample for Full Automation CI/CD with the Cloud Build of GCP.
You need to use the [Cloud Build](https://cloud.google.com/cloud-build/?hl=ja)

![chart](./cloud-build-auto-eng.png)

1. Create GKE Cluster for each a environment on GCP
	- develop cluster-name : dev-cb-sample-server
	- staging cluster-name : stg-cb-sample-server
	- product cluster-name : prd-cb-sample-server
2. Fork or Clone devops repository the [cb-sample-server-devops](https://github.com/sanleon/cb-sample-server-devops)
3. You need to setup [kms](https://cloud.google.com/kms/) for cb-sample-server-devops
4. You need to rewrite that 'id: Decrypt ssh key' and 'id: Set rsa file' part of [cloudbuild.yaml](./cloudbuild.yaml), [dev.cloudbuild.yaml](./dev.cloudbuild.yaml)
5. You need to set your devops repository url[3] of github on [cloudbuild.yaml](./cloudbuild.yaml), [dev.cloudbuild.yaml](./dev.cloudbuild.yaml)
6. Create Trigger of Docker image build for Push to branch(develop, staging, master) of source repository on Cloud Build.
	- Cloud Build trigger setup for develop : [dev.cloudbuild.yaml](./dev.cloudbuild.yaml)
	- Cloud Build trigger setup for staging, product : [cloudbuild.yaml](./cloudbuild.yaml)


## Local Build

### Prepare

- Install golang
- Install Docker

### Run
- Init(Need to run first this) : make init
- Build : make build
- Run App on Local Docker : make docker-build

 