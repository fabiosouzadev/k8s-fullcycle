<div align="center">
# K8s FullCycle

Este repositório contém exemplos para apaprendizado `Kubernetes` do curso FullCycle.
</div>

## Navegação no repositório
- [Pré-requisitos](#pré-requisitos)
- [Probes](#probes)
- [Resources](#resources)

## 💻Pré-requisitos
- [Docker](https://www.docker.com/)
- [DockerHub](https://hub.docker.com/)
- [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)

## Trabalhando com a imagem

### 🚀Build serviço go
```bash
docker build -t k8s-fullcycle .
```

### ☕Rodar servico go 
```bash
docker run --rm -p 80:80 k8s-fullcycle
```
### 🚀Subir imagem para DockerHub
```bash
docker push <seu-user-no-dockerhub>/k8s-fullcycle
```

## Rodando o Kind

### Criando o cluster
```bash
kind create cluster --config=k8s/kind.yaml --name=fullcycle
kubectl cluster-info --context kind-fullcycle
```
### Aplicando o arquivo de deployment
```bash
kubectl apply -f k8s/deployment.yaml
```
### Aplicando o service
```bash
kubectl apply -f k8s/service.yaml
```
...
### Aplicando o metrics-server (Com patch pra funcionar no kind)
```bash
kubectl apply -f k8s/metrics-server.yaml
```
...
## Probes

## Resources
