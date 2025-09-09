const selectedModels = new Set();

let checkAll = null;
let dropdownTriggerAll = null;

let modelModal = null;
let modelModalTitle = null;

document.addEventListener("DOMContentLoaded", () => {
	document
		.querySelectorAll('input[id^="check-"]:checked')
		.forEach((el) => selectedModels.add(el.id.replace("check-", "")));
	dropdownTriggerAll = document.getElementById("dropdownTriggerAll");
	checkAll = document.getElementById("checkAll");
	modelModal = document.getElementById("modelModal");
	modelModalTitle = document.getElementById("modelModalTitle");
	updateUIState();
});

function updateUIState() {
	checkAll.checked =
		selectedModels.size ===
		document.querySelectorAll('input[id^="check-"]').length;
	checkAll.indeterminate = selectedModels.size > 0 && !checkAll.checked;

	dropdownTriggerAll.disabled = selectedModels.size === 0;
}

function toggleSelectAll() {
	if (checkAll.checked) {
		document
			.querySelectorAll('input[id^="check-"]:not(:checked)')
			.forEach((el) => {
				el.checked = true;
				selectedModels.add(el.id.replace("check-", ""));
			});
	} else {
		document.querySelectorAll('input[id^="check-"]:checked').forEach((el) => {
			el.checked = false;
			selectedModels.delete(el.id.replace("check-", ""));
		});
	}
	updateUIState();
}

function toggleSelectModel(modelID) {
	if (document.getElementById("check-" + modelID).checked) {
		selectedModels.add(modelID);
	} else {
		selectedModels.delete(modelID);
	}
	updateUIState();
}

function openModelModal(title, model = null) {
	if (model !== null) {
		modelModalTitle.textContent = `Edit ${title}: ${model.name}`;

		// Add a hidden id input to the modelModal
		let idInput = modelModal.querySelector("#id");
		if (!idInput) {
			idInput = document.createElement("input");
			idInput.type = "hidden";
			idInput.id = "id";
			modelModal.appendChild(idInput);
		}
		idInput.value = model.id;

		for (const key in model) {
			const input = modelModal.querySelector(`#${key}`);
			if (input) {
				input.value = model[key];
			}
		}
	} else {
		modelModalTitle.textContent = `Add ${title}`;
		modelModal.querySelectorAll("input").forEach((input) => {
			input.value = "";
		});
		const idInput = modelModal.querySelector("#id");
		if (idInput) {
			idInput.remove();
		}
	}
}
