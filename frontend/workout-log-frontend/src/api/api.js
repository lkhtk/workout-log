import api from '@/lib/libAxios';

const base = '/api';
// const base = 'http://localhost:8000';
export function getAllWorkouts() {
  return api.get(`${base}/workouts`);
}

export function getAllWorkout(id) {
  return api.get(`${base}/workouts/${id}`);
}

export function addWorkout(workoutData) {
  return api.post(`${base}/workouts`, workoutData);
}

export function updateWorkout(workoutData) {
  const payload = {
    ID: workoutData.ID,
    name: workoutData.newName,
    description: workoutData.description,
    visible: workoutData.visible,
    tags: workoutData.tags,
  };
  return api.put(`${base}/workouts`, payload);
}

export function deleteWorkout(id) {
  return api.delete(`${base}/workouts/${id}`);
}
