# jmsocial
JM Social é uma aplicação fullstack em Go, criado durante o curso de Go do Gustavo Gallego.

O projeto é dividido em 3 aplicações:

*Mysql-Docker: Contem script do docker composer e um script sql inicial, com alguns dados de exemplos.

*Api:  Consiste numa aplicação back end, feita em Go, utilizando frameworks/Pacotes:
github.com/gorilla/mux   // Para Criar Rotas da aplicação
github.com/badoux/checkmail  // Validação de e-mail
github.com/dgrijalva/jwt-go  // Para converter dados em JSON 
github.com/go-sql-driver/mysql   // Driver Mysql
github.com/joho/godotenv  // Pacote utilizado para ler variaveis de ambinete
golang.org/x/crypto  //Pacote utilizado para criptografar e gerar hash das senhas.


*Webapp:  Interface frontend, feita em Go,  utlizando recursos como:
github.com/gorilla/mux   // Criar servidor web e rotas da aplicação
github.com/gorilla/securecookie  // Pacote utilizado para interagir com cookies
github.com/joho/godotenv // Pacote utilizado para ler as variaveis de ambiente:  arqivo  .env da raiz da aplicação
  
  
  Para rodar a aplicação é necessario realizar os seguintes procedimentos:
  
  1 - Criar imagem docker para o webapp e para a api:
  
    PARA CRIAR IMAGEM DO API:
      Entra no diretorio api e executa o seguinte comando:
        sudo docker build -t jmsocial-api:latest .


    PARA CRIAR IMAGEM DO WEBAPP:
      Entra no diretorio webapp e executa o seguinte comando:
        sudo docker build -t jmsocial-web:latest .
  
    PARA O MYSQL NÃO É NECESSARIO CRIAR IMAGEM;

  2 - Iniciar o Docker Compose: 

    Entra do diretorio Mysql-docker e executa o comando:
      docker compose up 
      
  
  lembrando que a porta de conexão com a webapp é  8080

  http://localhost:8080

  
  
  
  
  
  
  
