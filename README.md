# To Do List
A Todo List application in Golang which exposes an API.

# How to use
1. Run `make serve` locally
    - What's `make` you ask? Click [here](https://letmegooglethat.com/?q=GNU+make)!
2. Do a `POST` call to `http://localhost:8080/todos` using your favorite API Client using the following body

```
{
    "id": "6080a379-60f7-447f-a018-da8e6a92d2a5",
    "title": "pizza",
    "dueDate": "2022-07-07T17:20:50.52Z",
    "description": "eat pizza and be happy",
    "labels": ["food", "love"]
}
```
3. Do a `GET` call to `http://localhost:8080/todos` and see if you get a collection of todos
4. Do a `GET` call to `http://localhost:8080/todos/6080a379-60f7-447f-a018-da8e6a92d2a5` and check if you get the same todo
5. Repeat step 2 & 3 without the id and edit the body
```
{
    "title": "pizza",
    "dueDate": "2022-07-07T17:20:50.52Z",
    "description": "eat pizza and be happy",
    "labels": ["food", "love"]
}
```
6. Go crazy!!!

# To Dos
1. Add updating of todos
2. Add deletion of todos
3. Make isDone a timestamp called done_at
4. Add unit tests
