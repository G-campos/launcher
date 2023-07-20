- Funções do launcher:
  - *Autenticação do usuário*
  - *Listar programas*
  - *Baixar e instalar os programas selecionados*
  - *Verificação de atualizações ao instalar*
  - *Verificação de updates regulares*
  - *Instala atualizações*
  - *Em caso de usuário inativo, inutiliza os apps*
  - *Ao abrir um app, inicia o launcher se não estiver rodando*
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
  - *Baixar e instalar os programas selecionados*:
    - *Baixa o programa num diretório temporário*
    - *Após finalizar o download, move para diretório de instalação*
  - *Verificação de atualizações ao instalar*:
    - *É informado de update disponíveis*
    - *Verificação regularmente por updates*
  - *Instala atualizações*:
    - *Baixa a atualização num diretório temporário*
    - *Após finalizar o download, pede se quer instalar a atualização*
    - *Remove os arquivos das versões anteriores do programa*
    - *Move o update para diretório de instalação*
  - *Ao abrir um app, inicia o launcher se não estiver rodando*
  - *Notificações dos apps*:
    - *Notificações Push direcionadas a cada app (ex: código inserido por usuário ...)*
    - *Notificações Push atualização disponível*

---

- Responsabilidades do makefile:
  - *Gerar build do launcher*
  - *Gerar instalador via CLI do Inno Setup*

---

- Variaveis de ambiente:
  - USER_EMAIL
  - USER_TOKEN
  - USER_GITHUB_TOKEN
