# Cache LRU

## Problema

Nesse desafio você precisa desenvolver uma estrutura de dados que se comporte como um Cache LRU que é um tipo de dado que guarda até X itens, sendo que quando já há X itens, o item com acesso mais antigo será removido.

### Requisitos

* Escreva uma estrutura/classe que representa o Cache LRU .
* O tipo definido possui três métodos/funções:
    * Função construtura; recebe a capacidade do cache, um número inteiro.
    * Método Get; recebe uma chave string e retorna um valor de qualquer tipo.
    * Método Set; recebe a chave e o valor.
* O método Get retorna -1 se não existir item para chave recebida.
* O método Get memoriza a ordem de acesso de cada chave.
* O método Set:
    * Salva um valor baseado numa chave se o número de items salvos for abaixo da capacidade.
    * Substitui o item com acesso mais antigo se a capacidade de itens foi atingida