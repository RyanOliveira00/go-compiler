# Compilador TypeScript-Inspired

Um compilador simples para uma linguagem de programação inspirada em Typescript, implementado em Go.

## Índice

- [Sobre a Linguagem](#sobre-a-linguagem)
- [Características](#características)
- [Sintaxe](#sintaxe)
- [Como Executar](#como-executar)
- [Exemplos](#exemplos)
- [Implementação](#implementação)

## Sobre a Linguagem

Este projeto implementa um compilador para uma linguagem de programação estaticamente tipada com sintaxe inspirada em Go. A linguagem suporta tipos básicos, estruturas de controle de fluxo e operações de entrada/saída.

## Características

### Tipos de Dados

- `int`: Números inteiros
- `float`: Números de ponto flutuante
- `string`: Textos
- `bool`: Valores booleanos

### Declarações

- Variáveis podem ser declaradas com tipo explícito ou inferido
- Constantes são suportadas
- Escopo de bloco

### Operadores

- Aritméticos: `+`, `-`, `*`, `/`, `%`
- Comparação: `==`, `!=`, `<`, `<=`, `>`, `>=`
- Lógicos: `&&`, `||`, `!`
- Atribuição: `=`, `+=`, `-=`

### Estruturas de Controle

- `if/else` para condicionais
- `while` para loops
- Blocos delimitados por chaves

### Entrada/Saída

- `print()` para saída
- `read()` para entrada

## Sintaxe

### Declarações de Variáveis

```go
let x: int;              // Declaração com tipo
let y = 42;              // Tipo inferido
const pi = 3.14;         // Constante
let nome: string;        // String
```

### Estruturas de Controle

```go
if (x > 0) {
    print(x);
} else {
    print("negativo");
};

while (i < 10) {
    print(i);
    i = i + 1;
};
```

### Entrada e Saída

```go
print("Digite seu nome:");
read(nome);
print("Olá " + nome);
```

## Como Executar

1. Requisitos:

   - Go 1.16 ou superior
   - Git

2. Clone o repositório:

```bash
git clone https://github.com/RyanOliveira00/go-compiler.git
cd go-compiler
```

3. Execute o compilador:

```bash
go run src/main.go
```

4. Use o REPL interativo ou compile arquivos:

```bash
# REPL
>> let x = 42;
>> print(x);
42
```

## Exemplos

### Exemplo 1: Calculadora Simples

```go
let a = 10;
let b = 5;
print(a + b);    // 15
print(a * b);    // 50
print(a / b);    // 2
```

### Exemplo 2: Loop com Condicionais

```go
let i = 0;
while (i < 5) {
    if (i % 2 == 0) {
        print("par");
    } else {
        print("ímpar");
    };
    i = i + 1;
};
```

### Exemplo 3: Entrada de Dados

```go
let idade: int;
print("Digite sua idade:");
read(idade);
if (idade >= 18) {
    print("Maior de idade");
} else {
    print("Menor de idade");
};
```

## Implementação

### Estrutura do Projeto

```
src/
├── ast/            # Árvore sintática abstrata
├── lexer/          # Análise léxica
├── parser/         # Análise sintática
├── compiler/       # Geração de código
└── main.go         # Ponto de entrada
```

### Pipeline de Compilação

1. **Lexer**: Tokenização do código fonte
2. **Parser**: Geração da AST
3. **Compiler**: Execução/Interpretação do código

### Decisões de Design

1. **Sistema de Tipos**

   - Tipagem estática com inferência
   - Tipos básicos inspirados em Go
   - Verificação de tipos em tempo de compilação

2. **Sintaxe**

   - Inspirada em Go para familiaridade
   - Ponto e vírgula obrigatório
   - Blocos delimitados por chaves

3. **Execução**

   - Interpretador em vez de geração de código nativo
   - REPL para facilitar testes e aprendizado
   - Sistema de ambiente para variáveis

4. **Tratamento de Erros**
   - Mensagens de erro detalhadas
   - Validação de tipos
   - Verificação de variáveis não declaradas

### Limitações Atuais

- Sem suporte a funções
- Sem arrays ou estruturas de dados complexas
- Sem garbage collection
- Operações limitadas com strings

### Possíveis Extensões Futuras

- Implementar funções e procedimentos
- Adicionar arrays e slices
- Suportar estruturas (structs)
- Adicionar mais operadores e tipos de dados
