function changePassword() {
    resetErrors();
    let formData = new FormData(document.forms.userChangePasswordForm);
    let jsonObj = {}
    formData.forEach((value, key) => {
        jsonObj[key] = value;
    });
    let xhr = new XMLHttpRequest();
    xhr.open("PUT", "/api/user/change_password");
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            let resp = JSON.parse(xhr.responseText)
            if (resp.status) {
                let alertChangePassword = document.getElementById('alertChangePassword')
                alertChangePassword.hidden = false
                setTimeout(function() {
                    alertChangePassword.hidden = true
                }, 3000)
            } else {
                setErrors(resp.errors);
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}

function setErrors(errs) {
    for (let i = 0; i < errs.length; i++) {
        let errBlocks = document.querySelectorAll('[data-err-field-name]');
        errBlocks.forEach((item) => {
            if (item.getAttribute('data-err-field-name') === errs[i].split('|')[0]) {
                item.innerHTML = errs[i].split('|')[1];
                item.previousElementSibling.classList.add('is-invalid');
            }
        })
    }
}

function resetErrors() {
    let passwordForm = document.forms.userChangePasswordForm;
    let inputs = Array.from(passwordForm.elements);
    inputs.forEach((item) => {
        item.classList.remove('is-invalid');
    })
}