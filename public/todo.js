import data from "./data.js"

export default () => {
    const {
        message,
        todos,
        taskName,
        currentUsername
    } = data()

    const resetMessage = () => {
        message.value = ""
    }

    const getUserTodoList = async (username) => {
        resetMessage()
        username = encodeURIComponent(username)
        currentUsername.value = username
        const result = await fetch(`/todo-list/get/?username=${username}`).then(response => response.json())
        todos.value = [].concat(result.todo)
    }

    const createTodo = async (username) => {
        resetMessage()
        username = encodeURIComponent(username)
        const task = encodeURIComponent(taskName.value)
        const result = await fetch(`/todo-list/create/?username=${username}&task=${task}`, {
            method: "POST"
        }).then(response => response.json())
        if (result.status !== 0) {
            message.value = `Failed to create todo, because of ${result.message}`
            return
        }
        getUserTodoList(username)
        message.value = result.message
        taskName.value = ""
    }

    return {
        message,
        resetMessage,
        todos,
        taskName,
        getUserTodoList,
        createTodo,
    }
}