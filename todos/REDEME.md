```
mutation createTodo{
  createTodo(todo:{text:"nice",done:false}){
    done
    text
  }
}

query todos{
  todos{
    id
    text
    done
  }
}

query lasttodo{
  lastTodo{
    id
    text
    done
  }
}


query todo{
   todo(id:1){
    done
    text
  }
}

mutation updateTodo{
  updateTodo(id:1,changes:{text:"golang",done:false}){
    done
    text
  }
}
```