- Funções do launcher:
  - *Autenticação do usuário*
  - *Listar programas*
  - *Baixar e instalar os programas selecionados*
  - *Verificação de atualizações ao instalar*
  - *Verificação de updates regulares*
  - *Instala atualizações*
  - *Em caso de usuário inativo, inutiliza os apps*
  - *Pré Execução dos apps*
  - *Notificações dos apps*

---

- Funções detalhadas:
  - *Autenticação do usuário*:
    - *Remove qualquer env file anterior*
    - *Login com email e senha*
    - *Em caso de usuário inativo, inutiliza os apps já instalados*
    - *Após autenticar*:
      - *Gera um arquivo de variaveis de ambiente para o usuário*
      - *Grava o email do usuário no env file*
      - *Grava um token de identificação no env file*
      - *Grava um token criptografado do github no env file*
  - *Listar programas*:
    - *Listar os programas disponíveis para o usuário*
    - *Listar versões de cada programa*
  - *Baixar/instalar/desinstalar os programas selecionados*:
    - *Baixa o programa num diretório temporário*
    - *Após finalizar o download, move para diretório de instalação*
    - *No Caso de programas instalados, exibe a opçao de desinstalar*
    - *Alerta de programas descontinuados e da a opçao de desinstalar*
  - *Verificação de atualizações ao instalar*:
    - *É informado de update disponíveis*
    - *Verificação regularmente por updates*
  - *Instala atualizações*:
    - *Baixa a atualização num diretório temporário*
    - *Após finalizar o download, pede se quer instalar a atualização*
    - *Remove os arquivos das versões anteriores do programa*
    - *Move o update para diretório de instalação*
  - *Pré Execução dos apps*:
    - *Exibe uma splesh screen*
    - *Em segundo plano, inicia o launcher se não estiver rodando*
  - *Notificações dos apps*:
    - *Notificações Push direcionadas a cada app (ex: novo código de peças criado por usuário tal ...)*
    - *Notificações Push atualização disponível*

---

- Responsabilidades do makefile:
  - *Gerar build do launcher*
  - *Gerar o instalador via CLI do Inno Setup*

---

- Variaveis de ambiente:
  - USER_EMAIL
  - USER_TOKEN
  - USER_GITHUB_TOKEN
