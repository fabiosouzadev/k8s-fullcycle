<div align="center">

# K8s FullCycle

Este repositório contém exemplos para aprendizado `Kubernetes` do curso FullCycle.

</div>

## Navegação no repositório

- [Pré-requisitos](#pré-requisitos)
- [Rodando o Kind](#rodando-o-kind)
- [Pods](#pods)
- [Replicaset](#replicaset)
- [Deployment](#deployment)
- [Services](#services)
- [Variaveis de ambiente](#variaveis-de-ambiente)
    - [ConfigMap](#configmap)
    - [Secrets](#secrets)
- [Probes](#probes)
    - [LivenessProbe](#livenessProbe)
    - [ReadinessProbe](#readinessProbe)
    - [StartupProbe](#startupProbe)
- [Resources e HPA](#resources-e-hpa)
  - [Aplicando o matrics-server](#aplicando-o-metrics-server)
  - [Resources](#resources)
  - [HPA](#hpa)
    - [Aplicando Hpa](#aplicando-hpa)
- [🧪Stress Test](#stress-test)
    - [Fortio](#fortio)
- [Statefulset e volumes persistentes](#statefulset-e-volumes-persistentes)
    - [PersistentVolume](#persistentvolume)
    - [PersistentVolumeClaims](#persistentvolumeclaims)

## 💻Pré-requisitos

- [Docker](https://www.docker.com/)
- [DockerHub](https://hub.docker.com/)
- [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)

## Rodando o Kind
```bash
### Criando o cluster
kind create cluster --config=k8s/kind.yaml --name=[nome-do-cluster]
kubectl cluster-info --context kind-[nome-do-cluster]

## Trabalhando com a imagem
### 🚀Build go server
docker build -t k8s-fullcycle:v[version]  k8s-fullcycle:latest .
### ☕Rodar servico go
# docker run --rm -p 80:80 k8s-fullcycle:v[version]

### 🚀Subir imagem para DockerHub
docker push <seu-user-no-dockerhub>/k8s-fullcycle

### Aplicando o service
kubectl apply -f k8s/service.yaml

### Aplicando configmaps
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/configmap-family.yaml

### Aplicando secret
kubectl apply -f k8s/secret.yaml

### Aplicando metrics-server 
kubectl apply -f k8s/metrics-server.yaml

### Aplicando hpa
kubectl apply -f k8s/hpa.yaml

### Aplicando persistent-volume-claim
kubectl apply -f k8s/persistent-volume-claim.yaml

### Aplicando o arquivo de deployment
kubectl apply -f k8s/deployment.yaml

```


## Deployment

## Services

## Variaveis de ambiente

### ConfigMap

### Secrets

## Probes

### LivenessProbe
```yaml
### LivenessProbe --> conteiner saudavel ####
....
spec
  containers:
...
    livenessProbe:
      httpGet:
        path: /healthcheck
        port: 5000 -> porta do container
      periodSeconds: 5    -> teste de X em X segundos
      failureThreshold: 1 -> tentivas necessarias pra o k8s reiniciar o pod
      timeoutSeconds: 1
      successThreshold: 1
      initialDelaySeconds: 10 -> tempo de espera pra começar a verificação
....
      (Caso a requisição tenha problemas o container é reiniciado)
EXEC
$ kubectl delete deployments.app goserver
$ kubectl apply -f k8s/deployment.yaml && kubectl get pods -w
```
### ReadinessProbe
```yaml
### ReadinessProbe --> conteiner pronto pra receber conexões ####
....
spec
  containers:
...
    readinessProbe:
      httpGet:
        path: /healthcheck
        port: 5000 -> porta do container
      periodSeconds: 5    -> teste de X em X segundos
      failureThreshold: 1 -> tentivas necessarias pra o k8s reiniciar o pod
      initialDelaySeconds: 10 -> tempo de espera pra começar a verificação
....
      (Caso a requisição tenha problemas o container não fica disponivel para receber requisições - NON Ready)
EXEC
$ kubectl delete deployments.app goserver
$ kubectl apply -f k8s/deployment.yaml && kubectl get pods -w
```
### StartupProbe
```yaml
### StartupProbe --> roda apenas 1 vez na inicialização do container ####
....
spec
  containers:
...
    startupProbe:
      httpGet:
        path: /healthcheck
        port: 5000  -> porta do container
      periodSeconds: 3    -> teste de X em X segundos
      failureThreshold: 30 -> tentativas de verificacao do healthcheck
      initialDelaySeconds: 10 -> tempo de espera pra começar a verificação
....
      (Nesse caso ocorrera a verificacao em até 90 segundos e se estiver ok - readinessProbe e livenessProbe começam a funcionar)
```
```bash
 kubectl delete deployments.app goserver
 kubectl apply -f k8s/deployment.yaml && kubectl get pods -w
```
## Resources e HPA
### Aplicando o metrics-server

> (Com patch pra funcionar no kind)

```bash
kubectl apply -f k8s/metrics-server.yaml
```

### Resources

```yaml
...
spec:
  containers:
    ...
    resources:
      requests:
        cpu: 100m # ou "0.1" => 0.1 vCPU
        memory: 128Mi
      limits:
        cpu: 250m
        memory: 256Mi
    ...
```

`requests` é referente ao minimo de recursos provisionados para o container.
`limits` se refere a quantidade máxima de recursos que um container deve utilizar.

Para cpu a unidade de medida é o

> vCPU = 1000m (milicores)
> 1/2 vCPU = 500m ou 0.5 vCPU

Para memory a unidade é Mi = Mb

> 20Mi = 20Mb

Caso não existam rescursos suficientes, o pod ficará em `PENDING` até o que o cluster tenha recursos disponivel para provisionar.

### HPA

> Horizontal Pod Autoscaling

```yaml
# k8s/hpa.yaml
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  name: goserver-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: goserver
  minReplicas: 1
  maxReplicas: 30
  targetCPUUtilizationPercentage: 25
```

### Aplicando hpa
```yaml
# k8s/hpa.yaml
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  name: goserver-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: goserver
  minReplicas: 1
  maxReplicas: 5
  targetCPUUtilizationPercentage: 25
```

```bash
kubectl apply -f k8s/hpa.yaml
```

## 🧪Stress Test
> Stress Test com [fortio](https://github.com/fortio/fortio)
> Test para testar hpa

### Fortio
```bash
kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 200s -c 50 "http://[nome-do-service]:[porta-do-service]/healthcheck"
```
OU
```yaml
#kubectl apply -f k8s/fortio.yaml
apiVersion: v1
kind: Pod
metadata:
  name:  traffic-generator
  labels:
    app: traffic-generator
spec:
  containers:
  - name: fortio
    image: fortio/fortio
    args: ["load", "-t", "0", "-qps", "1800", "-c", "50", goserver-service:8080/healthcheck"]

```
```bash
kubectl apply -f k8s/fortio.yaml
```
## Statefulset e volumes persistentes
### PersistentVolume
>A PersistentVolume (PV) is a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using Storage Classes. It is a resource in the cluster just like a node is a cluster resource. PVs are volume plugins like Volumes, but have a lifecycle independent of any individual Pod that uses the PV. This API object captures the details of the implementation of the storage, be that NFS, iSCSI, or a cloud-provider-specific storage system.

Na tradução livre:

> Um PersistentVolume (PV) é uma parte do armazenamento no cluster que foi provisionado por um administrador ou provisionado dinamicamente usando classes de armazenamento. É um recurso no cluster, assim como um nó é um recurso de cluster. PVs são plug-ins de volume como Volumes, mas têm um ciclo de vida independente de qualquer Pod individual que usa o PV. Esse objeto API captura os detalhes da implementação do armazenamento, seja NFS, iSCSI ou um sistema de armazenamento específico do provedor de nuvem.

```yaml
#https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistent-volumes
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv0003
spec:
  capacity:
    storage: 5Gi
  volumeMode: Filesystem
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Recycle
  storageClassName: slow
  mountOptions:
    - hard
    - nfsvers=4.1
  nfs:
    path: /tmp
    server: 172.17.0.2
```
### PersistentVolumeClaims
```yaml
# k8s/persistent-volume-claim.yaml

# https://kubernetes.io/docs/concepts/storage/persistent-volumes/
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: goserver-pvc
spec:
  # AKS: default,managed-premium
  # GKE: standard
  # EKS: gp2 (custom)
  # Rook: rook-ceph-block,rook-ceph-fs
  # storageClassName: default
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
```
