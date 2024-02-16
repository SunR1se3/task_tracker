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
            });

            suggestionsMenu.appendChild(listItem);
        });

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