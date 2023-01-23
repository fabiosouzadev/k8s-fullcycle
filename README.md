# K8s FullCycle

Este repositório contém exemplos para uso do `Kubernetes` do curso FullCycle.

## 💻 Pré-requisitos
- [Docker](https://www.docker.com/)

## Uso

### 🚀 Build serviço go
```bash
docker build -t hello-go .
```

### ☕ Rodar servico go 
```bash
docker run --rm -p 80:80 hello-go
```

