export default {
  emits: ['change-component'],
  methods: {
    changeComponent(component, workoutId) {
      this.$emit('change-component', { component, workoutId });
    },
  },
};
