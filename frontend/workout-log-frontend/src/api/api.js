import api from '@/lib/libAxios';

const base = '/api';
export function getAllWorkouts(params = {}) {
  const baseUrl = '/api/workouts';
  const query = new URLSearchParams();

  query.set('page', params.page && params.page > 0 ? params.page : 1);

  if (params.date) {
    query.set('date', params.date);
  }

  if (params.gymType) {
    query.set('gymType', params.gymType);
  }

  if (params.muscle) {
    query.set('muscle', params.muscle);
  }

  const url = `${baseUrl}?${query.toString()}`;
  return api.get(url);
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

export function createMeasurement(measurement) {
  return api.post(`${base}/measurements`, measurement);
}

export function getLastMeasurement() {
  return api.get(`${base}/measurements/last`);
}
export function getMeasurements() {
  return api.get(`${base}/measurements`);
}

export function exportData() {
  return api.get(`${base}/user/export`, { responseType: 'blob' });
}

export function wipeData() {
  return api.post(`${base}/user/wipe`);
}

export function deleteUser() {
  return api.delete(`${base}/user`);
}

export function getTrends(period) {
  return api.get(`${base}/workouts/average?size=${period}`);
}

export function getTop5(period) {
  return api.get(`${base}/workouts/top5?size=${period}`);
}
