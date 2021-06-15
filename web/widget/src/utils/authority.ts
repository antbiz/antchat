export function getApiSid(): string {
  return localStorage.getItem('sid') || '';
};

export function setApiSid(sid: string) {
  localStorage.setItem('sid', sid);
};

export function getVid(): string {
  return localStorage.getItem('vid') || '';
};

export function setVid(vid: string) {
  localStorage.setItem('vid', vid);
};
