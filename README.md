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
  
  
  Para rodar a aplicação é necessario iniciar o docker
  subir a API
  subir a Webapp
  
  lembrando que a porta de conexão com a webapp é  8080
  
  A plicação pode ser acessada no link:  www.jmcloudpro.com/jmsocial
  
  
  
  
