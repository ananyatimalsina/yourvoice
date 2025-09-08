const selectedModels = new Set();

let checkAll = null;

document.addEventListener("DOMContentLoaded", () => {
	document
		.querySelectorAll('input[id^="check-"]:checked')
		.forEach((el) => selectedModels.add(el.id.replace("check-", "")));
	checkAll = document.getElementById("checkAll");
	updateUIState();
});

function updateUIState() {
	checkAll.checked =
		selectedModels.size ===
		document.querySelectorAll('input[id^="check-"]').length;
	checkAll.indeterminate = selectedModels.size > 0 && !checkAll.checked;

	document.getElementById("dropdownTriggerAll").disabled =
		selectedModels.size === 0;
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
