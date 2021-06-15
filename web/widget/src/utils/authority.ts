export function getApiSid(): string {
  return localStorage.getItem('sid') || '';
};

export function setApiSid(sid: string) {
  localStorage.setItem('sid', sid);
};
