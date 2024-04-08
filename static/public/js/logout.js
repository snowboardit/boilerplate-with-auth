const API_ENDPOINT = '/api/v1/logoff';

document.addEventListener('DOMContentLoaded', async function () {
  const response = await fetch(API_ENDPOINT),
    data = await response.json();

  if (response.ok && data.success) {
    return (window.location.href = '/');
  }
  alert(data.message);
  return (window.location.href = '/');
});
