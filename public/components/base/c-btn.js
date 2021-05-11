export default {
    render: function() {
        return Vue.h("button", {
            class: [
                'border-2',
                'border-blue-500',
                'bg-blue-500',
                'text-white',
                'rounded',
                'px-1',
                'py-0.5'
            ]
        }, this.$slots.default())
    }
}