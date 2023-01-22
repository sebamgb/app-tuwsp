const form = document.querySelector('form')

form.addEventListener('submit', (e) => {
  // avoid to send
  e.preventDefault();
  // if not validateForm
  if(!validateForm()){
    return; // get out
  }
  // send form
  form.submit();
})
async function validateForm() {
  var email = document.getElementById("email").value;
  var password = document.getElementById("password").value;
  var errorMessage = document.getElementById("error-message");
  // emty values
  if (!email || !password) {
    errorMessage.innerHTML = "Por favor, complete todos los campos.";
    // don't send
    return false;
  }
  // regex to email
  if (!email.match(/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/)) {
    errorMessage.style.border = "red solid .1em";
    errorMessage.innerHTML = "Por favor, ingrese una dirección de correo electrónico válida.";
    return false;
  }
  // password major to 8
  if (password.length < 8) {
    errorMessage.style.border = "red solid .1em";
    errorMessage.innerHTML = "La contraseña debe tener al menos 8 caracteres.";
    return false;
  }
  // try for
  try {
    const response = await fetch('http://127.0.0.1:3000/validate', {
      method: 'post',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        email: email,
        password: password
      })
    });
    // assign result to wait for response
    const result = await response.json();
    // if not success
    if (!result.success) {
      errorMessage.style.color = "red";
      errorMessage.style.border = "red solid .1em";
      errorMessage.innerHTML = result.message;
      return false;
    }
    // send form
    return true;
  } catch (error) {
    errorMessage.innerHTML = "Ocurrió un error al validar los datos. Por favor, inténtalo de nuevo más tarde.";
    return false;
  }
}
