import api from '@/lib/libAxios';

const base = '/api';
export function getAllWorkouts() {
  return api.get(`${base}/workouts`);
}

export function getWorkout(id) {
  return api.get(`${base}/workouts/${id}`);
}

export function createWorkout(workoutData) {
  return api.post(`${base}/workouts`, workoutData);
}

export async function updateWorkout(workoutData) {
  const payload = JSON.parse(JSON.stringify(workoutData));
  return api.put(`${base}/workouts/${payload.id}`, payload);
}

export function deleteWorkout(id) {
  return api.delete(`${base}/workouts/${id}`);
}

export async function checkToken(data) {
  const payload = JSON.stringify({ token: data });
  return api.post('/auth/google/sigin', payload);
}

export async function logOut() {
  return api.post('/auth/google/sigout');
}
