# CEP e Temperatura API

Este programa em Go é um servidor HTTP que fornece informações de endereço com base em CEP e dados de temperatura associados a uma localidade específica. O servidor utiliza duas APIs externas: ViaCEP para obter dados de endereço a partir do CEP e WeatherAPI para obter informações de temperatura com base na localidade.

## Funcionalidades

- **Obtenção de Informações de Endereço por CEP:**
  - O servidor permite a recuperação de detalhes de endereço (logradouro, bairro, localidade e UF) fornecidos um CEP válido.
  - Se o CEP não for válido (não contiver exatamente 8 dígitos ou for igual a "00000000"), o servidor retorna um erro com um status HTTP 422.

- **Previsão de Temperatura por Localidade:**
  - O servidor também fornece informações de temperatura atual em Celsius, Fahrenheit e Kelvin com base na localidade informada.
  - Se a localidade estiver ausente ou não puder ser obtida, o servidor retornará uma resposta com temperatura zero.

## Como utilizar no ambiente de deploy:
https://golang-temperature-api-cloudrun-wirapf7gfq-uc.a.run.app

ou com a URL direta

https://golang-temperature-api-cloudrun-wirapf7gfq-uc.a.run.app/cep/01001000

## Como Utilizar localmente:

1. **Requisitos:** 
&nbsp;
   - Certifique-se de ter o Go instalado na sua máquina.
&nbsp;
2. **Clonar o Repositório:**
&nbsp;

```bash
git clone https://github.com/kleytonsolinho/golang-temperature-api-cloudrun.git
```
&nbsp;
3. **Acesse a pasta do app:**
&nbsp;

```bash
cd golang-temperature-api-cloudrun
```
&nbsp;
4. **Rode o app:**
&nbsp;

```bash 
go run cmd/server/main.go
```

&nbsp;
5. **Dispare um cURL para a URL:**
&nbsp;

```bash
curl http:localhost:8080/cep/01001000
````