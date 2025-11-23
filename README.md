# lawyers-api

- Esse é um service para testar seus conhecimentos no curso observability in deep

## Oque ele faz ?

- é uma api que faz consultas em um banco de dados PG no estilo CRUD
- É instrumentada com o Open telemetry para resgatar métricas e trancing

## Como usar ela ?

- git clone na sua maquina local:

```bash
git clone (urlrep)
```

- Após subir seu cluster basta realizar o deploy usando o comando kubectl

```bash
kubectl apply -f ./k8s/values.yaml
```

- PS:
    - Não se esqueça de configurar suas envs de conexão no banco no seu configmap

- **AlERTA**

Para seguir com os testes é necessario criar um banco no seu cluster ou criar na plataform que vamos indicar no cursos

você pode subir o banco pg em qualquer lugar que desejar caso a solução apresentada no curso não esteja disponivel.
