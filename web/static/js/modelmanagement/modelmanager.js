const selectedModels = new Set();

let checkAll = null;
let dropdownTriggerAll = null;

function updateSelectedModels() {
	document
		.querySelectorAll('input[id^="check-"]:checked')
		.forEach((el) => selectedModels.add(el.id.replace("check-", "row-")));
	fetchElementsModelManager();
	updateUIState();
}

document.addEventListener("DOMContentLoaded", updateSelectedModels);

function fetchElementsModelManager() {
	dropdownTriggerAll = document.getElementById("dropdownTriggerAll");
	checkAll = document.getElementById("checkAll");
}

function updateUIState() {
	checkAll.checked =
		selectedModels.size ===
		document.querySelectorAll('input[id^="check-"]').length;
	checkAll.indeterminate = selectedModels.size > 0 && !checkAll.checked;

	dropdownTriggerAll.disabled = selectedModels.size < 2;
}

function toggleSelectAll() {
	if (checkAll.checked) {
		document
			.querySelectorAll('input[id^="check-"]:not(:checked)')
			.forEach((el) => {
				el.checked = true;
				selectedModels.add(el.id.replace("check-", "row-"));
			});
	} else {
		document.querySelectorAll('input[id^="check-"]:checked').forEach((el) => {
			el.checked = false;
			selectedModels.delete(el.id.replace("check-", "row-"));
		});
	}
	updateUIState();
}

function toggleSelectModel(modelID) {
	if (document.getElementById("check-" + modelID).checked) {
		selectedModels.add("row-" + modelID);
	} else {
		selectedModels.delete("row-" + modelID);
	}
	updateUIState();
}
