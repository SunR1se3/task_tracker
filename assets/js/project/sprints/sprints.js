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
                updateSprintCards(1);
            } else {
                setErrors(resp.errors);
            }
        }
    }
    xhr.send(JSON.stringify(jsonObj));
}

function setErrors(errs){
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

function getProjectSprintCards(page) {
    return new Promise((resolve, reject) => {
        let projectId = document.querySelector('[data-project-id]').dataset.projectId;
        let orderBy  = document.getElementById('sprintsOrder').value;
        let sprintTitle  = document.getElementById('sprintTitle').value;
        let queryRequest = `?projectId=${projectId}&offset=${page}&createdAtOrder=${orderBy}`;
        console.log(sprintTitle);
        if (sprintTitle) {
            if (sprintTitle.trim() !== "") {
                queryRequest += `&title=${sprintTitle}`;
            }
        }
        let xhr = new XMLHttpRequest();
        xhr.open("GET", `/project/sprint_cards` + queryRequest);
        xhr.onreadystatechange = function () {
            if (this.readyState === 4) {
                let resp = JSON.parse(xhr.responseText);
                if (resp.status) {
                    resolve(resp.data);
                } else {
                    reject(new Error("Ошибка запроса"));
                }
            }
        }
        xhr.send();
    });
}

function reorderSprints(e) {

}

function updateSprintCards(page) {
    getProjectSprintCards(page)
        .then(data => {
            let sprintCards = document.getElementById('sprintCards');
            sprintCards.innerHTML = data;
        })
        .catch(error => {
            console.error(error);
        });
}

$('#pagination-demo').twbsPagination({
    totalPages: 12,
    visiblePages: 7,
    first: 'В начало',
    last: 'В конец',
    next: false,
    prev: false,
    onPageClick: function (event, page) {
        updateSprintCards(page);
    }
});

