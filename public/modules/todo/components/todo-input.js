import data from "../data.js"
import Todo from "../composables/todo.js"

import CBtn from "../../../components/base/c-btn.js"
import CInput from "../../../components/base/c-input.js"

export default {
    template: `
    <div id="todo-input" class="py-2" v-if="currentUsername">
        <hr>
        <label>You're now {{ currentUsername }}, create task for {{ currentUsername }} </label>
        <br>
        <c-input type="text" :model-value="taskName" @update:model-value="taskName = $event" @keyup="resetMessage" @keyup.enter="createTodo"></c-input> 
        <c-btn type="button" @click="createTodo(currentUsername)">Add</c-btn>
        <p> {{ message }} </p>
    </div>
    `,
    components: {
        CInput,
        CBtn
    },
    setup() {
        const {
            currentUsername,
            taskName,
            message
        } = data()

        const {
            createTodo,
            resetMessage,
        } = Todo()

        return {
            currentUsername,
            taskName,
            message,
            createTodo,
            resetMessage
        }
    }
}