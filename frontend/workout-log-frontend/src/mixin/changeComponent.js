export default {
  emits: ['change-component'],
  methods: {
    changeComponent(component, payload) {
      this.$emit('change-component', { component, payload });
    },
  },
};
