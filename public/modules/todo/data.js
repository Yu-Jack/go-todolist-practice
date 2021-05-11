const todos = Vue.ref([])
const users = Vue.ref([])
const taskName = Vue.ref("")
const currentUsername = Vue.ref("")
const message = Vue.ref("")

export default () => {
    return {
        todos,
        users,
        taskName,
        currentUsername,
        message,
    }
}