# Comparador de arquivos

Aplicação via linha de comando que permite comparar arquivos linha por linha.

## Modos de uso:
    1: Default:
        Criar uma pasta na raíz do projeto, com os arquivos ancestral.txt e atual.txt.
        Nesse modo, esses 2 arquivos sempre serão comparados.

    2: Diretório:
        Rodar a aplicação, passando os argumentos:
            [dir] [caminho_diretório]
        Nesse modo, todos os arquivos nesse diretório serão comparados.

    3: Múltiplos arquivos:
        Rodar a aplicação, passando os argumentos:
            [multiple] [caminhos_arquivos]...
        Nesse modo, todos os arquivos passados serão comparados.

## Modo de comparação:
    Quantidade de linhas e conteúdo de cada linha, excluindo os espaços iniciais e finais. 
    