# gochat
Use Gemini from your terminal 

## Features
    A gemini chat from terminal.
   Saves time from going online everytime or to chatgpt
    Chats saving to have access to history 
    


##  Usage:

    You must have golang installed
    Create a google Ai Studio Account [googleaistudio]("https://aistudio.google.com/") 
    Generate an api key here https://aistudio.google.com/app/apikey
    Set and environment varialble of GEMINI_API_KEY on your bashrc , ```vim ~/.bashrc``` || ```code ~/.bashrc``` for vscode then ```export GEMINI_API_KEY=<your apikey>```
    To enable chat saving set and env variable of SAVE_GO_CHAT=true
    All chats are saved at "./chats", i.e your working dir /chats 
    install this package from github with go ```go install github.com/muhammadolammi/gochat@latest```
    Now run ```gochat```
    You can create an .env file alternatively in working directly , to manage your enviromnet variables.



# Incoming Features:
    Enable history



# tips: 
    if you use vscode run '$HOMEDIR/gochat/data/chats' , so you can see each of your chats as an md file.

