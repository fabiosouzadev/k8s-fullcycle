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
```yml
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
EXEC
$ kubectl delete deployments.app goserver
$ kubectl apply -f k8s/deployment.yaml && kubectl get pods -w
```
```yml
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
```

```yml
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
## Resources
