import data from "../data.js"

export default {
    template: `
    <div id="content" class="py-2">
        <h3>Todo Content</h3>
        <ol>
            <li v-for="todo in todos" class="m-3">
                {{ todo.name }}
                <br>
                <span class="text-xs text-gray-500">at {{ new Date(todo.create_at) }}</span>
            </li>
        </ol>
    </div>
    `,
    setup() {
        const {
            todos
        } = data()

        return {
            todos
        }
    }
}