import data from "../data.js"

export default {
    template: `
    <div>
        <h3>All Users, click user to see the todo list</h3>
        <ol>
            <li v-for="user in users">
                <a href="#" class="text-blue-500" @click="$emit('get-todo', user.name)">{{ user.name }}</a>
            </li>
        </ol>
    </div>
    `,
    setup() {
        const {
            users
        } = data()

        Vue.onMounted(async () => {
            const result = await fetch("/users").then(response => response.json())
            users.value = users.value.concat(result.users)
        })

        return {
            users
        }
    }
}