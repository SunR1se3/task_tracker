function projectAutocomplete(data) {
    const autocompleteData = data;

    const autocompleteInput = document.getElementById('autocompleteInput');

    autocompleteInput.addEventListener('input', function () {
        const inputText = this.value.toLowerCase();
        const suggestionsDropdown = document.getElementById('suggestionsDropdown');

        if (!suggestionsDropdown) {
            return;
        }

        const suggestionsMenu = suggestionsDropdown.querySelector('.dropdown-menu');
        suggestionsMenu.innerHTML = '';

        const suggestions = autocompleteData.filter(item => item.fio.toLowerCase().includes(inputText));
        let i = 0
        suggestions.forEach(suggestion => {
            const listItem = document.createElement('a');
            listItem.classList.add('dropdown-item');
            listItem.href = '#';
            listItem.id = `${suggestion.id}`
            listItem.textContent = suggestion.fio;

            listItem.addEventListener('click', function () {
                autocompleteInput.value = suggestion.fio;
                autocompleteInput.dataset.selectedId = suggestion.id
                suggestionsMenu.innerHTML = '';
                setVisibilityAddButton();
            });
            suggestionsMenu.appendChild(listItem);
        });
        if (autocompleteInput.value === "") {
            autocompleteInput.removeAttribute('data-selected-id');
            setVisibilityAddButton();
        }
        setVisibilityClearButton();
        suggestionsMenu.style.display = suggestions.length > 0 ? 'block' : 'none';
    });

    document.addEventListener('click', function (event) {
        const suggestionsDropdown = document.getElementById('suggestionsDropdown');
        if (!event.target.closest('#suggestionsDropdown')) {
            const suggestionsMenu = suggestionsDropdown.querySelector('.dropdown-menu');
            suggestionsMenu.style.display = 'none';
        }
    });
}

function initUserPicker() {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", `/api/user/picker`);
    xhr.onreadystatechange = function() {
        if (this.readyState === 4) {
            let resp = JSON.parse(xhr.responseText)
            if (resp.status) {
                return projectAutocomplete(JSON.parse(this.responseText).data)
            }
        }
    }
    xhr.send();
}

function setVisibilityAddButton() {
    let autocomplete = document.getElementById('autocompleteInput');
    let addBtn = document.getElementById('addBtn');

    if (autocomplete.hasAttribute('data-selected-id') && addBtn.style.display === 'none') {
        addBtn.style.display = 'block';
    }

    else if (!autocomplete.hasAttribute('data-selected-id') && addBtn.style.display === 'block') {
        addBtn.style.display = 'none';
    }
}

function setVisibilityClearButton() {
    let autocomplete = document.getElementById('autocompleteInput');
    let clearBtn = document.getElementById('clearBtn');
    if (autocomplete.value !== "") {
        clearBtn.style.display = 'block';
    } else {
        clearBtn.style.display = 'none';
    }
}

function clearInput() {
    let autocomplete = document.getElementById('autocompleteInput');
    autocomplete.value = "";
    autocomplete.removeAttribute('data-selected-id');
    setVisibilityClearButton();
    setVisibilityAddButton();
}