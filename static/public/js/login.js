const $ = (selector) => document.querySelector(selector),
  API_ENDPOINT = '/api/v1/login';

$('#login').addEventListener('click', async (e) => {
  e.preventDefault();
  const email = $('#email').value,
    password = $('#password').value;

  if (!email || !password) return;

  const form = new FormData();
  form.append('email', email);
  form.append('password', password);

  const response = await fetch(API_ENDPOINT, {
    method: 'POST',
    body: form,
  });

  $('#email').value = '';
  $('#password').value = '';

  if (!response.ok) {
    switch (response.status) {
      case 401:
        return alert('Invalid email or password.');
      case 422:
        return alert('Please enter your email and password.');
      default:
        return alert('Something went wrong. Please try again.');
    }
  }

  console.info('User authenticated successfully. Redirecting home.');
  window.location.href = '/'; // go home on success
});
