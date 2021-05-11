// see https://v3.vuejs.org/guide/component-basics.html#using-v-model-on-components how to implment input render function
export default {
    props: ["modelValue"],
    render: function() {
        return Vue.h("input", {
            class: ['border-2', 'border-black-300', 'py-0.5', 'px-1.5'],
            value: this.modelValue,
            onInput: (e) => this.$emit("update:modelValue", e.target.value),
        })
    }
}