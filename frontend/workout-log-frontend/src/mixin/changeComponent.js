export default {
  emits: ['change-component'],
  methods: {
    changeComponent(component, userId = 0) {
      this.$emit('change-component', { component, userId });
    },
  },
};
