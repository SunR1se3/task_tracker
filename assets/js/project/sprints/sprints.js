function createSprint() {
    resetErrors();
    // создать объект для формы
    let formData = new FormData(document.forms.sprintCreateForm);
    let jsonObj = {}
    formData.forEach((value, key) => {
        jsonObj[key] = value;
    });
    let projectId = document.querySelector('[data-project-id]').dataset.projectId;
    let xhr = new XMLHttpRequest();
    xhr.open("POST", `/api/sprint/${projectId}`);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            let resp = JSON.parse(xhr.responseText)
            if (resp.status) {
                bootstrap.Modal.getInstance(document.getElementById('sprintCreateModal')).hide();
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
    let sprintCreateForm = document.forms.sprintCreateForm;
    let inputs = Array.from(sprintCreateForm.elements);
    inputs.forEach((item) => {
        item.classList.remove('is-invalid');
    })
}

function sprintPagination() {
    $('#pagination-demo').twbsPagination({
        totalPages: 16,
        visiblePages: 6,
        next: 'Next',
        prev: 'Prev',
        onPageClick: function (event, page) {
            $('#page-content').text('Page ' + page) + ' content here';
        }
    });
}