# README

This is todo list project for practicing how to write go.  
There are two APIs, and db is file now.  

1. `curl http://localhost:7070/todo-list/get?username=user1`  
2. `curl -X POST http://localhost:7070/todo-list/create?username=user1&task=hi`  

## TODO

[ ] Change file schema, add create time for task.  
[ ] Add API to edit task name and remove task.  
[ ] Add API to add new user and remove user.  

Make sure the code and folder structure follow up the following advises.  
1. https://github.com/llitfkitfk/go-best-practice  
2. https://github.com/Ptt-official-app/Ptt-backend/issues/16  
3. https://ronmi.github.io/post/go/effectivego/  
4. https://openhome.cc/Gossip/Go/index.html  
5. https://golang.org/doc/effective_go#control-structures  
6. https://gobyexample-cn.github.io/