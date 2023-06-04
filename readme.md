# GRPC

GRPC é um framework RPC de alto desempenho e de código aberto, criado pelo Google com objetivo de facilitar a comunicação entre sistema. Ele pode ser usado em qualquer ambiente em que o HTTP é usado atualmente para comunicação entre clientes e servidores. No entanto, o GRPC é usado principalmente para criar APIs de serviço da Web. Ele permite que um cliente faça uma chamada de método diretamente em um servidor de serviço no mesmo ou em um sistema diferente, como se fosse um objeto local.

## Conceitos iniciais

- GRPC é um framework RPC de alto desempenho e de código aberto, criado pelo Google com objetivo de facilitar a comunicação entre sistema.
- Extremamente leve
- É mantido pela CNCF (Cloud Native Computing Foundation)

### Em quais casos posso utilizar

- Ideal para microsserviços
- Mobile, Browsers e Backend
- Geração das bibliotecas de forma automática
- Streaming Bidirecional utilizando HTTP/2

### Linguagens Suportadas

- GO
- Java
- C
- C++
- Node.js
- C#
- Python
- PHP
- Ruby
- Dart

## GRPC Http/2 e Protocol Buffers

### RPC - Remote Procedure Call

O RPC funciona da seguinte forma:  o cliente chama um método no servidor remoto, o servidor recebe a solicitação, executa o procedimento e envia a resposta de volta ao cliente. O cliente não precisa saber como o servidor executa o procedimento, ele só precisa saber a interface do procedimento.

### Protocol Buffers

Protocol Buffers é uma linguagem de descrição de interface neutra de plataforma para serialização de dados estruturados. É uma alternativa ao XML, JSON e outros formatos de troca de dados. Um arquivo .proto é usado para definir a estrutura dos dados que são serializados/deserializados. O arquivo .proto é compilado em uma classe que implementa a estrutura de dados e as rotinas de serialização/deserialização.

### Protocol Buffers vs JSON

- Protocol Buffers é 3 a 10 vezes menor que JSON
- Protocol Buffers é 20 a 100 vezes mais rápido que JSON
- Protocol Buffers é menos verboso que JSON
- Protocol Buffers é mais rápido para processar
- Protocol Buffers é mais fácil de escrever e manter
- Protocol Buffers é binário
- Protocol Buffers é mais fácil de usar, evoluir, debugar, testar,documentar e  ler

```proto
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}
```

### HTTP/2

- O nome oficial é HTTP/2, mas é comum vermos HTTP2 ou HTTP 2.0
- O nome original era HTTP Speed+Mobility
- Lançado em 2015
- Dados trafegados são binários
- Utiliza a mesma conexão TCP para enviar e receber dados (multiplexação)
- Server Push, permite que o servidor envie dados para o cliente antes mesmo dele solicitar
- Cabeçalhos são comprimidos, reduzindo o tamanho da requisição e resposta
- Permite o uso de criptografia TLS (HTTPS)
- Priorização de requisições, permite que o cliente informe ao servidor qual requisição é mais importante
- Permite o uso de streaming bidirecional, ou seja, o cliente e o servidor podem enviar dados simultaneamente
- Permite o uso de um canal de comunicação entre o cliente e o servidor, chamado de stream

## GRPC - API Unary

É uma forma de comunicação entre cliente e servidor, onde o cliente envia uma requisição para o servidor e recebe uma resposta. É o tipo de comunicação mais comum.

## GRPC - API Server Streaming

É uma forma de comunicação entre cliente e servidor, onde o cliente envia uma requisição para o servidor e recebe uma stream de respostas. O cliente precisa ficar escutando a stream de respostas até que o servidor envie todos os dados.

## GRPC - API Client Streaming

É uma forma de comunicação entre cliente e servidor, onde o cliente envia uma stream de requisições para o servidor e recebe uma resposta. O servidor precisa ficar escutando a stream de requisições até que o cliente envie todos os dados.

## GRPC - API Bidirectional Streaming

É uma forma de comunicação entre cliente e servidor, onde o cliente envia uma stream de requisições para o servidor e recebe uma stream de respostas. O cliente e o servidor precisam ficar escutando a stream de requisições e respostas até que o cliente envie todos os dados.

## GRPC - SSL/TLS

SSL (Secure Sockets Layer) é um protocolo de segurança que permite que os dados trafeguem de forma segura entre o cliente e o servidor. O SSL foi substituído pelo TLS (Transport Layer Security), mas o nome SSL ainda é utilizado.

## GRPC - Deadlines e Timeouts

Deadlines e timeouts são muito importantes para garantir que o cliente não fique esperando uma resposta do servidor para sempre. O deadline é o tempo máximo que o cliente está disposto a esperar por uma resposta do servidor. O timeout é o tempo máximo que o servidor está disposto a esperar para enviar uma resposta para o cliente.

## GRPC - Interceptadores

Interceptadores são funções que são executadas antes ou depois de uma chamada de método. Os interceptadores são muito úteis para fazer logs, autenticação, validação, etc.

## GRPC - Reflection

Reflection é uma API que permite que o cliente descubra os serviços que o servidor oferece. O Reflection é muito útil para fazer ferramentas de debug.

## GRPC vs REST

- REST é uma arquitetura
- GRPC é um framework RPC
- REST é baseado em recursos
- GRPC é baseado em serviços
- REST é stateless
- GRPC é stateful
- REST utiliza o protocolo HTTP/1.1
- GRPC utiliza o protocolo HTTP/2
- REST utiliza JSON
- GRPC utiliza Protocol Buffers
- REST utiliza os verbos HTTP
- GRPC utiliza os métodos RPC
- REST utiliza o padrão de projeto MVC
- GRPC utiliza o padrão de projeto Repository
- REST é mais simples
- GRPC é mais complexo
- REST é mais lento
- GRPC é mais rápido
- REST é mais popular
- GRPC é menos popular
- REST é mais maduro
- GRPC é menos maduro
- REST é mais flexível
- GRPC é menos flexível
- REST é mais fácil de aprender
- GRPC é mais difícil de aprender
