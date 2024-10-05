import api from '@/lib/libAxios';

const defaultPath = '/api/album';

export function getPhotos(dir) {
  return api.get(`${defaultPath}/${dir}`);
}
export function getAlbums() {
  return api.get('/api/albums');
}
export function getAll() {
  return api.get(defaultPath);
}
export function getAuth(login, pass) {
  return api.post('/api/login', { username: login, password: pass });
}
export function logout() {
  return api.post('/api/signout');
}
export function createAlb(albumData) {
  return api.post(defaultPath, albumData);
}

export function updateAlb(albumData) {
  const payload = {
    ID: albumData.ID,
    name: albumData.newName,
    description: albumData.description,
    visible: albumData.visible,
    tags: albumData.tags,
  };
  return api.put(defaultPath, payload);
}
export function deleteAlb(name) {
  return api.delete(`${defaultPath}/${name}`);
}
export function uploadImages(name, formData) {
  return api.post(`${defaultPath}/${name}/photo`, formData, {
    maxContentLength: -1,
    maxBodyLength: -1,
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
}
export function deleteImg(albName, fileName) {
  return api.delete(`${defaultPath}/${albName}/delete/${fileName}`);
}
export function sendMsg(msg) {
  return api.post('/api/msg', msg);
}
export function getMsgs() {
  return api.get('/api/msg');
}
export function delMsg(id) {
  return api.delete(`/api/msg/${id}`);
}
export function getConf() {
  return api.get('/api/config');
}
export function delConf(id) {
  return api.delete(`/api/config/${id}`);
}
export function addParam(param) {
  return api.post('/api/config', param);
}
export function getPersonInfo() {
  return api.get('/api/person');
}
export function setPersonInfo(param) {
  return api.put('/api/person', param);
}
