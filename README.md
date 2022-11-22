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

  
  
ENGLISH VERSION
JM Social is a fullstack Go application, created during Gustavo Gallego's Go course.

The project is divided into 3 sessions:

*Mysql-Docker: Contains a docker composer script and an initial sql script, with some example data.

*Api: Consists of a back end application, made in Go, using frameworks/Packages: github.com/gorilla/mux // To Create Application Routes github.com/badoux/checkmail // Email validation github.com/ dgrijalva/jwt-go // To convert data to JSON github.com/go-sql-driver/mysql // Mysql driver github.com/joho/godotenv // Package used to read environment variables golang.org/x/crypto //Package used to encrypt and hash passwords.

*Webapp: Frontend interface, made in Go, using features like: github.com/gorilla/mux // Create web server and application routes github.com/gorilla/securecookie // Package used to interact with cookies github.com/joho /godotenv // Package used to read environment variables: .env file from the root of the application

To run the application, perform the following procedures:

1 - Create docker image for the webapp and the api:

TO CREATE API IMAGE:
  Enter the api directory and run the following command:
    sudo docker build -t jmsocial-api:latest .


TO CREATE WEBAPP IMAGE:
  Go into the webapp directory and run the following command:
    sudo docker build -t jmsocial-web:latest .

FOR MYSQL IT IS NOT NECESSARY TO CREATE AN IMAGE;

2 - Start Docker Compose:

Enter from the Mysql-docker directory and run the command:
  docker compose up

Have in mind that the connection port with the webapp is 8080

http://localhost:8080  
  
  
  
  
