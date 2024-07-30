PicPay Simplificado

O PicPay Simplificado é uma plataforma de pagamentos simplificada. Nela é possível depositar e realizar transferências
de dinheiro entre usuários. Temos 2 tipos de usuários, os comuns e lojistas, ambos têm carteira com dinheiro e realizam
transferências entre eles.

### Requisitos

-----> Nao focar em volume, implementar funcionalidades antes de tudo
 1. A feature de autenticação é opcional, logo não deve ser o foco aqui.
 2. Testes são importantíssimos mas não podem atrasar o desenvolvimento dos demais pontos



A seguir estão algumas regras de negócio que são importantes para o funcionamento do PicPay Simplificado:

- [X] Para ambos tipos de usuário, precisamos do `Nome Completo`, `CPF`, `e-mail` e `Senha`. CPF/CNPJ e e-mails devem ser
  únicos no sistema. Sendo assim, seu sistema deve permitir apenas um cadastro com o mesmo CPF ou endereço de e-mail;

- [] Usuários podem enviar dinheiro (efetuar transferência) para lojistas e entre usuários;

- [] Lojistas **só recebem** transferências, não enviam dinheiro para ninguém;

- [] Validar se o usuário tem saldo antes da transferência;

- [] Antes de finalizar a transferência, deve-se consultar um serviço autorizador externo, use este mock para
  simular (https://run.mocky.io/v3/5794d450-d2e2-4412-8131-73d0293ac1cc);

- [] A operação de transferência deve ser uma transação (ou seja, revertida em qualquer caso de inconsistência) e o
  dinheiro deve voltar para a carteira do usuário que envia;

- [] No recebimento de pagamento, o usuário ou lojista precisa receber notificação (envio de email, sms) enviada por um serviço de terceiro e eventualmente este serviço pode estar indisponível/instável. Use este mock para simular o
  envio (https://run.mocky.io/v3/54dc2cf1-3add-45b5-b5a9-6bf7e7f1f4a6);

- [] Este serviço deve ser RESTFul.

> Tente ser o mais aderente possível ao que foi pedido, mas não se preocupe se não conseguir atender a todos os
> requisitos. Durante a entrevista vamos conversar sobre o que você conseguiu fazer e o que não conseguiu.

### Endpoint de transferência

Você pode implementar o que achar conveniente, porém vamos nos atentar **somente** ao fluxo de transferência entre dois
usuários. A implementação deve seguir o contrato abaixo.

```http request
POST /transfer
Content-Type: application/json

{
  "value": 100.0,
  "payer": 4,
  "payee": 15
}
```

Caso ache interessante, faça uma **proposta** de endpoint e apresente para os entrevistadores 