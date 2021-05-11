import TodoContent from "./components/todo-content.js"
import TodoUsers from "./components/todo-users.js"
import TodoInput from "./components/todo-input.js"
import Todo from "./composables/todo.js"

export default {
    template: `
        <todo-users @get-todo="getUserTodoList"></todo-users>
        <todo-input></todo-input>
        <hr>
        <todo-content></todo-content>
    `,
    components: {
        TodoContent,
        TodoUsers,
        TodoInput
    },
    setup() {
        const {
            getUserTodoList
        } = Todo()

        return {
            getUserTodoList
        }
    }
}