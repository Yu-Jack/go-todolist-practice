# README

This is todo list project for practicing how to write go.  
There are two APIs, and db is file now.  

1. `curl http://localhost:7070/todo-list/get?username=user1`  
2. `curl -X POST http://localhost:7070/todo-list/create?username=user1&task=hi`  
3. `curl -X POST http://localhost:7070/todo-list/delete?username=user1&task=hi`

## TODO

- [x] Change file schema, add create time for task.
- [x] Add API to edit task name and remove task.
- [x] Change response data structure.
- [ ] Add unit test.
- [ ] Add API to query user.
- [ ] Add API to create new user.
- [ ] Connect google calendar or google tasks.

Make sure the code and folder structure follow up the following advises.  
1. https://github.com/llitfkitfk/go-best-practice  
2. https://github.com/Ptt-official-app/Ptt-backend/issues/16  
3. https://ronmi.github.io/post/go/effectivego/  
4. https://openhome.cc/Gossip/Go/index.html  
5. https://golang.org/doc/effective_go#control-structures  
6. Go Example Chinese: https://gobyexample-cn.github.io/
7. Go Example Englisg: https://gobyexample.com/

## Data Structure

This is a very simple structure.

```json
{
  "user1": {
    "sequence": 1,
    "todo": [
      {
        "id": 1,
        "name": "eat1",
        "create_at": 1618491050221
      }
    ],
    "create_at": 1618491050221
  }
}
```
