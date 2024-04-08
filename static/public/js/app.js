const $ = (selector) => document.querySelector(selector);
const container = $('section#user');
const API_ENDPOINT = '/api/v1/user';

const getUser = async function () {
  const response = await fetch(API_ENDPOINT),
    data = await response.json(),
    { success, user } = data;

  if (!success) {
    return null;
  }

  return user;
};

function showUser(user) {
  const email = document.createElement('p'),
    token = document.createElement('p');
  email.id = 'user-email';
  token.id = 'user-token';
  email.innerHTML = `<strong>Email:</strong> ${user.email}`;
  token.innerHTML = `<strong>Token:</strong> ${user.token}`;
  container.append(email, token);
}

function showContent() {
  const user = $('section#user'),
    authMessage = $('#auth-message'),
    authenticatedMessage = $('#authenticated-message');
  user.classList.remove('d-none');
  authMessage.classList.add('d-none');
  authenticatedMessage.classList.remove('d-none');
}

async function setup() {
  // if user is authenticated, show more info
  const u = await getUser();
  if (u) {
    showContent();
    showUser(u);
  }
}

document.addEventListener('DOMContentLoaded', setup);
