# This Todo list for solving problem 
* Next Steps: 
    - [x] Solving duplication in slices 
    - [x] For Example "makeSlices" return:  [Mehran Moradi Mehran Moradi Mehrdad Moradi].
    - [x] Change userHandeler old function with new userHandeler function and fix makeSlice func with new return userHandeler return as map variabls. (From 2:48:30 Nana tutorial)
    - [x] Make persist user data and information in mysql database.
    - [x] Cleanup mysql.go in iternal.
        - [x] Make insert function. 
        - [x] Make select function. 
    - [x] Use mysql ticket with mysql internal function.
    - [X] Making rabbitmq connection and module.
        - [x] Making publish user info throw rabbitmq.
        - [x] Making consume and insert qury from rabbitmq.
        - [ ] Make http server to api server. 
            - [x] Bind http module to input user info func and integrate with rabbitmq. 
            - [x] Make function to gatther all static pages. 
            - [x] Make return after successful submit user's info.
            - [ ] Make client authentication and authorization with keycloak.
                - [x] Make client and admin area with keycloak.
                - [x] Integrate keycloak with "/" path.
                - [ ] Make access management to "Admin area" with role defined in keycloak.
        - [ ] Making pre-check for remaining ticket number throw rabbitmq. 
            - [ ] Make change in rabbitmq producer and consumer{1. seprate "message logic as input to func RabbitmqProducer and RabbitMQ Consumer."}
        - [ ] Making exchange and queue with code and bind queue with code to exchange.
    - [x] Make userValidation Function and sperate them to cmd directory. 
    - [x] Cleanup main.go file. 
    - [ ] Make sample ui.
    - [ ] Make build directory.
    - [ ] Make clean verbose messages.
    
---

* Issues: 
    - [x] Solve "Uncorrect showing cli comment #2".
    - [x] Fix remaining ticket number. "Uncorrect remaining ticket calculation #3"
    - [x] Fix conusmer number with every time running  rabbitmqConsumer in UserInfoPost function. "Fix consumer number #3"
    - [x] Fix ticketNumber validation condition to stop producing user info's after booked out all ticket. 
    - [x] Fix user info realm_access scope within keycloak module .

