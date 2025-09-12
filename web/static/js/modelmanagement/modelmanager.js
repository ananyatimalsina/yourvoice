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

function toggleSelectModel(element) {
	if (element.checked) {
		selectedModels.add(element.id.replace("check-", "row-"));
	} else {
		selectedModels.delete(element.id.replace("check-", "row-"));
	}
	updateUIState();
}

function orderModels(element, field) {
	performSearch({ orderBy: field });
	field = field.replace("-", "");
	orderBtn = document.getElementById("order-" + field);
	orderBtn.removeChild(orderBtn.firstChild);
	orderBtn.appendChild(element.firstChild.firstChild.cloneNode(true));
	orderBtn.firstChild.classList.remove("mr-2");
	document.querySelectorAll('button[id^="order-"]').forEach((el) => {
		if (el.id !== orderBtn.id) {
			el.removeChild(el.firstChild);
			// Hardcoded default icon for now
			el.appendChild(
				new DOMParser().parseFromString(
					'<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="" data-lucide="icon"><path d="m21 16-4 4-4-4"></path><path d="M17 20V4"></path><path d="m3 8 4-4 4 4"></path><path d="M7 4v16"></path></svg>',
					"image/svg+xml",
				).documentElement,
			);
		}
	});
}
