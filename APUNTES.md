MVC PATTERN for structuring: model - view - controller 

    Flujo del proyecto:
- primero se ejecuta main.go 
- primero viene el view de la app (tipo el frontend)
- desp vienen los mapeos, para saber cómo responder a las diferentes requests
- los mapeos interactúan con los controllers (hacen algún control simple y de lógica)
- los controllers interactúan con los services 
- services interactúan con los repositories 
- repositories interactúan con base de datos 

    Project structure:
- application: starts and create the web server
- controllers: in charge of handling the requests and sending them to different services in order to be handled and processed
- domain: core of our entire microservice, if we are working with the users microservices, then in the domain we will have a username. 
- services: entire business logic of our application  

HTTP Framework: gin-gonic/gin provides the first layer (http) engine that we are going to use. Frameworks make things such as writing, maintaining and scaling web apps easier.

user_dto.go : data transfer object 
user_dao.go : data access object (entire logic to persist and to retrieve this user from the data base) ONLY POINT WHERE WE INTERACT WITH THE DATABASE
user_marshaller : how the domain (user) is going to be presented to the final client

Package utils will be common for the entire microservice, contains errors, and datenow (so there´s no duplicated logic)

when importing the package “name_db”
you need the func init(){

if in a function, in case of an error you are returning a json, in case of success, keep returning a json, do not change the content type. 

when we pass something like this, "/users/:user_id" user_id is a parameter from the url 
when we pass something like /internal/users/search?status=active status is a query parameter that we are passing as a query

Marshal structs : present different versions of the same struct depending on the type of request, if its private or public 