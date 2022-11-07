(1) Equipe:   
* Caio Alexandre Campos Maciel
* Othavio Rudda da Cunha Araujo
* Pedro Henrique Alves Moutinho 

(2) O nosso projeto, chamado de SpotImages, é um sistema que oferece ao usuário, em um primeiro momento, a funcionalidade de cadastro e login na plataforma. Após isso, o mesmo pode salvar fotos e gerenciá-las, formando assim sua galeria de fotos pessoais.

(3) Tecnologias utilizadas:
* Como linguagem de programação principal foi utilizada Golang, com a interface para usuário utilizando bootstrap.

* Para a persistência dos dados, como usuários e imagens salvas, foi utilizado o banco de dados relacional PostgreSQL, por sua robustez e performance, satisfazendo assim nossas necessidades.

* Outra tecnologia também utilizada foi a do Docker, onde padronizamos assim um ambiente para que qualquer pessoa com aquele instalado na sua máquina consiga executar o nosso projeto.

(4) Para a execução do nosso projeto, basta ter o docker instalado e executar o seguinte comando:

Windows:
* sed -i 's/\r$//' startScript.sh  && chmod +x startScript.sh
* docker-compose up

Linux: 
* docker-compose up

(5) Link dos commits com refatorações:
https://github.com/KaioAlex/TP-Engenharia_de_Software_II/commit/cd36818d702e6449b3616f933cd2b0b2ca2f1993
