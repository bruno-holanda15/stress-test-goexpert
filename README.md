# Goexpert Stress Test

## Objetivo 
Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

## Entrada de Parâmetros via CLI

- `--url -u`: URL do serviço a ser testado.
- `--requests -r`: Número total de requests.
- `--concurrency -c`: Número de chamadas simultâneas.


## Execução do Teste

Realizar requests HTTP para a URL especificada.
Distribuir os requests de acordo com o nível de concorrência definido.
Garantir que o número total de requests seja cumprido.

### Geração de Relatório
- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).
- Exemplo de relatório

![relatório](./img/Screenshot%20from%202024-07-23%2020-05-49.png)

### Execução da aplicação

Criar imagem a partir do arquivo Dockerfile na raiz do projeto.
- `docker build -t stress-test .`

Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:
 - `docker run stress-test —url http://google.com —requests 1000 —concurrency 10`

Caso não seja passado as flags o sistema executará para argumentos default.
