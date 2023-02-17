const form = document.getElementById('form');
const submitButton = document.getElementById('send');

let timeout = null;

document.querySelectorAll('.form__body__group').forEach((group) => {
    const input = group.querySelector('input');
    input.addEventListener('keydown', (e) => {
        clearTimeout(timeout);
        timeout = setTimeout(() => {
            validateInput(group, input);
        }, 500);
    });
});

async function validateInput(group, input) {
    showError(group, input.value === '', 'Este campo es obligatorio');
    if (input.id === 'email') {
        if (input.value.match(/^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/)) {
            showError(group, false, '');
        } else {
            showError(group, true, 'El email no es válido');
        }
    }
}
    
async function showError(group, condition, errorMessage) {
    if (condition) {
        group.classList.remove('success');
        group.classList.add('error');
        group.querySelector('small').innerHTML = errorMessage
        submitButton.disabled = true;
    } else {
        group.classList.remove('error');
        group.classList.add('success');
        group.querySelector('small').innerHTML = '';
        submitButton.disabled = false;
    }
}